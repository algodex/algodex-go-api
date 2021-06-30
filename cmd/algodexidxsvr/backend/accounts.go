package backend

import (
	"context"
	"fmt"
	"sync"
)

var watchedData watchData

type accountMap map[string]*trackedAccount

type trackedAccount struct {
	sync.RWMutex
	// Public Account address
	Address string
	// Opted-in ASA IDs
	Assets []uint64
}

func (ta *trackedAccount) UpdateHoldings(ctx context.Context) error {
	holdings, err := getAccountHoldings(ctx, ta.Address)
	if err != nil {
		return err
	}
	ta.Lock()
	defer ta.Unlock()
	ta.Assets = make([]uint64, 0, len(holdings))
	for id := range holdings {
		ta.Assets = append(ta.Assets, id)
	}
	return nil
}

type watchData struct {
	sync.RWMutex
	watchedAccounts accountMap
}

func WatchAccount(ctx context.Context, address string) error {
	account := &trackedAccount{
		Address: address,
	}
	err := account.UpdateHoldings(ctx)
	if err != nil {
		return fmt.Errorf("couldn't add watch on account: %w", err)
	}
	watchedData.Lock()
	defer watchedData.Unlock()
	watchedData.watchedAccounts[address] = account
	return nil
}

func GetAccount(address string) *trackedAccount {
	return watchedData.GetAccount(address)
}

func GetAccounts() []*trackedAccount {
	return watchedData.GetWatchedAccounts()
}

func (w *watchData) GetAccount(address string) *trackedAccount {
	w.RLock()
	defer w.RUnlock()
	return watchedData.watchedAccounts[address]
}

func (w *watchData) GetWatchedAccounts() []*trackedAccount {
	w.RLock()
	defer w.RUnlock()
	accounts := make([]*trackedAccount, 0, len(w.watchedAccounts))
	for _, account := range w.watchedAccounts {
		accounts = append(accounts, account)
	}
	return accounts
}

func (w *watchData) IsWatchedAccount(toMatch map[string]bool, matched func(*trackedAccount)) {
	w.RLock()
	defer w.RUnlock()
	if w.watchedAccounts == nil {
		w.watchedAccounts = accountMap{}
	}
	for account, _ := range toMatch {
		if yard, found := w.watchedAccounts[account]; found {
			matched(yard)
		}
	}
}
