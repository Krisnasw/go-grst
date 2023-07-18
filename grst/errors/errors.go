package errors

import (
	"bytes"
	"context"
	fmt "fmt"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var jsonpbmarshal = &jsonpb.Marshaler{EmitDefaults: true}

func New(httpStatus int32, grpcStatus codes.Code, code int32, message string, otherErrors ...*ErrorDetail) error {
	if grpcStatus == codes.OK || httpStatus < 400 {
		return nil
	}

	gre := &Error{HTTPStatus: httpStatus, GRPCStatus: int32(grpcStatus), Code: code, Message: message, OtherErrors: otherErrors}
	st := status.New(codes.Code(gre.GRPCStatus), gre.Message)
	ds, errWD := st.WithDetails(gre)
	if errWD != nil {
		return st.Err()
	}
	return ds.Err()
}

func NewFromError(err error) (*Error, error) {
	s, ok := status.FromError(err)
	if !ok {
		return errConvert, fmt.Errorf(errConvert.Message)
	}
	statusDetails := s.Details()
	if len(statusDetails) <= 0 {
		errNoDetailsWithAdditionalMsg := &Error{
			HTTPStatus:  errNoDetails.HTTPStatus,
			GRPCStatus:  errNoDetails.GRPCStatus,
			Code:        errNoDetails.Code,
			Message:     fmt.Sprintf("%s: %s", errNoDetails.Message, s.Message()),
			OtherErrors: errNoDetails.OtherErrors,
		}
		return errNoDetailsWithAdditionalMsg, fmt.Errorf("%s: %s", errNoDetails.Message, err.Error())
	}

	var grpcRestErr *Error
	if customErrDetail, ok := statusDetails[0].(*Error); ok {
		grpcRestErr = customErrDetail
	} else {
		grpcRestErr = nil
	}
	if grpcRestErr == nil {
		return errInvalidFormat, fmt.Errorf(errInvalidFormat.Message)
	}
	return grpcRestErr, nil
}

func NewGeneralError500(errorMessage string) error {
	return New(http.StatusInternalServerError, codes.Internal, 999999, "General Error 500: "+errorMessage)
}
func NewGeneralError400(errorMessage string) error {
	return New(http.StatusBadRequest, codes.InvalidArgument, 999999, "General Error 400: "+errorMessage)
}

func CheckErrorByCode(err error, code int32) bool {
	grpcRestErr, errNew := NewFromError(err)
	if errNew != nil {
		return false
	}
	return grpcRestErr.Code == code
}

func (gre *Error) toJson() []byte {
	if gre == nil {
		gre = &Error{HTTPStatus: http.StatusOK, GRPCStatus: int32(codes.OK), Code: 200, Message: "OK"}
	}

	var res bytes.Buffer
	if err := jsonpbmarshal.Marshal(&res, gre); err != nil {
		return []byte(fmt.Sprintf(`
		{
			"http_status" : %d,
			"grpc_status" : %d,
			"code" : %d,
			"message" : "%s",
			"other_errors": []
		}
		`, gre.HTTPStatus, gre.GRPCStatus, gre.Code, gre.Message))
	}
	return res.Bytes()
}

var errConvert = &Error{HTTPStatus: http.StatusInternalServerError, GRPCStatus: int32(codes.Internal), Code: 99999, Message: "Error converting fault into grpc error", OtherErrors: []*ErrorDetail{}}
var errNoDetails = &Error{HTTPStatus: http.StatusInternalServerError, GRPCStatus: int32(codes.Internal), Code: 99998, Message: "Error contains no detail", OtherErrors: []*ErrorDetail{}}
var errInvalidFormat = &Error{HTTPStatus: http.StatusInternalServerError, GRPCStatus: int32(codes.Internal), Code: 99997, Message: "Error Invalid format error", OtherErrors: []*ErrorDetail{}}
var errMarshall = &Error{HTTPStatus: http.StatusInternalServerError, GRPCStatus: int32(codes.Internal), Code: 99996, Message: "Error marshall response error", OtherErrors: []*ErrorDetail{}}

func HTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	s := status.Convert(err)
	pb := s.Proto()
	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")
	w.Header().Set("Content-Type", marshaler.ContentType(pb))

	grpcRestErr, err := NewFromError(err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(grpcRestErr.toJson())
		return
	}

	buf, err := marshaler.Marshal(grpcRestErr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errMarshall.toJson())
		return
	}

	w.WriteHeader(int(grpcRestErr.HTTPStatus))
	if _, err := w.Write(buf); err != nil {
		logrus.Errorln("[Error] Failed write to response: %v", err)
	}

	return
}
