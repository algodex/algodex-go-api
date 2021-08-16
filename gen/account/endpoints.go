// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account endpoints
//
// Command:
// $ goa gen algodexidx/design

package account

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "account" service endpoints.
type Endpoints struct {
	Add         goa.Endpoint
	Delete      goa.Endpoint
	Deleteall   goa.Endpoint
	Get         goa.Endpoint
	GetMultiple goa.Endpoint
	List        goa.Endpoint
	Iswatched   goa.Endpoint
}

// NewEndpoints wraps the methods of the "account" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add:         NewAddEndpoint(s),
		Delete:      NewDeleteEndpoint(s),
		Deleteall:   NewDeleteallEndpoint(s),
		Get:         NewGetEndpoint(s),
		GetMultiple: NewGetMultipleEndpoint(s),
		List:        NewListEndpoint(s),
		Iswatched:   NewIswatchedEndpoint(s),
	}
}

// Use applies the given middleware to all the "account" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Delete = m(e.Delete)
	e.Deleteall = m(e.Deleteall)
	e.Get = m(e.Get)
	e.GetMultiple = m(e.GetMultiple)
	e.List = m(e.List)
	e.Iswatched = m(e.Iswatched)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "account".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		return nil, s.Add(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "account".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return nil, s.Delete(ctx, p)
	}
}

// NewDeleteallEndpoint returns an endpoint function that calls the method
// "deleteall" of service "account".
func NewDeleteallEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Deleteall(ctx)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "account".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetPayload)
		return s.Get(ctx, p)
	}
}

// NewGetMultipleEndpoint returns an endpoint function that calls the method
// "getMultiple" of service "account".
func NewGetMultipleEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetMultiplePayload)
		return s.GetMultiple(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "account".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		res, view, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTrackedAccountCollection(res, view)
		return vres, nil
	}
}

// NewIswatchedEndpoint returns an endpoint function that calls the method
// "iswatched" of service "account".
func NewIswatchedEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IswatchedPayload)
		return s.Iswatched(ctx, p)
	}
}
