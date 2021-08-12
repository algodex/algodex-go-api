package backend

import (
	"encoding/json"
)

// algoID is a synthetic 'ASA' ID used to represent an accounts' ALGO balance.
const algoID = 1

// Account is information we track for every Algorand account, tracking
// all holdings and asset information.
// Persistence (in-memory cache w/ fetch from persistence layer on-demand)
type Account struct {
	// Public Account address
	Address string
	// Round which the information was updated for
	Round uint64
	// Opted-in ASA / balance information
	Holdings map[uint64]*Holding
}

// AccountFromJSON converts from marshaled value to Account struct (if valid)
func AccountFromJSON(b []byte) (*Account, error) {
	var retAccount *Account
	err := json.Unmarshal(b, &retAccount)
	return retAccount, err
}

// ToJSON marshals to json for persistence layers
func (a *Account) ToJSON() ([]byte, error) {
	return json.Marshal(*a)
}

// Holding is just the basic asset-id / amount we track per
// account.  Information on each asset-id we fetch independently.
type Holding struct {
	AssetID uint64
	Amount  uint64
	Info    AssetInformation
}

// AssetInformation contains information about an Algorand ASA, such
// as its name, unit name, decimals, etc.
type AssetInformation struct {
	Deleted      bool   // Whether the asset still 'exists'
	Decimals     uint64 // Decimals of asset - 2 decimals means amount of 100 is 1.00
	MetadataHash []byte // dependent on asset use - may play into NFTs
	Name         string
	UnitName     string
	Url          string
}

// AssetInformationFromJSON converts from marshaled value to AssetInformation struct (if valid)
func AssetInformationFromJSON(b []byte) (*AssetInformation, error) {
	var assetInfo *AssetInformation
	err := json.Unmarshal(b, &assetInfo)
	return assetInfo, err
}

// ToJSON marshals to json for persistence layers
func (a *AssetInformation) ToJSON() ([]byte, error) {
	return json.Marshal(*a)
}
