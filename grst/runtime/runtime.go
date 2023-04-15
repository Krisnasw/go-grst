package runtime

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/krisnasw/go-grst/grst/errors"
	"google.golang.org/protobuf/proto"
)

var jsonpbmarshal = &jsonpb.Marshaler{EmitDefaults: true}

func ForwardResponseMessage(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// var buf bytes.Buffer
	buf, err := marshaler.Marshal(resp)
	if err != nil {
		errors.HTTPError(ctx, mux, marshaler, w, req, err)
		return
	}
	var data interface{}
	// log if something goes wrong with unmarshalling response
	if err := json.Unmarshal(buf, &data); err != nil {
		errors.HTTPError(ctx, mux, marshaler, w, req, err)
		return
	}

	// Parse Start Time
	latency := "0ms"
	starttime, errParse := time.Parse(time.RFC3339Nano, req.Header.Get("grst.starttime"))
	if errParse == nil {
		latency = time.Since(starttime).String()
	}

	// Handling file
	is_file, filename, filecontent := is_file_download_resp(buf)
	if resp != nil && is_file {
		w.Header().Set("Content-Description", "File Transfer")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(filecontent)
		return
	}

	// template key value for REST response
	formattedResponse := &ResponseSuccess{
		Data:        resp,
		ProcessTime: latency,
		HTTPStatus:  http.StatusOK,
	}

	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, formattedResponse, opts...)
}

var file_type = map[string]bool{
	"image/gif":                 true,
	"image/png":                 true,
	"image/jpg":                 true,
	"text/csv":                  true,
	"text/plain; charset=utf-8": true,
	"text/xml; charset=utf-8":   true,
	"application/zip":           true,
}

func is_file_download_resp(buf []byte) (is_file bool, filename string, filecontent []byte) {
	is_file, filename, filecontent = false, "", []byte("")

	var data_map map[string]string
	if err := json.Unmarshal(buf, &data_map); err != nil {
		return
	}

	for k, v := range data_map {
		if strings.ToLower(k) == "filecontent" {
			// validate base64 is correct
			c, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				return
			}

			filecontent = c
			content_type := http.DetectContentType(c)
			if v, ok := file_type[content_type]; ok {
				is_file = v
			}
		}
		if strings.ToLower(k) == "filename" {
			filename = v
		}
	}
	return
}
