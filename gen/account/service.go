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
	// Add Algorand account(s) to track
	Add(context.Context, *AddPayload) (err error)
	// Delete Algorand account(s) to track
	Delete(context.Context, *DeletePayload) (err error)
	// Get specific account
	Get(context.Context, *GetPayload) (res *Account, err error)
	// List all tracked accounts
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "full"
	List(context.Context, *ListPayload) (res TrackedAccountCollection, view string, err error)
	// Returns which of the passed accounts are currently being monitored
	Iswatched(context.Context, *IswatchedPayload) (res []string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "account"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"add", "delete", "get", "list", "iswatched"}

// AddPayload is the payload type of the account service add method.
type AddPayload struct {
	Address []string
}

// DeletePayload is the payload type of the account service delete method.
type DeletePayload struct {
	Address []string
}

// GetPayload is the payload type of the account service get method.
type GetPayload struct {
	// Public Account address
	Address string
}

// Account is the result type of the account service get method.
type Account struct {
	// Public Account address
	Address string
	// Round fetched
	Round uint64
	// Account Assets
	Holdings map[string]*Holding
}

// ListPayload is the payload type of the account service list method.
type ListPayload struct {
	// View to render
	View *string
}

// TrackedAccountCollection is the result type of the account service list
// method.
type TrackedAccountCollection []*TrackedAccount

// IswatchedPayload is the payload type of the account service iswatched method.
type IswatchedPayload struct {
	Address []string
}

// Holding defines an ASA Asset ID and its balance.  ID 1 represents ALGO
type Holding struct {
	// ASA ID (1 for ALGO)
	Asset uint64
	// Balance in asset base units
	Amount       uint64
	Decimals     uint64
	MetadataHash string
	Name         string
	UnitName     string
	URL          string
}

// A TrackedAccount is an Account returned by the indexer
type TrackedAccount struct {
	// Public Account address
	Address string
	// Round fetched
	Round uint64
	// Account Assets
	Holdings map[string]*Holding
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
	if vres.Round != nil {
		res.Round = *vres.Round
	}
	if vres.Holdings != nil {
		res.Holdings = make(map[string]*Holding, len(vres.Holdings))
		for key, val := range vres.Holdings {
			tk := key
			res.Holdings[tk] = transformAccountviewsHoldingViewToHolding(val)
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
		Round:   &res.Round,
	}
	if res.Holdings != nil {
		vres.Holdings = make(map[string]*accountviews.HoldingView, len(res.Holdings))
		for key, val := range res.Holdings {
			tk := key
			vres.Holdings[tk] = transformHoldingToAccountviewsHoldingView(val)
		}
	}
	return vres
}

// transformAccountviewsHoldingViewToHolding builds a value of type *Holding
// from a value of type *accountviews.HoldingView.
func transformAccountviewsHoldingViewToHolding(v *accountviews.HoldingView) *Holding {
	if v == nil {
		return nil
	}
	res := &Holding{
		Asset:        *v.Asset,
		Amount:       *v.Amount,
		Decimals:     *v.Decimals,
		MetadataHash: *v.MetadataHash,
		Name:         *v.Name,
		UnitName:     *v.UnitName,
		URL:          *v.URL,
	}

	return res
}

// transformHoldingToAccountviewsHoldingView builds a value of type
// *accountviews.HoldingView from a value of type *Holding.
func transformHoldingToAccountviewsHoldingView(v *Holding) *accountviews.HoldingView {
	res := &accountviews.HoldingView{
		Asset:        &v.Asset,
		Amount:       &v.Amount,
		Decimals:     &v.Decimals,
		MetadataHash: &v.MetadataHash,
		Name:         &v.Name,
		UnitName:     &v.UnitName,
		URL:          &v.URL,
	}

	return res
}
