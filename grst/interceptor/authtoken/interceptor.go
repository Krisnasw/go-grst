package authtoken

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	grst_errors "github.com/krisnasw/go-grst/grst/errors"
	authtokenapi "github.com/krisnasw/go-grst/grst/interceptor/authtoken/clients/authtokenapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// This is just example of auth token interceptor.

func UnaryServerInterceptor(authTokenClient authtokenapi.AuthTokenClient, fullMethods []string) grpc.UnaryServerInterceptor {
	fullMethodsMap := map[string]bool{}
	for _, fm := range fullMethods {
		fullMethodsMap[fm] = true
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if _, ok := fullMethodsMap[info.FullMethod]; !ok {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, grst_errors.New(int32(http.StatusInternalServerError), codes.Internal, 98001, "Failed to parse context")
		} else if len(md["authorization"]) <= 0 {
			return nil, grst_errors.New(int32(http.StatusBadRequest), codes.InvalidArgument, 98002, "Authorization on header is required")
		}

		token := md["authorization"][0]
		reqBody, err := json.Marshal(authtokenapi.ValidateTokenReq{AccessToken: token})
		if err != nil {
			return nil, grst_errors.New(500, codes.Unknown, 98101, err.Error())
		}
		successResp, errResp := authTokenClient.ValidateToken(bytes.NewBuffer(reqBody))
		if errResp != nil {
			return nil, grst_errors.New(int32(errResp.HTTPStatus), codes.Code(errResp.GRPCStatus), int32(errResp.Code), errResp.Message)
		}

		if len(successResp.Data) > 0 {
			for key, value := range successResp.Data {
				md.Set("grst."+key, fmt.Sprintf("%v", value))
			}
			ctx = metadata.NewIncomingContext(ctx, md)
		}
		return handler(ctx, req)
	}
}
