package algodexidx

import (
	"context"
	"fmt"
	"log"

	"algodexidx/cmd/algodexidxsvr/backend"
	account "algodexidx/gen/account"
)

// account service example implementation.
// The example methods log the requests and return zero values.
type accountsrvc struct {
	logger *log.Logger
}

// NewAccount returns the account service implementation.
func NewAccount(logger *log.Logger) account.Service {
	return &accountsrvc{logger}
}

// Add Algorand account to track
func (s *accountsrvc) Add(ctx context.Context, p string) (err error) {
	err = backend.WatchAccount(ctx, p)
	if err != nil {
		return fmt.Errorf("account watch add of address:%s, error:%w", p, err)
	}
	s.logger.Print("account.add", p)
	return
}

// Get specific account
func (s *accountsrvc) Get(ctx context.Context, p *account.GetPayload) (res *account.Account, err error) {
	s.logger.Print("account.get", p.Address)
	backendAccount := backend.GetAccount(p.Address)
	if backendAccount == nil {
		return nil, fmt.Errorf("account:%s is not watched or other error", p.Address)
	}
	res = &account.Account{
		Address: backendAccount.Address,
		Assets:  make([]uint64, 0, len(backendAccount.Assets)),
	}
	for _, id := range backendAccount.Assets {
		res.Assets = append(res.Assets, id)
	}
	return
}

// List all tracked accounts
func (s *accountsrvc) List(ctx context.Context) (res account.TrackedAccountCollection, view string, err error) {
	view = "full"
	s.logger.Print("account.list")
	for _, acct := range backend.GetAccounts() {
		res = append(res, &account.TrackedAccount{
			Address: acct.Address,
			Assets:  acct.Assets,
		})
	}
	return
}
