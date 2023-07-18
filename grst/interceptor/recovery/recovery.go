package recovery

import (
	"context"
	"net/http"
	"runtime/debug"

	"github.com/krisnasw/go-grst/grst/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(http.StatusInternalServerError, codes.Internal, 99999, "Unexpected error: recovering is start")
				logrus.WithField("stack", string(debug.Stack())).Errorln("panic recovered:", r)
			}
		}()

		return handler(ctx, req)
	}
}
