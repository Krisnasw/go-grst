package logger

import (
	"context"
	"os"
	"time"

	"github.com/krisnasw/go-grst/grst/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor function
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		grpcStatus := &errors.Error{HTTPStatus: 200, GRPCStatus: 0, Code: 0, Message: ""}
		if err != nil {
			grpcStatus, _ = errors.NewFromError(err)
		}

		hostname, errOs := os.Hostname()
		if errOs != nil {
			hostname = "unknown"
		}

		md, _ := metadata.FromIncomingContext(ctx)
		sessionId := ""
		if len(md.Get("grst.session.id")) > 0 {
			sessionId = md.Get("grst.session.id")[0]
		}

		latency := time.Since(start)
		l := logrus.WithTime(time.Now().UTC()).
			WithField("hostname", hostname).
			WithField("http_status", grpcStatus.HTTPStatus).
			WithField("grpc_status", grpcStatus.GRPCStatus).
			WithField("full_method", info.FullMethod).
			WithField("latency", latency).
			WithField("request_id", sessionId).
			WithField("payload", req)

		if grpcStatus.HTTPStatus > 499 {
			l.Errorln(grpcStatus.Message)
		} else if grpcStatus.HTTPStatus > 399 {
			l.Warnln(grpcStatus.Message)
		} else {
			l.Infoln(grpcStatus.Message)
		}

		return resp, err
	}
}
