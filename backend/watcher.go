package backend

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/getsentry/sentry-go"
)

// addressAndResult is what is sent to the background account updater goroutines channel and
// contains an address  to 'update' and an optional channel to signal once the account is confirmed updated.
type addressAndResult struct {
	address string
	result  chan<- *Account
	ctx     context.Context
}

func (a addressAndResult) String() string {
	return a.address
}

type accountUpdateChanType chan addressAndResult

type watcher struct {
	logger            *log.Logger
	algoClient        *algod.Client
	persist           Persistor
	accountUpdateChan accountUpdateChanType
}

func newWatcher(log *log.Logger, algoClient *algod.Client, persistor Persistor) *watcher {
	return &watcher{
		logger:     log,
		algoClient: algoClient,
		persist:    persistor,
	}
}

func (w *watcher) WatchAccounts(ctx context.Context, addresses ...string) error {
	err := w.persist.WatchAccounts(ctx, addresses...)
	if err != nil {
		return fmt.Errorf("error in WatchAccounts: %w", err)
	}
	for _, address := range addresses {
		// if there are a ton of accounts
		w.accountUpdateChan <- addressAndResult{address: address}
	}
	return nil
}

func (w *watcher) UnwatchAccounts(ctx context.Context, addresses ...string) error {
	err := w.persist.UnwatchAccounts(ctx, addresses...)
	if err != nil {
		return fmt.Errorf("error in UnwatchAccounts: %w", err)
	}
	return nil
}

func (w *watcher) GetAccount(ctx context.Context, address string) (*Account, error) {
	account, err := w.persist.GetAccount(ctx, address)
	if account == nil {
		// not cached yet - force update from chain and update cache
		account, err = w.updateAccountInfo(ctx, address)
		if err != nil {
			return nil, fmt.Errorf("error fetching un-cached account info, error:%w", err)
		}
	}
	return account, nil

}

func (w *watcher) GetAccounts(ctx context.Context) []*Account {
	accounts, err := w.persist.GetWatchedAccounts(ctx)
	if err != nil {
		return nil
	}
	retAccounts := make([]*Account, 0, len(accounts))
	for _, address := range accounts {
		account, err := w.GetAccount(ctx, address)
		if err != nil {
			w.logger.Printf("GetAccounts error fetching account address:%s, error:%s", address, err.Error())
			continue
		}
		retAccounts = append(retAccounts, account)
	}
	return retAccounts
}

func (w *watcher) start(ctx context.Context) {
	go func(localHub *sentry.Hub) {
		localHub.ConfigureScope(
			func(scope *sentry.Scope) {
				scope.SetTag("goroutine", "accountWatcher")
			},
		)
		w.accountWatcher(ctx)
	}(sentry.CurrentHub().Clone())
}

func (w *watcher) accountWatcher(ctx context.Context) {
	defer w.logger.Println("Exited accountWatcher")
	wg := sync.WaitGroup{}

	w.logger.Printf("Num cores:%d", runtime.NumCPU())
	w.accountUpdateChan = make(accountUpdateChanType, 1000)
	// Create parallel update workers - 4x core count
	for i := 0; i < runtime.NumCPU()*4; i++ {
		wg.Add(1)
		go func(localHub *sentry.Hub) {
			localHub.ConfigureScope(
				func(scope *sentry.Scope) {
					scope.SetTag("goroutine", fmt.Sprintf("accountUpdater:%d/%d", i+1, runtime.NumCPU()*4))
				},
			)
			w.accountUpdater(ctx, &wg)
		}(sentry.CurrentHub().Clone())
	}

	// Get last round we saw...
	startRound, _ := w.persist.GetLastRound(ctx)

	// Then all watched accounts.
	allWatched, err := w.persist.GetWatchedAccounts(ctx)
	if err != nil {
		w.logger.Panicf("error getting watched accounts: %v", err)
	}

	// Get current round from node...
	nodeStatus, err := w.algoClient.Status().Do(ctx)
	if err != nil {
		w.logger.Panicf("error node status: %v", err)
	}
	//
	// If we're farther away than 1000 blocks (in case we're running against a non-archival node) or farther away
	// from our last round than the number of watched accounts, then just start at the nodes' current round
	// and re-fetch all accounts explicitly.
	// If not - just start where we left off and update the accounts as we go.
	if (nodeStatus.LastRound-startRound) > 1000 || uint64(len(allWatched)) < (nodeStatus.LastRound-startRound) {
		startRound = nodeStatus.LastRound

		// queue update of ALL watched accounts
		w.logger.Printf("Starting far enough behind, just updating all accounts (%d total)", len(allWatched))
		go func() {
			for _, address := range allWatched {
				// if there are a ton of accounts
				w.accountUpdateChan <- addressAndResult{address: address}
			}
		}()
	}

	// Then we just watch for updates in each new block
	go func(localHub *sentry.Hub) {
		localHub.ConfigureScope(
			func(scope *sentry.Scope) {
				scope.SetTag("goroutine", "blockWatcher")
			},
		)
		w.blockWatcher(ctx, startRound)
	}(sentry.CurrentHub().Clone())

	// Wait until we're told to exit...
	<-ctx.Done()
	// Shut down the updater channel (which will terminate the accountUpdaters once they're caught up)
	close(w.accountUpdateChan)
	// Now wait until all the updaters finish...
	wg.Wait()
}

