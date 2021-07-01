// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account HTTP server types
//
// Command:
// $ goa gen algodexidx/design

package server

import (
	account "algodexidx/gen/account"
	accountviews "algodexidx/gen/account/views"
)

// GetResponseBody is the type of the "account" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
	// Opted-in ASA IDs
	Assets []uint64 `form:"assets" json:"assets" xml:"assets"`
}

// TrackedAccountResponseCollection is the type of the "account" service "list"
// endpoint HTTP response body.
type TrackedAccountResponseCollection []*TrackedAccountResponse

// TrackedAccountResponseFullCollection is the type of the "account" service
// "list" endpoint HTTP response body.
type TrackedAccountResponseFullCollection []*TrackedAccountResponseFull

// TrackedAccountResponse is used to define fields on response body types.
type TrackedAccountResponse struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
}

// TrackedAccountResponseFull is used to define fields on response body types.
type TrackedAccountResponseFull struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
	// Opted-in ASA IDs
	Assets []uint64 `form:"assets" json:"assets" xml:"assets"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "account" service.
func NewGetResponseBody(res *account.Account) *GetResponseBody {
	body := &GetResponseBody{
		Address: res.Address,
	}
	if res.Assets != nil {
		body.Assets = make([]uint64, len(res.Assets))
		for i, val := range res.Assets {
			body.Assets[i] = val
		}
	}
	return body
}

// NewTrackedAccountResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "account" service.
func NewTrackedAccountResponseCollection(res accountviews.TrackedAccountCollectionView) TrackedAccountResponseCollection {
	body := make([]*TrackedAccountResponse, len(res))
	for i, val := range res {
		body[i] = marshalAccountviewsTrackedAccountViewToTrackedAccountResponse(val)
	}
	return body
}

// NewTrackedAccountResponseFullCollection builds the HTTP response body from
// the result of the "list" endpoint of the "account" service.
func NewTrackedAccountResponseFullCollection(res accountviews.TrackedAccountCollectionView) TrackedAccountResponseFullCollection {
	body := make([]*TrackedAccountResponseFull, len(res))
	for i, val := range res {
		body[i] = marshalAccountviewsTrackedAccountViewToTrackedAccountResponseFull(val)
	}
	return body
}

// NewGetPayload builds a account service get endpoint payload.
func NewGetPayload(address string) *account.GetPayload {
	v := &account.GetPayload{}
	v.Address = address

	return v
}

// NewListPayload builds a account service list endpoint payload.
func NewListPayload(view *string) *account.ListPayload {
	v := &account.ListPayload{}
	v.View = view

	return v
}
