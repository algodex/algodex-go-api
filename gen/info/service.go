// Code generated by goa v3.4.3, DO NOT EDIT.
//
// info service
//
// Command:
// $ goa gen algodexidx/design

package info

import (
	"context"
)

// The info service provides information on version data, etc.
type Service interface {
	// Returns version information for the service
	Version(context.Context) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "info"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"version"}
