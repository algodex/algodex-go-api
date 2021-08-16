package algodexidx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"algodexidx/backend"
	"algodexidx/gen/account"
	"github.com/algorand/go-algorand-sdk/types"
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
	if err := backend.FailIfNotAuthorized(ctx); err != nil {
		return err
	}
	if p == nil || len(p.Address) == 0 {
		return errors.New("must provide address(es) to watch")
	}
	for _, address := range p.Address {
		_, err = types.DecodeAddress(address)
		if err != nil {
			return fmt.Errorf("address:%v not valid: %w", address, err)
		}
	}
	// pass on to persistence backend... (redis for eg)
	err = s.backend.WatchAccounts(ctx, p.Address...)
	if err != nil {
		return fmt.Errorf("account watch persistence add of addresses:%v, error:%w", p.Address, err)
	}
	return
}

// Delete Algorand account(s) to track
func (s *accountsrvc) Delete(ctx context.Context, p *account.DeletePayload) (err error) {
	s.logger.Println("account.delete", p.Address)
	if err := backend.FailIfNotAuthorized(ctx); err != nil {
		return err
	}
	if p == nil || len(p.Address) == 0 {
		return errors.New("must provide address(es) to remove")
	}
	// pass on to persistence backend... (redis for eg)
	err = s.backend.UnwatchAccounts(ctx, p.Address...)
	if err != nil {
		return fmt.Errorf("account watch persistence delete of addresses:%v, error:%w", p.Address, err)
	}
	return
}

// Delete all tracked algorand account(s).  Used for resetting everything
func (s *accountsrvc) Deleteall(ctx context.Context) (err error) {
	s.logger.Print("account.deleteall")
	if err := backend.FailIfNotAuthorized(ctx); err != nil {
		return err
	}
	return s.backend.Reset(ctx)
}

// Get specific account
func (s *accountsrvc) Get(ctx context.Context, p *account.GetPayload) (res *account.Account, err error) {
	s.logger.Println("account.get", p.Address)
	if err := backend.FailIfNotAuthorized(ctx); err != nil {
		return nil, err
	}
	backendAccount, err := s.backend.GetAccount(ctx, p.Address)
	if backendAccount == nil || err != nil {
		return nil, fmt.Errorf("account:%s couldn't be retrieved or other error: %w", p.Address, err)
	}
	res = &account.Account{
		Address:  backendAccount.Address,
		Round:    backendAccount.Round,
		Holdings: backendHoldingToDSLHolding(backendAccount.Holdings),
	}
	return
}

// Get account(s)
func (s *accountsrvc) GetMultiple(ctx context.Context, p *account.GetMultiplePayload) (res []*account.Account, err error) {
	s.logger.Println("account.getMultiple", p.Address)
	if err := backend.FailIfNotAuthorized(ctx); err != nil {
		return nil, err
	}
	for _, address := range p.Address {
		backendAccount, err := s.backend.GetAccount(ctx, address)
		if backendAccount == nil || err != nil {
			return nil, fmt.Errorf("account:%s couldn't be retrieved or other error: %w", address, err)
		}
		res = append(res, &account.Account{
			Address:  backendAccount.Address,
			Round:    backendAccount.Round,
			Holdings: backendHoldingToDSLHolding(backendAccount.Holdings),
		})
	}
	return
}

// List all tracked accounts
func (s *accountsrvc) List(ctx context.Context, p *account.ListPayload) (
	res account.TrackedAccountCollection, view string, err error,
) {
	if err = backend.FailIfNotAuthorized(ctx); err != nil {
		return
	}
	view = "default"
	if p.View != nil {
		view = *p.View
	}
	s.logger.Println("account.list, view:", view)
	for _, acct := range s.backend.GetAccounts(ctx) {
		res = append(
			res, &account.TrackedAccount{
				Address:  acct.Address,
				Round:    acct.Round,
				Holdings: backendHoldingToDSLHolding(acct.Holdings),
			},
		)
	}
	return
}

// Returns which of the passed accounts are currently being monitored
func (s *accountsrvc) Iswatched(ctx context.Context, p *account.IswatchedPayload) (res []string, err error) {
	s.logger.Print("account.iswatched", p.Address)
	if err := backend.FailIfNotAuthorized(ctx); err != nil {
		return nil, err
	}
	return s.backend.IsWatchedAccount(ctx, p.Address)
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
