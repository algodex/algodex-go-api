package algodexidx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"algodexidx/backend"
	"algodexidx/gen/account"
)

// account service example implementation.
// The example methods log the requests and return zero values.
type accountsrvc struct {
	logger  *log.Logger
	backend backend.Itf
}

// NewAccount returns the account service implementation.
func NewAccount(logger *log.Logger, itf backend.Itf) account.Service {
	return &accountsrvc{logger, itf}
}

// Add Algorand account to track
func (s *accountsrvc) Add(ctx context.Context, p *account.AddPayload) (err error) {
	s.logger.Println("account.add", p.Address)
	if p == nil || len(p.Address) == 0 {
		return errors.New("must provide address(es) to watch")
	}
	s.backend.WatchAccounts(ctx, p.Address...)
	err = backend.WatchAccounts(ctx, p.Address...)
	if err != nil {
		return fmt.Errorf("account watch add of addresses:%v, error:%w", p.Address, err)
	}
	return
}

// Get specific account
func (s *accountsrvc) Get(ctx context.Context, p *account.GetPayload) (res *account.Account, err error) {
	s.logger.Println("account.get", p.Address)
	backendAccount := backend.GetAccount(p.Address)
	if backendAccount == nil {
		return nil, fmt.Errorf("account:%s is not watched or other error", p.Address)
	}
	res = &account.Account{
		Address:  backendAccount.Address,
		Holdings: backendHoldingToDSLHolding(backendAccount.Holdings),
	}
	return
}

// List all tracked accounts
func (s *accountsrvc) List(ctx context.Context, p *account.ListPayload) (
	res account.TrackedAccountCollection, view string, err error,
) {
	view = "default"
	if p.View != nil {
		view = *p.View
	}
	s.logger.Println("account.list, view:", view)
	for _, acct := range backend.GetAccounts() {
		res = append(
			res, &account.TrackedAccount{
				Address:  acct.Address,
				Holdings: backendHoldingToDSLHolding(acct.Holdings),
			},
		)
	}
	return
}

func backendHoldingToDSLHolding(backendHolding map[uint64]*backend.Holding) map[string]*account.Holding {
	retHolding := make(map[string]*account.Holding, len(backendHolding))
	for key, holding := range backendHolding {
		retHolding[strconv.FormatUint(key, 10)] = &account.Holding{
			Asset:        holding.AssetID,
			Amount:       holding.Amount,
			Decimals:     holding.Info.Decimals,
			MetadataHash: string(holding.Info.MetadataHash),
			Name:         holding.Info.Name,
			UnitName:     holding.Info.UnitName,
			URL:          holding.Info.Url,
		}
	}
	return retHolding
}
