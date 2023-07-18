package grst

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/krisnasw/go-grst/grst/builtin/healthcheck"
	"github.com/krisnasw/go-grst/grst/builtin/validationrule"
	grst_context "github.com/krisnasw/go-grst/grst/context"
	"github.com/krisnasw/go-grst/grst/errors"
	runtimeint "github.com/krisnasw/go-grst/grst/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type OptionFunc func(*Server) error
type Server struct {
	// GRPC
	grpcServer       *grpc.Server
	grpcPort         int
	grpcServerOption GRPCServerOption

	// REST
	enableRest             bool
	restServer             *http.Server
	restPort               int
	restServerOption       RESTServerOption
	restHandlerAdditionals map[string]RestHandlerAdditional
	corsPolicy             CORSPolicy
	enableHealthcheck      bool

	// Grpc Gateway
	grpcGatewayMuxServer   *runtime.ServeMux
	grpcClientForRest      *grpc.ClientConn
	restHandlers           []RestHandler
	errorHandler           runtime.ErrorHandlerFunc
	forwardResponseMessage func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error)
	forwardHeaderToContext []string
	forwardResponseOption  func(ctx context.Context, req http.ResponseWriter, resp proto.Message) error
	// other
	lock *sync.Mutex
}

func NewServer(grpcPort int, restPort int, enableRest bool, options ...OptionFunc) (*Server, error) {
	jsonpbMarshaller := &runtime.JSONPb{}
	jsonpbMarshaller.EmitUnpopulated = true
	s := &Server{
		// GRPC
		grpcServer: nil,
		grpcPort:   grpcPort,
		grpcServerOption: GRPCServerOption{
			unaryServerInterceptors: map[string]grpc.UnaryServerInterceptor{},
			maxRecvMsgSize:          1024 * 1024 * 20,
			maxSendMsgSize:          1024 * 1024 * 20,
			maxMsgSize:              1024 * 1024 * 20,
		},

		// REST
		enableRest: enableRest,
		restServer: nil,
		restPort:   restPort,
		restServerOption: RESTServerOption{
			readTimeout:  5 * time.Second,
			writeTimeout: 10 * time.Second,
		},
		restHandlers:           []RestHandler{},
		restHandlerAdditionals: map[string]RestHandlerAdditional{},
		corsPolicy: CORSPolicy{
			allowedOrigin:    map[string]bool{"*": true},
			allowAllOrigin:   true,
			preflightHeaders: DefaultCORSPreflightHeaders,
			preflightMethods: DefaultCORSPreflightMethods,
			withCredential:   false,
		},
		enableHealthcheck: true,

		// Grpc Gateway
		grpcGatewayMuxServer:   nil,
		grpcClientForRest:      nil,
		errorHandler:           errors.HTTPError,
		forwardResponseMessage: runtimeint.ForwardResponseMessage,
		forwardHeaderToContext: []string{
			grst_context.CONTEXT_CLIENT_APPNAME.String(),
			grst_context.CONTEXT_CLIENT_APPVERSION.String(),
			grst_context.CONTEXT_CLIENT_APPVERSIONCODE.String(),
			grst_context.CONTEXT_CLIENT_MANUFACTURER.String(),
			grst_context.CONTEXT_CLIENT_MODEL.String(),
			grst_context.CONTEXT_CLIENT_PLATFORM.String(),
			grst_context.CONTEXT_CLIENT_PLATFORMVERSION.String(),
			grst_context.CONTEXT_CLIENT_SDKVERSION.String(),
			"apikey",
		},
		forwardResponseOption: func(ctx context.Context, req http.ResponseWriter, resp proto.Message) error {
			return nil
		},
		// other
		lock: &sync.Mutex{},
	}

	// Run the options on it
	for _, option := range options {
		if err := option(s); err != nil {
			return nil, err
		}
	}

	s.grpcGatewayMuxServer = runtime.NewServeMux(
		runtime.WithForwardResponseOption(s.forwardResponseOption),
		runtime.WithErrorHandler(s.errorHandler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpbMarshaller),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			startTime := time.Now().UTC().Format(time.RFC3339Nano)
			req.Header.Set("grst.starttime", startTime)
			mdMap := map[string]string{
				"grst.starttime": startTime,
			}
			for _, header := range s.forwardHeaderToContext {
				if header == "starttime" {
					continue
				}
				headerValue := req.Header.Get(header)
				if headerValue != "" {
					mdMap["grst."+header] = req.Header.Get(header)
				}
			}
			md := metadata.New(mdMap)
			return md
		}),
	)

	// initialize built in validator rule
	validationrule.Initialize()

	s.grpcServer = s.newGrpcServer()
	s.restServer = s.newRestServer()

	var err error
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s.grpcClientForRest, err = grpc.Dial(s.GetGrpcAddr(), opts...)
	if err != nil {
		return nil, err
	}

	if s.enableHealthcheck {
		healthcheckHndl := healthcheck.New(s.grpcClientForRest, "_status")
		s.RegisterRestHandlerAdditional("GET", healthcheckHndl.GetEndpointURL(), healthcheckHndl.HealthcheckHandler)
	}
	return s, err
}

func (s *Server) ListenAndServeGrst() <-chan error {
	ch := make(chan error, 2)
	go func() {
		ch <- s.ListenAndServeGRPC()
	}()

	if s.enableRest {
		go func() {
			ch <- s.ListenAndServeREST()
		}()
	}

	return ch
}

// WithCustomErrorHandler To change error response handler (including error format)
func WithCustomErrorHandler(customErrorHandler runtime.ErrorHandlerFunc) OptionFunc {
	return func(s *Server) error {
		s.errorHandler = customErrorHandler
		return nil
	}
}

