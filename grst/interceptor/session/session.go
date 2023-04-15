package session

import (
	"context"

	"github.com/google/uuid"
	grst_context "github.com/krisnasw/go-grst/grst/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	CONTEXT_SESSION_ID grst_context.ContextKey = "session.id"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		sessionIds := md.Get("grst." + CONTEXT_SESSION_ID.String())
		if len(sessionIds) == 0 {
			md.Set("grst."+CONTEXT_SESSION_ID.String(), uuid.New().String())
			ctx = metadata.NewIncomingContext(ctx, md)
		} else if sessionIds[0] == "" {
			md.Set("grst."+CONTEXT_SESSION_ID.String(), uuid.New().String())
			ctx = metadata.NewIncomingContext(ctx, md)
		}
		return handler(ctx, req)
	}
}
