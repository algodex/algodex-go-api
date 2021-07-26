package backend

import (
	"context"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

type Holding struct {
	AssetID uint64
	Amount  uint64
	Info    AssetInformation
}
type holdingsMap map[uint64]*Holding

func getAccountHoldings(ctx context.Context, account string) (holdingsMap, error) {
	info, err := accountInformation(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("get account holdings:%s : %w", account, err)
	}
	holdings := make(holdingsMap, len(info.Assets)+1)
	holdings[algoID] = &Holding{
		AssetID: algoID,
		Amount:  info.Amount,
		Info: AssetInformation{
			Deleted:  false,
			Decimals: 6,
			Name:     "ALGO",
			UnitName: "ALGO",
		},
	}
	for _, asset := range info.Assets {
		assetInfo := assetCache.Fetch(ctx, asset.AssetId)
		if assetInfo == nil {
			continue
		}
		holdings[asset.AssetId] = &Holding{
			AssetID: asset.AssetId,
			Amount:  asset.Amount,
			Info:    *assetInfo,
		}
	}
	return holdings, nil
}

func accountInformation(ctx context.Context, account string) (models.Account, error) {
	accountInfo, err := algoClient.AccountInformation(account).Do(ctx)
	if err != nil {
		return models.Account{}, fmt.Errorf("fetching account info for:%s : %w", account, err)
	}
	return accountInfo, nil
}
