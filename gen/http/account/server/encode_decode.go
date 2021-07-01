// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account HTTP server encoders and decoders
//
// Command:
// $ goa gen algodexidx/design

package server

import (
	account "algodexidx/gen/account"
	accountviews "algodexidx/gen/account/views"
	"context"
	"io"
	"net/http"
	"unicode/utf8"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddResponse returns an encoder for responses returned by the account
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeAddRequest returns a decoder for requests sent to the account add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body string
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := body

		return payload, nil
	}
}

// EncodeGetResponse returns an encoder for responses returned by the account
// get endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.Account)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the account get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			address string
			err     error

			params = mux.Vars(r)
		)
		address = params["address"]
		if utf8.RuneCountInString(address) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("address", address, utf8.RuneCountInString(address), 58, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(address)

		return payload, nil
	}
}

// EncodeListResponse returns an encoder for responses returned by the account
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(accountviews.TrackedAccountCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewTrackedAccountResponseCollection(res.Projected)
		case "full":
			body = NewTrackedAccountResponseFullCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the account list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			view *string
			err  error
		)
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "full") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "full"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(view)

		return payload, nil
	}
}

// marshalAccountviewsTrackedAccountViewToTrackedAccountResponse builds a value
// of type *TrackedAccountResponse from a value of type
// *accountviews.TrackedAccountView.
func marshalAccountviewsTrackedAccountViewToTrackedAccountResponse(v *accountviews.TrackedAccountView) *TrackedAccountResponse {
	res := &TrackedAccountResponse{
		Address: *v.Address,
	}

	return res
}

// marshalAccountviewsTrackedAccountViewToTrackedAccountResponseFull builds a
// value of type *TrackedAccountResponseFull from a value of type
// *accountviews.TrackedAccountView.
func marshalAccountviewsTrackedAccountViewToTrackedAccountResponseFull(v *accountviews.TrackedAccountView) *TrackedAccountResponseFull {
	res := &TrackedAccountResponseFull{
		Address: *v.Address,
	}
	if v.Assets != nil {
		res.Assets = make([]uint64, len(v.Assets))
		for i, val := range v.Assets {
			res.Assets[i] = val
		}
	}

	return res
}
