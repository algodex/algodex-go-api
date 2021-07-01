// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account client
//
// Command:
// $ goa gen algodexidx/design

package account

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "account" service client.
type Client struct {
	AddEndpoint  goa.Endpoint
	GetEndpoint  goa.Endpoint
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "account" service client given the endpoints.
func NewClient(add, get, list goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:  add,
		GetEndpoint:  get,
		ListEndpoint: list,
	}
}

// Add calls the "add" endpoint of the "account" service.
func (c *Client) Add(ctx context.Context, p string) (err error) {
	_, err = c.AddEndpoint(ctx, p)
	return
}

// Get calls the "get" endpoint of the "account" service.
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *Account, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Account), nil
}

// List calls the "list" endpoint of the "account" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res TrackedAccountCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(TrackedAccountCollection), nil
}
