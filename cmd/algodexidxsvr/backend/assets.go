package backend

import (
	"context"
	"sync"
)

// algoID is a synthetic 'ASA' ID used to represent an accounts' ALGO balance.
const algoID = 1

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

type assetMap struct {
	sync.RWMutex
	assets map[uint64]*AssetInformation
}

// Global asset cache
var assetCache = assetMap{assets: map[uint64]*AssetInformation{}}

// Returns existing (static) asset informaton for the specified ID, or loads
// the asset info from the chain, caching it for later fetches.
func (a *assetMap) Fetch(ctx context.Context, assetID uint64) *AssetInformation {
	a.RLock()
	if asset, found := a.assets[assetID]; found {
		a.RUnlock()
		return asset
	}
	a.RUnlock()

	assetData, err := algoClient.GetAssetByID(assetID).Do(ctx)
	if err != nil {
		// TODO: because asset data isn't returned as distinct data - we have no choice but to just return blank data
		// need logger instance as well to log error..... punt to log

		// Return 'blank' info for now - will refetch again on next try
		return &AssetInformation{}
	}
	assetInfo := &AssetInformation{
		Deleted:      assetData.Deleted,
		Decimals:     assetData.Params.Decimals,
		MetadataHash: assetData.Params.MetadataHash,
		Name:         assetData.Params.Name,
		UnitName:     assetData.Params.UnitName,
		Url:          assetData.Params.Url,
	}

	a.Lock()
	a.assets[assetID] = assetInfo
	a.Unlock()

	return assetInfo
}
