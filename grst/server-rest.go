package grst

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
)

type RESTServerOption struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
}
type RestHandler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
type RestHandlerAdditional struct {
	Method   string
	Endpoint string
	Handler  runtime.HandlerFunc
}

func (s *Server) newRestServer() *http.Server {
	return &http.Server{
		Addr:         s.GetRestAddr(),
		Handler:      nil,
		ReadTimeout:  s.restServerOption.readTimeout,
		WriteTimeout: s.restServerOption.writeTimeout,
	}
}

func (s *Server) ListenAndServeREST() error {
	log.Println(fmt.Sprintf("Initializing REST server :%d", s.restPort))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	s.restServer.Handler = wrapMuxWithCors(mux, s.corsPolicy)

	// init rest handler
	for _, h := range s.restHandlers {
		if err := h(ctx, s.grpcGatewayMuxServer, s.grpcClientForRest); err != nil {
			return err
		}
	}

	// init rest handler additional
	for _, v := range s.restHandlerAdditionals {
		s.grpcGatewayMuxServer.Handle(v.Method, runtime.MustPattern(runtime.NewPattern(1, []int{int(utilities.OpLitPush), 0}, []string{v.Endpoint}, "")), runtime.HandlerFunc(v.Handler))
	}

	// spawn all rest handler
	mux.Handle("/", s.grpcGatewayMuxServer)

	log.Println(fmt.Sprintf("REST server is started :%d", s.restPort))
	if err := s.restServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}
