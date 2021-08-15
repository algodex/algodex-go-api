package backend

import (
	"context"
	"net/http"

	"goa.design/goa/v3/middleware"
)

type (
	// private type used to define context keys
	ctxKey int
)

const (
	// RemoteIPKey is the request context key used to store the remote IP address set by the SetRemoteIP middleware.
	RemoteIPKey ctxKey = iota + 1
)

func SetRemoteIP(options ...middleware.RequestIDOption) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), RemoteIPKey, r.RemoteAddr)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetRemoteIP returns the remote ip set into our context
func GetRemoteIP(ctx context.Context) string {
	if i := ctx.Value(RemoteIPKey); i != nil {
		return i.(string)
	}
	return ""
}
