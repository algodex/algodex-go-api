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
	"strings"
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
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddPayload(&body)

		return payload, nil
	}
}

// EncodeAddError returns an encoder for errors returned by the add account
// endpoint.
func EncodeAddError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewAddAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the
// account delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the account
// delete endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			address []string
			err     error

			params = mux.Vars(r)
		)
		{
			addressRaw := params["address"]
			addressRawSlice := strings.Split(addressRaw, ",")
			address = make([]string, len(addressRawSlice))
			for i, rv := range addressRawSlice {
				address[i] = rv
			}
		}
		for _, e := range address {
			err = goa.MergeErrors(err, goa.ValidatePattern("address[*]", e, "^[A-Z2-7]{58}$"))
			if utf8.RuneCountInString(e) < 58 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("address[*]", e, utf8.RuneCountInString(e), 58, true))
			}
			if utf8.RuneCountInString(e) > 58 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("address[*]", e, utf8.RuneCountInString(e), 58, false))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(address)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete
// account endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteAllResponse returns an encoder for responses returned by the
// account deleteAll endpoint.
func EncodeDeleteAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// EncodeDeleteAllError returns an encoder for errors returned by the deleteAll
// account endpoint.
func EncodeDeleteAllError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteAllAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
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
		err = goa.MergeErrors(err, goa.ValidatePattern("address", address, "^[A-Z2-7]{58}$"))
		if utf8.RuneCountInString(address) < 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("address", address, utf8.RuneCountInString(address), 58, true))
		}
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

// EncodeGetError returns an encoder for errors returned by the get account
// endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetMultipleResponse returns an encoder for responses returned by the
// account getMultiple endpoint.
func EncodeGetMultipleResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*account.Account)
		enc := encoder(ctx, w)
		body := NewGetMultipleResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetMultipleRequest returns a decoder for requests sent to the account
// getMultiple endpoint.
func DecodeGetMultipleRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body GetMultipleRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateGetMultipleRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewGetMultiplePayload(&body)

		return payload, nil
	}
}

// EncodeGetMultipleError returns an encoder for errors returned by the
// getMultiple account endpoint.
func EncodeGetMultipleError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetMultipleAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
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

// EncodeListError returns an encoder for errors returned by the list account
// endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeIsWatchedResponse returns an encoder for responses returned by the
// account isWatched endpoint.
func EncodeIsWatchedResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIsWatchedRequest returns a decoder for requests sent to the account
// isWatched endpoint.
func DecodeIsWatchedRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body IsWatchedRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateIsWatchedRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewIsWatchedPayload(&body)

		return payload, nil
	}
}

// EncodeIsWatchedError returns an encoder for errors returned by the isWatched
// account endpoint.
func EncodeIsWatchedError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "access_denied":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewIsWatchedAccessDeniedResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalAccountHoldingToHoldingResponseBody builds a value of type
// *HoldingResponseBody from a value of type *account.Holding.
func marshalAccountHoldingToHoldingResponseBody(v *account.Holding) *HoldingResponseBody {
	res := &HoldingResponseBody{
		Asset:        v.Asset,
		Amount:       v.Amount,
		Decimals:     v.Decimals,
		MetadataHash: v.MetadataHash,
		Name:         v.Name,
		UnitName:     v.UnitName,
		URL:          v.URL,
	}

	return res
}

// marshalAccountAccountToAccountResponse builds a value of type
// *AccountResponse from a value of type *account.Account.
func marshalAccountAccountToAccountResponse(v *account.Account) *AccountResponse {
	res := &AccountResponse{
		Address: v.Address,
		Round:   v.Round,
	}
	if v.Holdings != nil {
		res.Holdings = make(map[string]*HoldingResponse, len(v.Holdings))
		for key, val := range v.Holdings {
			tk := key
			res.Holdings[tk] = marshalAccountHoldingToHoldingResponse(val)
		}
	}

	return res
}

// marshalAccountHoldingToHoldingResponse builds a value of type
// *HoldingResponse from a value of type *account.Holding.
func marshalAccountHoldingToHoldingResponse(v *account.Holding) *HoldingResponse {
	res := &HoldingResponse{
		Asset:        v.Asset,
		Amount:       v.Amount,
		Decimals:     v.Decimals,
		MetadataHash: v.MetadataHash,
		Name:         v.Name,
		UnitName:     v.UnitName,
		URL:          v.URL,
	}

	return res
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
		Round:   *v.Round,
	}
	if v.Holdings != nil {
		res.Holdings = make(map[string]*HoldingResponse, len(v.Holdings))
		for key, val := range v.Holdings {
			tk := key
			res.Holdings[tk] = marshalAccountviewsHoldingViewToHoldingResponse(val)
		}
	}

	return res
}

// marshalAccountviewsHoldingViewToHoldingResponse builds a value of type
// *HoldingResponse from a value of type *accountviews.HoldingView.
func marshalAccountviewsHoldingViewToHoldingResponse(v *accountviews.HoldingView) *HoldingResponse {
	res := &HoldingResponse{
		Asset:        *v.Asset,
		Amount:       *v.Amount,
		Decimals:     *v.Decimals,
		MetadataHash: *v.MetadataHash,
		Name:         *v.Name,
		UnitName:     *v.UnitName,
		URL:          *v.URL,
	}

	return res
}
