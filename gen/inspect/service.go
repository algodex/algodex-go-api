// Code generated by goa v3.4.3, DO NOT EDIT.
//
// inspect service
//
// Command:
// $ goa gen algodexidx/design

package inspect

import (
	"context"
)

// The inspect service provides msgpack decoding services
type Service interface {
	// Unpack a msgpack body (base64 encoded) returning 'goal clerk inspect' output
	Unpack(context.Context, *UnpackPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "inspect"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"unpack"}

// UnpackPayload is the payload type of the inspect service unpack method.
type UnpackPayload struct {
	Msgpack *string
}
