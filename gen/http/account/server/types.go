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
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "account" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	Address []string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// GetMultipleRequestBody is the type of the "account" service "getMultiple"
// endpoint HTTP request body.
type GetMultipleRequestBody struct {
	Address []string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// IsWatchedRequestBody is the type of the "account" service "isWatched"
// endpoint HTTP request body.
type IsWatchedRequestBody struct {
	Address []string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// GetResponseBody is the type of the "account" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
	// Round fetched
	Round uint64 `form:"round" json:"round" xml:"round"`
	// Account Assets
	Holdings map[string]*HoldingResponseBody `form:"holdings" json:"holdings" xml:"holdings"`
}

// GetMultipleResponseBody is the type of the "account" service "getMultiple"
// endpoint HTTP response body.
type GetMultipleResponseBody []*AccountResponse

// TrackedAccountResponseCollection is the type of the "account" service "list"
// endpoint HTTP response body.
type TrackedAccountResponseCollection []*TrackedAccountResponse

// TrackedAccountResponseFullCollection is the type of the "account" service
// "list" endpoint HTTP response body.
type TrackedAccountResponseFullCollection []*TrackedAccountResponseFull

// AddAccessDeniedResponseBody is the type of the "account" service "add"
// endpoint HTTP response body for the "access_denied" error.
type AddAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteAccessDeniedResponseBody is the type of the "account" service "delete"
// endpoint HTTP response body for the "access_denied" error.
type DeleteAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteAllAccessDeniedResponseBody is the type of the "account" service
// "deleteAll" endpoint HTTP response body for the "access_denied" error.
type DeleteAllAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetAccessDeniedResponseBody is the type of the "account" service "get"
// endpoint HTTP response body for the "access_denied" error.
type GetAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetMultipleAccessDeniedResponseBody is the type of the "account" service
// "getMultiple" endpoint HTTP response body for the "access_denied" error.
type GetMultipleAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListAccessDeniedResponseBody is the type of the "account" service "list"
// endpoint HTTP response body for the "access_denied" error.
type ListAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// IsWatchedAccessDeniedResponseBody is the type of the "account" service
// "isWatched" endpoint HTTP response body for the "access_denied" error.
type IsWatchedAccessDeniedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// HoldingResponseBody is used to define fields on response body types.
type HoldingResponseBody struct {
	// ASA ID (1 for ALGO)
	Asset uint64 `form:"asset" json:"asset" xml:"asset"`
	// Balance in asset base units
	Amount       uint64 `form:"amount" json:"amount" xml:"amount"`
	Decimals     uint64 `form:"decimals" json:"decimals" xml:"decimals"`
	MetadataHash string `form:"metadataHash" json:"metadataHash" xml:"metadataHash"`
	Name         string `form:"name" json:"name" xml:"name"`
	UnitName     string `form:"unitName" json:"unitName" xml:"unitName"`
	URL          string `form:"url" json:"url" xml:"url"`
}

// AccountResponse is used to define fields on response body types.
type AccountResponse struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
	// Round fetched
	Round uint64 `form:"round" json:"round" xml:"round"`
	// Account Assets
	Holdings map[string]*HoldingResponse `form:"holdings" json:"holdings" xml:"holdings"`
}

// HoldingResponse is used to define fields on response body types.
type HoldingResponse struct {
	// ASA ID (1 for ALGO)
	Asset uint64 `form:"asset" json:"asset" xml:"asset"`
	// Balance in asset base units
	Amount       uint64 `form:"amount" json:"amount" xml:"amount"`
	Decimals     uint64 `form:"decimals" json:"decimals" xml:"decimals"`
	MetadataHash string `form:"metadataHash" json:"metadataHash" xml:"metadataHash"`
	Name         string `form:"name" json:"name" xml:"name"`
	UnitName     string `form:"unitName" json:"unitName" xml:"unitName"`
	URL          string `form:"url" json:"url" xml:"url"`
}

// TrackedAccountResponse is used to define fields on response body types.
type TrackedAccountResponse struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
}

