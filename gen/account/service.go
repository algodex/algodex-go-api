// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account service
//
// Command:
// $ goa gen algodexidx/design

package account

import (
	accountviews "algodexidx/gen/account/views"
	"context"
)

// The account service specifies which Algorand accounts to track
type Service interface {
	// Add Algorand account to track
	Add(context.Context, string) (err error)
	// Get specific account
	Get(context.Context, *GetPayload) (res *Account, err error)
	// List all tracked accounts
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "full"
	List(context.Context, *ListPayload) (res TrackedAccountCollection, view string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "account"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"add", "get", "list"}

// GetPayload is the payload type of the account service get method.
type GetPayload struct {
	// Public Account address
	Address string
}

// Account is the result type of the account service get method.
type Account struct {
	// Public Account address
	Address string
	// Opted-in ASA IDs
	Assets []uint64
}

// ListPayload is the payload type of the account service list method.
type ListPayload struct {
	// View to render
	View *string
}

// TrackedAccountCollection is the result type of the account service list
// method.
type TrackedAccountCollection []*TrackedAccount

// A TrackedAccount is an Account returned by the indexer
type TrackedAccount struct {
	// Public Account address
	Address string
	// Opted-in ASA IDs
	Assets []uint64
}

// NewTrackedAccountCollection initializes result type TrackedAccountCollection
// from viewed result type TrackedAccountCollection.
func NewTrackedAccountCollection(vres accountviews.TrackedAccountCollection) TrackedAccountCollection {
	var res TrackedAccountCollection
	switch vres.View {
	case "default", "":
		res = newTrackedAccountCollection(vres.Projected)
	case "full":
		res = newTrackedAccountCollectionFull(vres.Projected)
	}
	return res
}

// NewViewedTrackedAccountCollection initializes viewed result type
// TrackedAccountCollection from result type TrackedAccountCollection using the
// given view.
func NewViewedTrackedAccountCollection(res TrackedAccountCollection, view string) accountviews.TrackedAccountCollection {
	var vres accountviews.TrackedAccountCollection
	switch view {
	case "default", "":
		p := newTrackedAccountCollectionView(res)
		vres = accountviews.TrackedAccountCollection{Projected: p, View: "default"}
	case "full":
		p := newTrackedAccountCollectionViewFull(res)
		vres = accountviews.TrackedAccountCollection{Projected: p, View: "full"}
	}
	return vres
}

// newTrackedAccountCollection converts projected type TrackedAccountCollection
// to service type TrackedAccountCollection.
func newTrackedAccountCollection(vres accountviews.TrackedAccountCollectionView) TrackedAccountCollection {
	res := make(TrackedAccountCollection, len(vres))
	for i, n := range vres {
		res[i] = newTrackedAccount(n)
	}
	return res
}

// newTrackedAccountCollectionFull converts projected type
// TrackedAccountCollection to service type TrackedAccountCollection.
func newTrackedAccountCollectionFull(vres accountviews.TrackedAccountCollectionView) TrackedAccountCollection {
	res := make(TrackedAccountCollection, len(vres))
	for i, n := range vres {
		res[i] = newTrackedAccountFull(n)
	}
	return res
}

// newTrackedAccountCollectionView projects result type
// TrackedAccountCollection to projected type TrackedAccountCollectionView
// using the "default" view.
func newTrackedAccountCollectionView(res TrackedAccountCollection) accountviews.TrackedAccountCollectionView {
	vres := make(accountviews.TrackedAccountCollectionView, len(res))
	for i, n := range res {
		vres[i] = newTrackedAccountView(n)
	}
	return vres
}

// newTrackedAccountCollectionViewFull projects result type
// TrackedAccountCollection to projected type TrackedAccountCollectionView
// using the "full" view.
func newTrackedAccountCollectionViewFull(res TrackedAccountCollection) accountviews.TrackedAccountCollectionView {
	vres := make(accountviews.TrackedAccountCollectionView, len(res))
	for i, n := range res {
		vres[i] = newTrackedAccountViewFull(n)
	}
	return vres
}

// newTrackedAccount converts projected type TrackedAccount to service type
// TrackedAccount.
func newTrackedAccount(vres *accountviews.TrackedAccountView) *TrackedAccount {
	res := &TrackedAccount{}
	if vres.Address != nil {
		res.Address = *vres.Address
	}
	return res
}

// newTrackedAccountFull converts projected type TrackedAccount to service type
// TrackedAccount.
func newTrackedAccountFull(vres *accountviews.TrackedAccountView) *TrackedAccount {
	res := &TrackedAccount{}
	if vres.Address != nil {
		res.Address = *vres.Address
	}
	if vres.Assets != nil {
		res.Assets = make([]uint64, len(vres.Assets))
		for i, val := range vres.Assets {
			res.Assets[i] = val
		}
	}
	return res
}

// newTrackedAccountView projects result type TrackedAccount to projected type
// TrackedAccountView using the "default" view.
func newTrackedAccountView(res *TrackedAccount) *accountviews.TrackedAccountView {
	vres := &accountviews.TrackedAccountView{
		Address: &res.Address,
	}
	return vres
}

// newTrackedAccountViewFull projects result type TrackedAccount to projected
// type TrackedAccountView using the "full" view.
func newTrackedAccountViewFull(res *TrackedAccount) *accountviews.TrackedAccountView {
	vres := &accountviews.TrackedAccountView{
		Address: &res.Address,
	}
	if res.Assets != nil {
		vres.Assets = make([]uint64, len(res.Assets))
		for i, val := range res.Assets {
			vres.Assets[i] = val
		}
	}
	return vres
}
