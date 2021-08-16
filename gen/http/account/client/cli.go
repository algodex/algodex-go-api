// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account HTTP client CLI support package
//
// Command:
// $ goa gen algodexidx/design

package client

import (
	account "algodexidx/gen/account"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildAddPayload builds the payload for the account add endpoint from CLI
// flags.
func BuildAddPayload(accountAddBody string) (*account.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(accountAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": [\n         \"4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU\",\n         \"6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY\"\n      ]\n   }'")
		}
		if body.Address == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &account.AddPayload{}
	if body.Address != nil {
		v.Address = make([]string, len(body.Address))
		for i, val := range body.Address {
			v.Address[i] = val
		}
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the account delete endpoint from
// CLI flags.
func BuildDeletePayload(accountDeleteAddress string) (*account.DeletePayload, error) {
	var err error
	var address []string
	{
		err = json.Unmarshal([]byte(accountDeleteAddress), &address)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for address, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU\",\n      \"6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY\"\n   ]'")
		}
	}
	v := &account.DeletePayload{}
	v.Address = address

	return v, nil
}

// BuildGetPayload builds the payload for the account get endpoint from CLI
// flags.
func BuildGetPayload(accountGetAddress string) (*account.GetPayload, error) {
	var err error
	var address string
	{
		address = accountGetAddress
		if utf8.RuneCountInString(address) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("address", address, utf8.RuneCountInString(address), 58, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &account.GetPayload{}
	v.Address = address

	return v, nil
}

// BuildGetMultiplePayload builds the payload for the account getMultiple
// endpoint from CLI flags.
func BuildGetMultiplePayload(accountGetMultipleBody string) (*account.GetMultiplePayload, error) {
	var err error
	var body GetMultipleRequestBody
	{
		err = json.Unmarshal([]byte(accountGetMultipleBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": [\n         \"4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU\",\n         \"6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY\"\n      ]\n   }'")
		}
		if body.Address == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &account.GetMultiplePayload{}
	if body.Address != nil {
		v.Address = make([]string, len(body.Address))
		for i, val := range body.Address {
			v.Address[i] = val
		}
	}

	return v, nil
}

// BuildListPayload builds the payload for the account list endpoint from CLI
// flags.
func BuildListPayload(accountListView string) (*account.ListPayload, error) {
	var err error
	var view *string
	{
		if accountListView != "" {
			view = &accountListView
			if view != nil {
				if !(*view == "default" || *view == "full") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "full"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &account.ListPayload{}
	v.View = view

	return v, nil
}

// BuildIswatchedPayload builds the payload for the account iswatched endpoint
// from CLI flags.
func BuildIswatchedPayload(accountIswatchedBody string) (*account.IswatchedPayload, error) {
	var err error
	var body IswatchedRequestBody
	{
		err = json.Unmarshal([]byte(accountIswatchedBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": [\n         \"4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU\",\n         \"6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY\"\n      ]\n   }'")
		}
		if body.Address == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &account.IswatchedPayload{}
	if body.Address != nil {
		v.Address = make([]string, len(body.Address))
		for i, val := range body.Address {
			v.Address[i] = val
		}
	}

	return v, nil
}