// TrackedAccountResponseFull is used to define fields on response body types.
type TrackedAccountResponseFull struct {
	// Public Account address
	Address string `form:"address" json:"address" xml:"address"`
	// Round fetched
	Round uint64 `form:"round" json:"round" xml:"round"`
	// Account Assets
	Holdings map[string]*HoldingResponse `form:"holdings" json:"holdings" xml:"holdings"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "account" service.
func NewGetResponseBody(res *account.Account) *GetResponseBody {
	body := &GetResponseBody{
		Address: res.Address,
		Round:   res.Round,
	}
	if res.Holdings != nil {
		body.Holdings = make(map[string]*HoldingResponseBody, len(res.Holdings))
		for key, val := range res.Holdings {
			tk := key
			body.Holdings[tk] = marshalAccountHoldingToHoldingResponseBody(val)
		}
	}
	return body
}

// NewGetMultipleResponseBody builds the HTTP response body from the result of
// the "getMultiple" endpoint of the "account" service.
func NewGetMultipleResponseBody(res []*account.Account) GetMultipleResponseBody {
	body := make([]*AccountResponse, len(res))
	for i, val := range res {
		body[i] = marshalAccountAccountToAccountResponse(val)
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

// NewAddAccessDeniedResponseBody builds the HTTP response body from the result
// of the "add" endpoint of the "account" service.
func NewAddAccessDeniedResponseBody(res *goa.ServiceError) *AddAccessDeniedResponseBody {
	body := &AddAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteAccessDeniedResponseBody builds the HTTP response body from the
// result of the "delete" endpoint of the "account" service.
func NewDeleteAccessDeniedResponseBody(res *goa.ServiceError) *DeleteAccessDeniedResponseBody {
	body := &DeleteAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteAllAccessDeniedResponseBody builds the HTTP response body from the
// result of the "deleteAll" endpoint of the "account" service.
func NewDeleteAllAccessDeniedResponseBody(res *goa.ServiceError) *DeleteAllAccessDeniedResponseBody {
	body := &DeleteAllAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetAccessDeniedResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "account" service.
func NewGetAccessDeniedResponseBody(res *goa.ServiceError) *GetAccessDeniedResponseBody {
	body := &GetAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetMultipleAccessDeniedResponseBody builds the HTTP response body from
// the result of the "getMultiple" endpoint of the "account" service.
func NewGetMultipleAccessDeniedResponseBody(res *goa.ServiceError) *GetMultipleAccessDeniedResponseBody {
	body := &GetMultipleAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListAccessDeniedResponseBody builds the HTTP response body from the
// result of the "list" endpoint of the "account" service.
func NewListAccessDeniedResponseBody(res *goa.ServiceError) *ListAccessDeniedResponseBody {
	body := &ListAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewIsWatchedAccessDeniedResponseBody builds the HTTP response body from the
// result of the "isWatched" endpoint of the "account" service.
func NewIsWatchedAccessDeniedResponseBody(res *goa.ServiceError) *IsWatchedAccessDeniedResponseBody {
	body := &IsWatchedAccessDeniedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAddPayload builds a account service add endpoint payload.
func NewAddPayload(body *AddRequestBody) *account.AddPayload {
	v := &account.AddPayload{}
	v.Address = make([]string, len(body.Address))
	for i, val := range body.Address {
		v.Address[i] = val
	}

	return v
}

// NewDeletePayload builds a account service delete endpoint payload.
func NewDeletePayload(address []string) *account.DeletePayload {
	v := &account.DeletePayload{}
	v.Address = address

	return v
}

// NewGetPayload builds a account service get endpoint payload.
func NewGetPayload(address string) *account.GetPayload {
	v := &account.GetPayload{}
	v.Address = address

	return v
}

// NewGetMultiplePayload builds a account service getMultiple endpoint payload.
func NewGetMultiplePayload(body *GetMultipleRequestBody) *account.GetMultiplePayload {
	v := &account.GetMultiplePayload{}
	v.Address = make([]string, len(body.Address))
	for i, val := range body.Address {
		v.Address[i] = val
	}

	return v
}

// NewListPayload builds a account service list endpoint payload.
func NewListPayload(view *string) *account.ListPayload {
	v := &account.ListPayload{}
	v.View = view

	return v
}

// NewIsWatchedPayload builds a account service isWatched endpoint payload.
func NewIsWatchedPayload(body *IsWatchedRequestBody) *account.IsWatchedPayload {
	v := &account.IsWatchedPayload{}
	v.Address = make([]string, len(body.Address))
	for i, val := range body.Address {
		v.Address[i] = val
	}

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	for _, e := range body.Address {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.address[*]", e, "^[A-Z2-7]{58}$"))
		if utf8.RuneCountInString(e) < 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address[*]", e, utf8.RuneCountInString(e), 58, true))
		}
		if utf8.RuneCountInString(e) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address[*]", e, utf8.RuneCountInString(e), 58, false))
		}
	}
	return
}

// ValidateGetMultipleRequestBody runs the validations defined on
// GetMultipleRequestBody
func ValidateGetMultipleRequestBody(body *GetMultipleRequestBody) (err error) {
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	for _, e := range body.Address {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.address[*]", e, "^[A-Z2-7]{58}$"))
		if utf8.RuneCountInString(e) < 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address[*]", e, utf8.RuneCountInString(e), 58, true))
		}
		if utf8.RuneCountInString(e) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address[*]", e, utf8.RuneCountInString(e), 58, false))
		}
	}
	return
}

// ValidateIsWatchedRequestBody runs the validations defined on
// IsWatchedRequestBody
func ValidateIsWatchedRequestBody(body *IsWatchedRequestBody) (err error) {
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	for _, e := range body.Address {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.address[*]", e, "^[A-Z2-7]{58}$"))
		if utf8.RuneCountInString(e) < 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address[*]", e, utf8.RuneCountInString(e), 58, true))
		}
		if utf8.RuneCountInString(e) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address[*]", e, utf8.RuneCountInString(e), 58, false))
		}
	}
	return
}
