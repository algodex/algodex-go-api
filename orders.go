package algodexidx

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"algodexidx/backend"
	orders "algodexidx/gen/orders"
	"github.com/jmoiron/sqlx"
)

// orders service example implementation.
// The example methods log the requests and return zero values.
type orderssrvc struct {
	logger  *log.Logger
	backend backend.Itf
}

// NewOrders returns the orders service implementation.
func NewOrders(logger *log.Logger, itf backend.Itf) orders.Service {
	return &orderssrvc{logger, itf}
}

// Get all open orders for a specific asset or owners
func (s *orderssrvc) Get(ctx context.Context, p *orders.GetPayload) (res *orders.Orders, err error) {
	res = &orders.Orders{}
	var assetID uint64
	if p.AssetID != nil {
		assetID = *p.AssetID
	}
	s.logger.Printf("orders.forAsset, asset:%d, ownerAddrs:%#v", assetID, p.OwnerAddr)
	if p.AssetID == nil && len(p.OwnerAddr) == 0 {
		return nil, orders.MakeMissingParameters(fmt.Errorf("assetId or ownerAddr are empty"))
	}
	buyApplicationID, err := strconv.ParseUint(os.Getenv("ALGODEX_ALGO_ESCROW_APP"), 10, 64)
	sellApplicationID, err := strconv.ParseUint(os.Getenv("ALGODEX_ASA_ESCROW_APP"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("missing escrow app id")
	}

	// Handle buy orders
	res.BuyASAOrdersInEscrow, err = execOrderbookQuery(ctx, s.backend, buyApplicationID, p)
	if err != nil {
		return nil, fmt.Errorf("error in query execution: %w", err)
	}
	res.SellASAOrdersInEscrow, err = execOrderbookQuery(ctx, s.backend, sellApplicationID, p)
	if err != nil {
		return nil, fmt.Errorf("error in query execution: %w", err)
	}
	// Walk all the assets in each record so we can reformat some of the prices
	var assetDecimals = map[uint64]uint64{}

	updateAssetData := func(orders []*orders.Order) error {
		for _, order := range orders {
			decimals, found := assetDecimals[order.AssetID]
			if !found {
				assetInfo, err := s.backend.GetAssetInfo(ctx, order.AssetID)
				if err != nil {
					return err
				}
				assetDecimals[order.AssetID] = assetInfo.Decimals
				decimals = assetInfo.Decimals
			}
			asaAmount := float64(order.AsaAmount)
			floatAsaPrice, _ := strconv.ParseFloat(order.AsaPrice, 64)
			if asaAmount == 0 {
				asaAmount = float64(order.AlgoAmount) / floatAsaPrice
			}
			order.FormattedPrice = getFormattedPrice(floatAsaPrice, decimals)
			order.FormattedASAAmount = getFormattedASA_Amount(asaAmount, decimals)
			order.Decimals = decimals
		}
		return nil
	}
	if err = updateAssetData(res.BuyASAOrdersInEscrow); err != nil {
		return nil, fmt.Errorf("error fetching asset info: %w", err)
	}
	if err = updateAssetData(res.SellASAOrdersInEscrow); err != nil {
		return nil, fmt.Errorf("error fetching asset info: %w", err)
	}

	return

}

func getFormattedPrice(price float64, decimals uint64) string {
	return fmt.Sprintf("%.06f", price*math.Pow10(int(decimals)-6))
}

func getFormattedASA_Amount(amount float64, decimals uint64) string {
	return fmt.Sprintf(fmt.Sprintf("%%.0%df", decimals), amount/math.Pow10(int(decimals)))
}

func execOrderbookQuery(ctx context.Context, itf backend.Itf, applicatonID uint64, p *orders.GetPayload) ([]*orders.Order, error) {
	var (
		db     *sqlx.DB
		stmt   *sqlx.Stmt
		inArgs []interface{}
		result []*orders.Order
		err    error
	)
	db, err = itf.GetRawSQLHandle(ctx)
	if err != nil {
		return nil, err
	}
	if p.AssetID != nil {
		stmt, err = db.Preparex(
			`SELECT cast((denominator/numerator) as decimal(30,12)) AS 'assetLimitPriceInAlgos',
				 cast((denominator/numerator) as decimal(30,12)) AS 'asaPrice',
				 denominator AS 'assetLimitPriceD', numerator AS 'assetLimitPriceN',
				 algoAmount, asaAmount, assetid as 'assetId', appid as 'appId',
				 address as 'escrowAddress', ownerAddress,
				 minimum AS 'minimumExecutionSizeInAlgo',
				 round, unix_time
				 FROM orderbook WHERE appid = ? AND assetid = ?
				 ORDER BY round DESC;`,
		)
	} else {
		query, args, newErr := sqlx.In(
			`SELECT cast((denominator/numerator) as decimal(30,12)) AS 'assetLimitPriceInAlgos',
				 cast((denominator/numerator) as decimal(30,12)) AS 'asaPrice',
				 denominator AS 'assetLimitPriceD', numerator AS 'assetLimitPriceN',
				 algoAmount, asaAmount, assetid as 'assetId', appid as 'appId',
				 address as 'escrowAddress', ownerAddress,
				 minimum AS 'minimumExecutionSizeInAlgo',
				 round, unix_time
				 FROM orderbook WHERE appid = ? AND ownerAddress in ( ? )
				 ORDER BY round DESC;`, applicatonID, p.OwnerAddr,
		)
		if newErr != nil {
			return nil, fmt.Errorf("error in in preparation syntax: %w", newErr)
		}
		query = db.Rebind(query)
		inArgs = args
		stmt, err = db.Preparex(query)
	}
	if err != nil {
		return nil, fmt.Errorf("error in prepare call: %w", err)
	}
	if p.AssetID != nil {
		err = stmt.SelectContext(ctx, &result, applicatonID, *p.AssetID)
	} else {
		err = stmt.SelectContext(ctx, &result, inArgs...)
	}
	return result, err
}
