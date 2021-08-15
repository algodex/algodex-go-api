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
	AddEndpoint       goa.Endpoint
	DeleteEndpoint    goa.Endpoint
	DeleteallEndpoint goa.Endpoint
	GetEndpoint       goa.Endpoint
	ListEndpoint      goa.Endpoint
	IswatchedEndpoint goa.Endpoint
}

// NewClient initializes a "account" service client given the endpoints.
func NewClient(add, delete_, deleteall, get, list, iswatched goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:       add,
		DeleteEndpoint:    delete_,
		DeleteallEndpoint: deleteall,
		GetEndpoint:       get,
		ListEndpoint:      list,
		IswatchedEndpoint: iswatched,
	}
}

// Add calls the "add" endpoint of the "account" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (err error) {
	_, err = c.AddEndpoint(ctx, p)
	return
}

// Delete calls the "delete" endpoint of the "account" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}

// Deleteall calls the "deleteall" endpoint of the "account" service.
func (c *Client) Deleteall(ctx context.Context) (err error) {
	_, err = c.DeleteallEndpoint(ctx, nil)
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

// Iswatched calls the "iswatched" endpoint of the "account" service.
func (c *Client) Iswatched(ctx context.Context, p *IswatchedPayload) (res []string, err error) {
	var ires interface{}
	ires, err = c.IswatchedEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]string), nil
}