func (w *watcher) accountUpdater(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// Endlessly loop on accounts to update the assets until our channel is closed
	for accountUpdate := range w.accountUpdateChan {
		newctx := ctx
		if accountUpdate.ctx != nil {
			newctx = accountUpdate.ctx
		}
		span := sentry.StartSpan(
			newctx, "account-update",
			sentry.TransactionName(fmt.Sprintf("update: %s", accountUpdate)),
		)

		w.logger.Printf("Account holdings update, account:%s", accountUpdate)
		// fetching the account also updates the persistence layer for later cached fetch
		updatedAccount, err := w.updateAccountInfo(span.Context(), accountUpdate.address)
		if err != nil {
			w.logger.Printf("error fetching holdings for account:%v - will retry, error:%v", accountUpdate, err.Error())
			w.accountUpdateChan <- accountUpdate
			span.Finish()
			continue
		}
		if accountUpdate.result != nil {
			accountUpdate.result <- updatedAccount
		}
		span.Finish()
	}
}

func (w *watcher) blockWatcher(ctx context.Context, startRound uint64) {
	defer w.logger.Println("Exited blockWatcher")
	round := startRound
	w.logger.Printf("Starting at round:%d", round)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			break
		}
		block, _ := w.algoClient.Block(round).Do(ctx)
		foundAccounts := map[string]struct{}{}
		// Iterate every transaction in the block, marking every account
		// updated in any way in that block.
		for _, txn := range block.Payset {
			if !txn.Txn.Receiver.IsZero() {
				foundAccounts[txn.Txn.Receiver.String()] = struct{}{}
			}
			if !txn.Txn.CloseRemainderTo.IsZero() {
				foundAccounts[txn.Txn.CloseRemainderTo.String()] = struct{}{}
			}
			if !txn.Txn.Sender.IsZero() {
				foundAccounts[txn.Txn.Sender.String()] = struct{}{}
			}
			if !txn.Txn.AssetReceiver.IsZero() {
				foundAccounts[txn.Txn.AssetReceiver.String()] = struct{}{}
			}
			if !txn.Txn.AssetSender.IsZero() {
				foundAccounts[txn.Txn.AssetSender.String()] = struct{}{}
			}
			if !txn.Txn.AssetCloseTo.IsZero() {
				foundAccounts[txn.Txn.AssetCloseTo.String()] = struct{}{}
			}
			if !txn.Txn.FreezeAccount.IsZero() {
				foundAccounts[txn.Txn.FreezeAccount.String()] = struct{}{}
			}
			// Any accounts referenced as part of contract operations
			if txn.Txn.ApplicationID != 0 {
				//w.logger.Printf("Block: %d, Appid:%d, Sender:%s", round, txn.Txn.ApplicationID, txn.Txn.Sender.String())
				for i, args := range txn.Txn.ApplicationArgs {
					//w.logger.Printf("ContractArgs: [%d] %s", i, string(args))
					if i == 2 {
						address := types.Address{}
						copy(address[:], args)
						//w.logger.Printf("ContractArgs: [%d] %s", i, address.String())
					}
				}
				for _, account := range txn.Txn.Accounts {
					//w.logger.Printf("ContractAccount: [%d] %s", i, account.String())
					foundAccounts[account.String()] = struct{}{}
				}
			}
		}
		// Now get our unique list of accounts - put into list form... and see which match our currently 'watched' accounts
		updatedAccounts := make([]string, 0, len(foundAccounts))
		for account := range foundAccounts {
			updatedAccounts = append(updatedAccounts, account)
		}
		// Now we check our unique map of 'touched' accounts against our map of 'watched' accounts
		// and add to our queue of accounts to update balances for in the background
		matches, err := w.persist.GetWatchedAccountMatches(ctx, updatedAccounts)
		if err != nil {
			w.logger.Printf("blockWatcher error in GetWatchedAccountMatches: %s", err.Error())
			// start over with same round number - maybe its a temporary issue w/ persistance layer
			// TODO: need more formal error logging...
			time.Sleep(100 * time.Millisecond)
			continue
		} else {
			if len(matches) > 0 {
				span := sentry.StartSpan(
					ctx, "dirty-account-dispatch",
					sentry.TransactionName("update dirty accounts"),
				)

				results := make(chan *Account, len(matches))
				start := time.Now()
				w.logger.Printf("block:%d - queuing account updates of %d accounts", round, len(matches))
				for _, address := range matches {
					w.accountUpdateChan <- addressAndResult{address: address, result: results, ctx: span.Context()}
				}
				// wait until we've received every result...
				for i := 0; i < len(matches); i++ {
					select {
					case <-ctx.Done():
						return
					case <-results:
						break
					case <-time.After(20 * time.Second):
						w.logger.Panicf(
							"Something went wrong in fetching block data.  More than 20 seconds have elapsed waiting for result %d of %d, exiting for now",
							i, len(matches),
						)
					}
				}
				w.logger.Printf("block:%d - account updates complete - elapsed:%v", round, time.Now().Sub(start))
				span.Finish()
			}
		}
		w.persist.SetLastRound(ctx, round)
		if round%100 == 0 {
			// just to show we're alive... lets show some logging activity periodically
			w.logger.Printf("Updated round:%d", round)
		}

		_, _ = w.algoClient.StatusAfterBlock(round).Do(ctx)
		round++
	}
}

