package backend

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
)

func accountWatcher(ctx context.Context, logger *log.Logger) {
	defer logger.Println("Exited yardWalletUpdater")
	wg := sync.WaitGroup{}

	accountUpdateChan := make(chan *trackedAccount, 1000)
	// Create parallel update workers - 4x core count
	for i := 0; i < runtime.NumCPU()*4; i++ {
		wg.Add(1)
		go accountUpdater(ctx, &wg, logger, accountUpdateChan)
	}

	// at startup - update ALL watched accounts
	for _, account := range watchedData.GetWatchedAccounts() {
		accountUpdateChan <- account
	}
	// Then we just watch for updates in each new block
	go blockWatcher(ctx, logger, accountUpdateChan)

	<-ctx.Done()
	close(accountUpdateChan)
	wg.Wait()
}

func accountUpdater(ctx context.Context, wg *sync.WaitGroup, logger *log.Logger, updateChan chan *trackedAccount) {
	defer wg.Done()
	// Endlessly loop on accounts to update the assets for until our channel is closed
	for account := range updateChan {
		logger.Printf("Account holdings update, account:%s", account.Address)
		err := account.UpdateHoldings(ctx)
		if err != nil {
			logger.Printf("error fetching holdings for account:%v, error:%v", account.Address, err.Error())
			continue
		}
	}
}

func blockWatcher(ctx context.Context, logger *log.Logger, updateChan chan *trackedAccount) {
	defer logger.Println("Exited blockYardWatcher")
	nodeStatus, err := algoClient.Status().Do(ctx)
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
		block, _ := algoClient.Block(round).Do(ctx)
		foundAccounts := map[string]bool{}
		for _, txn := range block.Payset {
			if !txn.Txn.Receiver.IsZero() {
				foundAccounts[txn.Txn.Receiver.String()] = true
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
		watchedData.IsWatchedAccount(
			foundAccounts, func(account *trackedAccount) {
				logger.Printf("Block with transactions, block:%d, account:%s", round, account.Address)
				updateChan <- account
			},
		)
		_, _ = algoClient.StatusAfterBlock(round).Do(ctx)
		round++
	}
}
