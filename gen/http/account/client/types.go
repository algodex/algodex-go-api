// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account HTTP client types
//
// Command:
// $ goa gen algodexidx/design

package client

import (
	account "algodexidx/gen/account"
	accountviews "algodexidx/gen/account/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "account" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	Address []string `form:"address" json:"address" xml:"address"`
}

// DeleteRequestBody is the type of the "account" service "delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	Address []string `form:"address" json:"address" xml:"address"`
}

// GetResponseBody is the type of the "account" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Public Account address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Round fetched
	Round *uint64 `form:"round,omitempty" json:"round,omitempty" xml:"round,omitempty"`
	// Account Assets
	Holdings map[string]*HoldingResponseBody `form:"holdings,omitempty" json:"holdings,omitempty" xml:"holdings,omitempty"`
}

// ListResponseBody is the type of the "account" service "list" endpoint HTTP
// response body.
type ListResponseBody []*TrackedAccountResponse

// HoldingResponseBody is used to define fields on response body types.
type HoldingResponseBody struct {
	// ASA ID (1 for ALGO)
	Asset *uint64 `form:"asset,omitempty" json:"asset,omitempty" xml:"asset,omitempty"`
	// Balance in asset base units
	Amount       *uint64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	Decimals     *uint64 `form:"decimals,omitempty" json:"decimals,omitempty" xml:"decimals,omitempty"`
	MetadataHash *string `form:"metadataHash,omitempty" json:"metadataHash,omitempty" xml:"metadataHash,omitempty"`
	Name         *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	UnitName     *string `form:"unitName,omitempty" json:"unitName,omitempty" xml:"unitName,omitempty"`
	URL          *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// TrackedAccountResponse is used to define fields on response body types.
type TrackedAccountResponse struct {
	// Public Account address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Round fetched
	Round *uint64 `form:"round,omitempty" json:"round,omitempty" xml:"round,omitempty"`
	// Account Assets
	Holdings map[string]*HoldingResponse `form:"holdings,omitempty" json:"holdings,omitempty" xml:"holdings,omitempty"`
}

// HoldingResponse is used to define fields on response body types.
type HoldingResponse struct {
	// ASA ID (1 for ALGO)
	Asset *uint64 `form:"asset,omitempty" json:"asset,omitempty" xml:"asset,omitempty"`
	// Balance in asset base units
	Amount       *uint64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	Decimals     *uint64 `form:"decimals,omitempty" json:"decimals,omitempty" xml:"decimals,omitempty"`
	MetadataHash *string `form:"metadataHash,omitempty" json:"metadataHash,omitempty" xml:"metadataHash,omitempty"`
	Name         *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	UnitName     *string `form:"unitName,omitempty" json:"unitName,omitempty" xml:"unitName,omitempty"`
	URL          *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "account" service.
func NewAddRequestBody(p *account.AddPayload) *AddRequestBody {
	body := &AddRequestBody{}
	if p.Address != nil {
		body.Address = make([]string, len(p.Address))
		for i, val := range p.Address {
			body.Address[i] = val
		}
	}
	return body
}

// NewDeleteRequestBody builds the HTTP request body from the payload of the
// "delete" endpoint of the "account" service.
func NewDeleteRequestBody(p *account.DeletePayload) *DeleteRequestBody {
	body := &DeleteRequestBody{}
	if p.Address != nil {
		body.Address = make([]string, len(p.Address))
		for i, val := range p.Address {
			body.Address[i] = val
		}
	}
	return body
}

// NewGetAccountOK builds a "account" service "get" endpoint result from a HTTP
// "OK" response.
func NewGetAccountOK(body *GetResponseBody) *account.Account {
	v := &account.Account{
		Address: *body.Address,
		Round:   *body.Round,
	}
	v.Holdings = make(map[string]*account.Holding, len(body.Holdings))
	for key, val := range body.Holdings {
		tk := key
		v.Holdings[tk] = unmarshalHoldingResponseBodyToAccountHolding(val)
	}

	return v
}

// NewListTrackedAccountCollectionOK builds a "account" service "list" endpoint
// result from a HTTP "OK" response.
func NewListTrackedAccountCollectionOK(body ListResponseBody) accountviews.TrackedAccountCollectionView {
	v := make([]*accountviews.TrackedAccountView, len(body))
	for i, val := range body {
		v[i] = unmarshalTrackedAccountResponseToAccountviewsTrackedAccountView(val)
	}

	return v
}

// ValidateGetResponseBody runs the validations defined on GetResponseBody
func ValidateGetResponseBody(body *GetResponseBody) (err error) {
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Round == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("round", "body"))
	}
	if body.Holdings == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("holdings", "body"))
	}
	if body.Address != nil {
		if utf8.RuneCountInString(*body.Address) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", *body.Address, utf8.RuneCountInString(*body.Address), 58, false))
		}
	}
	for _, v := range body.Holdings {
		if v != nil {
			if err2 := ValidateHoldingResponseBody(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateHoldingResponseBody runs the validations defined on
// HoldingResponseBody
func ValidateHoldingResponseBody(body *HoldingResponseBody) (err error) {
	if body.Asset == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("asset", "body"))
	}
	if body.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("amount", "body"))
	}
	if body.Decimals == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("decimals", "body"))
	}
	if body.MetadataHash == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("metadataHash", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.UnitName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("unitName", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	return
}

// ValidateTrackedAccountResponse runs the validations defined on
// TrackedAccountResponse
func ValidateTrackedAccountResponse(body *TrackedAccountResponse) (err error) {
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Round == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("round", "body"))
	}
	if body.Holdings == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("holdings", "body"))
	}
	if body.Address != nil {
		if utf8.RuneCountInString(*body.Address) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", *body.Address, utf8.RuneCountInString(*body.Address), 58, false))
		}
	}
	for _, v := range body.Holdings {
		if v != nil {
			if err2 := ValidateHoldingResponse(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateHoldingResponse runs the validations defined on HoldingResponse
func ValidateHoldingResponse(body *HoldingResponse) (err error) {
	if body.Asset == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("asset", "body"))
	}
	if body.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("amount", "body"))
	}
	if body.Decimals == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("decimals", "body"))
	}
	if body.MetadataHash == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("metadataHash", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.UnitName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("unitName", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	return
}
