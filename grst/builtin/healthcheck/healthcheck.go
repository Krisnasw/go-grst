package healthcheck

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type Healthcheck interface {
	HealthcheckHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string)
	GetEndpointURL() string
}

type healthcheck struct {
	conn        *grpc.ClientConn
	endpointURL string
}

func New(conn *grpc.ClientConn, endpointURL string) Healthcheck {
	if endpointURL == "" {
		endpointURL = "_status"
	}
	return &healthcheck{conn, endpointURL}
}

func (h *healthcheck) HealthcheckHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Content-Type", "text/plain")
	if s := h.conn.GetState(); s != connectivity.Ready {
		http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
		return
	}
	w.Write([]byte("ok"))
}

func (h *healthcheck) GetEndpointURL() string {
	return h.endpointURL
}
