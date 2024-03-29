// Code generated by goa v3.4.3, DO NOT EDIT.
//
// inspect HTTP server
//
// Command:
// $ goa gen algodexidx/design

package server

import (
	inspect "algodexidx/gen/inspect"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the inspect service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Unpack http.Handler
	CORS   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the inspect service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *inspect.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Unpack", "POST", "/inspect/unpack"},
			{"CORS", "OPTIONS", "/inspect/unpack"},
		},
		Unpack: NewUnpackHandler(e.Unpack, mux, decoder, encoder, errhandler, formatter),
		CORS:   NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "inspect" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Unpack = m(s.Unpack)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the inspect endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUnpackHandler(mux, h.Unpack)
	MountCORSHandler(mux, h.CORS)
}

// MountUnpackHandler configures the mux to serve the "inspect" service
// "unpack" endpoint.
func MountUnpackHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleInspectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/inspect/unpack", f)
}

// NewUnpackHandler creates a HTTP handler which loads the HTTP request and
// calls the "inspect" service "unpack" endpoint.
func NewUnpackHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUnpackRequest(mux, decoder)
		encodeResponse = EncodeUnpackResponse(encoder)
		encodeError    = EncodeUnpackError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "unpack")
		ctx = context.WithValue(ctx, goa.ServiceKey, "inspect")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service inspect.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleInspectOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/inspect/unpack", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleInspectOrigin applies the CORS response headers corresponding to the
// origin for the service inspect.
func HandleInspectOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "*")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
