package backend

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

type watcher struct {
	logger            *log.Logger
	algoClient        *algod.Client
	accountUpdateChan chan *trackedAccount
}

func newWatcher(log *log.Logger, algoClient *algod.Client, persistor Persistor) *watcher {
	return &watcher{
		logger:     log,
		algoClient: algoClient,
	}
}

func (w *watcher) start(ctx context.Context) {
	go w.accountWatcher(ctx)
}

func (w *watcher) accountWatcher(ctx context.Context) {
	defer w.logger.Println("Exited accountWatcher")
	wg := sync.WaitGroup{}

	w.accountUpdateChan = make(chan *trackedAccount, 1000)
	// Create parallel update workers - 4x core count
	for i := 0; i < runtime.NumCPU()*4; i++ {
		wg.Add(1)
		go w.accountUpdater(ctx, &wg)
	}

	// at startup - update ALL watched accounts
	for _, account := range watchedData.GetWatchedAccounts() {
		w.accountUpdateChan <- account
	}
	// Then we just watch for updates in each new block
	go w.blockWatcher(ctx)

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
	for account := range w.accountUpdateChan {
		w.logger.Printf("Account holdings update, account:%s", account.Address)
		err := account.UpdateHoldings(ctx)
		if err != nil {
			w.logger.Printf("error fetching holdings for account:%v, error:%v", account.Address, err.Error())
			continue
		}
	}
}

func (w *watcher) blockWatcher(ctx context.Context) {
	defer w.logger.Println("Exited blockWatcher")
	nodeStatus, err := w.algoClient.Status().Do(ctx)
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}
	round := nodeStatus.LastRound

	for {
		select {
		case <-ctx.Done():
			return
		default:
			break
		}
		block, _ := w.algoClient.Block(round).Do(ctx)
		foundAccounts := map[string]bool{}
		// Iterate every transaction in the block, marking every account
		// updated in any way in that block.
		for _, txn := range block.Payset {
			if !txn.Txn.Receiver.IsZero() {
				foundAccounts[txn.Txn.Receiver.String()] = true
			}
			if !txn.Txn.CloseRemainderTo.IsZero() {
				foundAccounts[txn.Txn.CloseRemainderTo.String()] = true
			}
			if !txn.Txn.Sender.IsZero() {
				foundAccounts[txn.Txn.Sender.String()] = true
			}
			if !txn.Txn.AssetReceiver.IsZero() {
				foundAccounts[txn.Txn.AssetReceiver.String()] = true
			}
			if !txn.Txn.AssetSender.IsZero() {
				foundAccounts[txn.Txn.AssetSender.String()] = true
			}
			if !txn.Txn.AssetCloseTo.IsZero() {
				foundAccounts[txn.Txn.AssetCloseTo.String()] = true
			}
		}
		// Now we check our unique map of 'touched' accounts against our map of 'watched' accounts
		// and add to our queue of accounts to update balances for in the background
		watchedData.IsWatchedAccount(
			foundAccounts, func(account *trackedAccount) {
				w.logger.Printf("Block with transactions, block:%d, account:%s", round, account.Address)
				w.accountUpdateChan <- account
			},
		)
		_, _ = w.algoClient.StatusAfterBlock(round).Do(ctx)
		round++
	}
}
