package backend

import "encoding/json"

// Account is information we track for every Algorand account, tracking
// all holdings and asset information.
// Persistence (in-memory cache w/ fetch from persistence layer on-demand)
type Account struct {
	// Public Account address
	Address string
	// Opted-in ASA information
	Holdings map[uint64]BaseHolding
}

// BaseHolding is just the basic asset-id / amount we track per
// account.  Information on each asset-id we fetch independently.
type BaseHolding struct {
	AssetID uint64
	Amount  uint64
}

func (a *Account) ToJSON() ([]byte, error) {
	return json.Marshal(*a)
}

func AccountFromJSON(b []byte) (*Account, error) {
	var retAccount *Account
	err := json.Unmarshal(b, retAccount)
	return retAccount, err
}
