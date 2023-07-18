package grst

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	outgoingInterceptor "github.com/krisnasw/go-grst/grst/interceptor/outgoing"
	"google.golang.org/grpc"
)

type GRPCServerOption struct {
	unaryServerInterceptors map[string]grpc.UnaryServerInterceptor
	maxRecvMsgSize          int
	maxSendMsgSize          int
	maxMsgSize              int
}

func (s *Server) newGrpcServer() *grpc.Server {
	s.lock.Lock()
	defer s.lock.Unlock()

	unaryInterceptors := []grpc.UnaryServerInterceptor{}
	for _, i := range s.grpcServerOption.unaryServerInterceptors {
		unaryInterceptors = append(unaryInterceptors, i)
	}
	unaryInterceptors = append(unaryInterceptors, outgoingInterceptor.UnaryServerInterceptor())
	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(unaryInterceptors...),
		grpc.MaxRecvMsgSize(s.grpcServerOption.maxRecvMsgSize),
		grpc.MaxSendMsgSize(s.grpcServerOption.maxSendMsgSize),
		grpc.MaxMsgSize(s.grpcServerOption.maxMsgSize),
	)

	return srv
}

func (s *Server) ListenAndServeGRPC() error {
	log.Println(fmt.Sprintf("Initializing gRPC server :%d", s.grpcPort))
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		return fmt.Errorf("Initializing gRPC server :%d. %v", s.grpcPort, err)
	}

	log.Println(fmt.Sprintf("gRPC server is started :%d", s.grpcPort))
	err = s.grpcServer.Serve(l)
	return err
}
