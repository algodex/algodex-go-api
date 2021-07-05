// Code generated by goa v3.4.3, DO NOT EDIT.
//
// inspect HTTP server encoders and decoders
//
// Command:
// $ goa gen algodexidx/design

package server

import (
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUnpackResponse returns an encoder for responses returned by the
// inspect unpack endpoint.
func EncodeUnpackResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUnpackRequest returns a decoder for requests sent to the inspect
// unpack endpoint.
func DecodeUnpackRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UnpackRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewUnpackPayload(&body)

		return payload, nil
	}
}
