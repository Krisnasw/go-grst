package grst

import (
	"net/http"
	"strings"

	grst_context "github.com/krisnasw/go-grst/grst/context"
)

type CORSPolicy struct {
	allowedOrigin    map[string]bool
	allowAllOrigin   bool
	preflightHeaders []string
	preflightMethods []string
	withCredential   bool
}

func (c CORSPolicy) isAllowedOrigin(origin string) bool {
	if c.allowAllOrigin {
		return true
	} else if value, ok := c.allowedOrigin[origin]; ok {
		return value
	}
	return false
}
func (c CORSPolicy) getAllOrigins() []string {
	resp := []string{}
	for origin, ok := range c.allowedOrigin {
		if ok {
			resp = append(resp, origin)
		}
	}
	return resp
}

var DefaultCORSPreflightHeaders = []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization",
	grst_context.CONTEXT_CLIENT_APPNAME.String(),
	grst_context.CONTEXT_CLIENT_APPVERSION.String(),
	grst_context.CONTEXT_CLIENT_APPVERSIONCODE.String(),
	grst_context.CONTEXT_CLIENT_MANUFACTURER.String(),
	grst_context.CONTEXT_CLIENT_MODEL.String(),
	grst_context.CONTEXT_CLIENT_PLATFORM.String(),
	grst_context.CONTEXT_CLIENT_PLATFORMVERSION.String(),
	grst_context.CONTEXT_CLIENT_SDKVERSION.String(),
}
var DefaultCORSPreflightMethods = []string{"GET", "HEAD", "POST", "PUT", "DELETE"}

func wrapMuxWithCors(h http.Handler, corsPolicy CORSPolicy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if corsPolicy.allowAllOrigin || corsPolicy.isAllowedOrigin(origin) {
			// if origin := r.Header.Get("Origin"); origin != "" {
			if corsPolicy.allowAllOrigin {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			if corsPolicy.withCredential {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsPolicy.preflightHeaders, ","))
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsPolicy.preflightMethods, ","))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