// WithCustomResponseMessageForwarder To change/manipulate response message (including success response format and behavior)
func WithCustomResponseMessageForwarder(customForwardResponseMessage func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error)) OptionFunc {
	return func(s *Server) error {
		s.forwardResponseMessage = customForwardResponseMessage
		return nil
	}
}

// WithCustomCORSOrigin default Access-Control-Allow-Origin: *
func WithCustomCORSOrigins(origins []string) OptionFunc {
	return func(s *Server) error {
		allowedOrigin := map[string]bool{}
		allowAllOrigin := false
		for _, origin := range origins {
			allowedOrigin[origin] = true
			if origin == "*" {
				allowAllOrigin = true
			}
		}
		s.lock.Lock()
		defer s.lock.Unlock()
		s.corsPolicy.allowedOrigin = allowedOrigin
		s.corsPolicy.allowAllOrigin = allowAllOrigin
		return nil
	}
}

// WithCustomCORSCredential Access-Control-Allow-Credentials. default false
func WithCustomCORSCredential(enable bool) OptionFunc {
	return func(s *Server) error {
		s.corsPolicy.withCredential = enable
		return nil
	}
}

// WithCustomCORSPreflightHeader default Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, appname, appversion, appversioncode, manufacturer, model, platform, platformversion, sdkversion, User-Agent (grst.DefaultCORSPreflightHeaders)
func WithCustomCORSPreflightHeaders(headers []string) OptionFunc {
	return func(s *Server) error {
		s.corsPolicy.preflightHeaders = headers
		return nil
	}
}

// default default Access-Control-Allow-Methods: GET, HEAD, POST, PUT, DELETE (grst.DefaultCORSPreflightMethods)
func WithCustomCORSPreflightMethods(methods []string) OptionFunc {
	return func(s *Server) error {
		s.corsPolicy.preflightMethods = methods
		return nil
	}
}

// WithCustomCORSPreflightHeader default Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, appname, appversion, appversioncode, manufacturer, model, platform, platformversion, sdkversion, User-Agent (grst.DefaultCORSPreflightHeaders)
func WithCustomForwardResponseOption(forwardResponseOption func(ctx context.Context, req http.ResponseWriter, resp proto.Message) error) OptionFunc {
	return func(s *Server) error {
		s.forwardResponseOption = forwardResponseOption
		return nil
	}
}

func SetGrpcServerOption(maxRecvMsgSize, maxSendMsgSize, maxMsgSize int) OptionFunc {
	return func(s *Server) error {
		if maxRecvMsgSize > 0 {
			s.grpcServerOption.maxRecvMsgSize = maxRecvMsgSize
		}
		if maxSendMsgSize > 0 {
			s.grpcServerOption.maxSendMsgSize = maxSendMsgSize
		}
		if maxMsgSize > 0 {
			s.grpcServerOption.maxMsgSize = maxMsgSize
		}
		return nil
	}
}

func SetRestServerOption(readTimeout time.Duration, writeTimeout time.Duration) OptionFunc {
	return func(s *Server) error {
		s.restServerOption.readTimeout = readTimeout
		s.restServerOption.writeTimeout = writeTimeout
		return nil
	}
}

// UseDefaultHealthcheck using default healthcheck endpoint. default: true, you can access on http://{host}/_status
func UseDefaultHealthcheck(ok bool) OptionFunc {
	return func(s *Server) error {
		s.enableHealthcheck = ok
		return nil
	}
}

// AddForwardHeaderToContext to forward request header to handler's metadata context (with prefix grst.*). For example, if you pass req header "country", then you can access the data in metadata context with grst.country | Notes: starttime / grst.starttime is reserved
func AddForwardHeaderToContext(headerKey []string) OptionFunc {
	return func(s *Server) error {
		s.forwardHeaderToContext = append(s.forwardHeaderToContext, headerKey...)
		return nil
	}
}

func RegisterGRPCUnaryInterceptor(name string, unaryServerInterceptor grpc.UnaryServerInterceptor, includeMethods ...string) OptionFunc {
	return func(s *Server) error {
		if len(includeMethods) > 0 {
			includeMethodsMap := make(map[string]bool)
			for _, m := range includeMethods {
				includeMethodsMap[m] = true
			}
			s.grpcServerOption.unaryServerInterceptors[name] = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				if _, ok := includeMethodsMap[info.FullMethod]; ok {
					return unaryServerInterceptor(ctx, req, info, handler)
				}
				return handler(ctx, req)
			}
			return nil
		}

		s.grpcServerOption.unaryServerInterceptors[name] = unaryServerInterceptor
		return nil
	}
}

// GetGrpcServer get current grpc server object
func (s *Server) GetGrpcServer() *grpc.Server {
	return s.grpcServer
}

// GetRestServer get current rest server object
func (s *Server) GetRestServer() *http.Server {
	return s.restServer
}

// GetGrpcAddr will return return :{port}
func (s *Server) GetGrpcAddr() string {
	return fmt.Sprintf(":%d", s.grpcPort)
}

// GetRestPort will return return :{port}
func (s *Server) GetRestAddr() string {
	return fmt.Sprintf(":%d", s.restPort)
}

// RegisterRestHandler registering rest handler from grpc-gateway
func (s *Server) RegisterRestHandler(h RestHandler) {
	s.restHandlers = append(s.restHandlers, h)
}

// RegisterRestHandlerAdditional registering rest handler additional (customize / not from *.proto)
func (s *Server) RegisterRestHandlerAdditional(method, endpoint string, hf runtime.HandlerFunc) {
	s.restHandlerAdditionals[method+"|"+endpoint] = RestHandlerAdditional{method, endpoint, hf}
}

// GetForwardResponseMessage get current forwardResponseMessage (useful for intercepting response)
func (s *Server) GetForwardResponseMessage() func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	return s.forwardResponseMessage
}
