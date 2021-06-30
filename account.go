package algodexidx

import (
	"context"
	"log"

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
func (s *accountsrvc) Add(ctx context.Context, p string) (res *account.Account, err error) {
	res = &account.Account{}
	s.logger.Print("account.add", p)
	return
}

// Get specific account
func (s *accountsrvc) Get(ctx context.Context, p *account.GetPayload) (res *account.Account, err error) {
	res = &account.Account{}
	s.logger.Print("account.get", p.Address)
	return
}

// List all tracked accounts
func (s *accountsrvc) List(ctx context.Context) (res account.TrackedAccountCollection, view string, err error) {
	view = "default"
	s.logger.Print("account.list")
	return
}
