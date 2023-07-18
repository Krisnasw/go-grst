package authtokenapi

import (
	"io"

	"google.golang.org/grpc/codes"
)

type AuthTokenClient interface {
	ValidateToken(reqBody io.Reader) (*SuccessResp, *ErrorResp)
}

type SuccessResp struct {
	Status      int                    `json:"http_status"`
	ProcessTime string                 `json:"process_time"`
	Data        map[string]interface{} `json:"data"`
}

type ErrorResp struct {
	HTTPStatus  int           `json:"http_status"`
	GRPCStatus  codes.Code    `json:"grpc_status"`
	Code        int           `json:"code"`
	Message     string        `json:"message"`
	OtherErrors []ErrorDetail `json:"other_errors"`
}

type ErrorDetail struct {
	Code    int    `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}
