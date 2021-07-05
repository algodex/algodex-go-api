// Code generated by goa v3.4.3, DO NOT EDIT.
//
// inspect endpoints
//
// Command:
// $ goa gen algodexidx/design

package inspect

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "inspect" service endpoints.
type Endpoints struct {
	Unpack goa.Endpoint
}

// NewEndpoints wraps the methods of the "inspect" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Unpack: NewUnpackEndpoint(s),
	}
}

// Use applies the given middleware to all the "inspect" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Unpack = m(e.Unpack)
}

// NewUnpackEndpoint returns an endpoint function that calls the method
// "unpack" of service "inspect".
func NewUnpackEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UnpackPayload)
		return nil, s.Unpack(ctx, p)
	}
}
