// Code generated by goa v3.4.3, DO NOT EDIT.
//
// HTTP request path constructors for the account service.
//
// Command:
// $ goa gen algodexidx/design

package server

import (
	"fmt"
	"net/url"
	"strings"
)

// AddAccountPath returns the URL path to the account service add HTTP endpoint.
func AddAccountPath() string {
	return "/account"
}

// DeleteAccountPath returns the URL path to the account service delete HTTP endpoint.
func DeleteAccountPath(address []string) string {
	addressSlice := make([]string, len(address))
	for i, v := range address {
		addressSlice[i] = url.QueryEscape(v)
	}
	return fmt.Sprintf("/account/%v", strings.Join(addressSlice, ","))
}

// DeleteallAccountPath returns the URL path to the account service deleteall HTTP endpoint.
func DeleteallAccountPath() string {
	return "/account/all"
}

// GetAccountPath returns the URL path to the account service get HTTP endpoint.
func GetAccountPath(address string) string {
	return fmt.Sprintf("/account/%v", address)
}

// ListAccountPath returns the URL path to the account service list HTTP endpoint.
func ListAccountPath() string {
	return "/account"
}

// IswatchedAccountPath returns the URL path to the account service iswatched HTTP endpoint.
func IswatchedAccountPath() string {
	return "/account/iswatched"
}