// updateAccountInfo fetches the latest balance information from the node for the specified account and all of its
// assets - updating the persistence layer with updated asset info and the resulting
func (w *watcher) updateAccountInfo(ctx context.Context, address string) (*Account, error) {
	accountInfo, err := w.algoClient.AccountInformation(address).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching account accountInfo for:%s : %w", address, err)
	}
	retAccount := &Account{
		Address:  address,
		Round:    accountInfo.Round,
		Holdings: make(map[uint64]*Holding, len(accountInfo.Assets)+1),
	}

	retAccount.Holdings[algoID] = &Holding{
		AssetID: algoID,
		Amount:  accountInfo.Amount,
		Info: AssetInformation{
			Deleted:  false,
			Decimals: 6,
			Name:     "ALGO",
			UnitName: "ALGO",
		},
	}
	// Walk the assets from the nodes' current account info - asset id / amount held
	for _, asset := range accountInfo.Assets {
		var (
			assetInfo *AssetInformation
			err       error
		)
		assetInfo, err = w.persist.GetAssetInfo(ctx, asset.AssetId)
		if err != nil {
			w.logger.Printf("error in getAssetInfo:%s", err.Error())
			continue
		}
		if assetInfo == nil {
			// not cached yet - fetch it
			assetInfo, err = w.getCurrentAssetInfo(ctx, asset.AssetId)
			if err != nil {
				w.logger.Printf("error fetching un-cached asset, error:%s", err.Error())
				continue
			}
			if err = w.persist.SetAssetInfo(ctx, asset.AssetId, assetInfo); err != nil {
				// couldn't set into persistence layer but can still set into our return value so don't skip this one...
				w.logger.Printf("error setting fetching asset info, error:%s", err.Error())
			}
		}
		retAccount.Holdings[asset.AssetId] = &Holding{
			AssetID: asset.AssetId,
			Amount:  asset.Amount,
			Info:    *assetInfo,
		}
	}
	if err = w.persist.UpdateAccount(ctx, retAccount); err != nil {
		w.logger.Printf("updateAccountInfo - updating account:%s, error:%s", address, err.Error())
	}

	return retAccount, nil
}

func (w *watcher) getCurrentAssetInfo(ctx context.Context, assetID uint64) (*AssetInformation, error) {
	assetData, err := w.algoClient.GetAssetByID(assetID).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("getCurrentAssetInfo error in fetch of asset id:%d, error:%w", assetID, err)
	}
	return &AssetInformation{
		Deleted:      assetData.Deleted,
		Decimals:     assetData.Params.Decimals,
		MetadataHash: assetData.Params.MetadataHash,
		Name:         assetData.Params.Name,
		UnitName:     assetData.Params.UnitName,
		Url:          assetData.Params.Url,
	}, nil
}
