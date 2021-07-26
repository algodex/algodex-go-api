package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Persistor interface {
	Close() error

	GetLastRound(context.Context) (uint64, error)
	SetLastRound(ctx context.Context, round uint64) error
	GetWatchedAccounts(context.Context) ([]string, error)
	WatchAccounts(ctx context.Context, addresses ...string) error
	GetAssetInfo(ctx context.Context, assetID uint64) (AssetInformation, error)
	SetAssetInfo(ctx context.Context, asset AssetInformation) error
	//AddWatches(context.Context, )
	//GetWatchedAccounts() []*Account
	//GetAccount(address string) (Account, error)
}

/*
type Account struct {
// Public Account address
Address string
// Opted-in ASA information
Holdings map[uint64]*Holding
}

}
*/

func initPersistance(ctx context.Context) *persistor {
	ret := &persistor{
		redis: redis.NewClient(
			&redis.Options{
				Addr: os.Getenv("ALGODEX_REDIS_ADDR"),
			},
		),
	}
	sqlConn, err := sqlx.Connect(
		"mysql", fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("ALGODEX_DB_USER"), os.Getenv("ALGODEX_DB_PASS"),
			os.Getenv("ALGODEX_DB_HOST"), os.Getenv("ALGODEX_DB_PORT"),
			os.Getenv("ALGODEX_DB_NAME"),
		),
	)
	if err != nil {
		log.Panicf("error in opening sql connection: %v", err)
	}
	// Some examples we don't care about... yet
	//sqlConn.SetConnMaxLifetime(time.Minute * 3)
	//sqlConn.SetMaxOpenConns(10)
	//sqlConn.SetMaxIdleConns(10)
	ret.sql = sqlConn

	//var round uint64
	//sqlConn.Get(&round, "SELECT round FROM config")
	//log.Println("round from mysql is:", round)

	return ret
}

type persistor struct {
	redis *redis.Client
	sql   *sqlx.DB
}

func redisKey(typeName, key string) string {
	return "dexidx:" + typeName + ":" + key
}

func (p *persistor) Close() error {
	return p.redis.Close()
	// don't call p.sql.Close() - it's rare to close out the sql drivers
}

func (p *persistor) GetLastRound(ctx context.Context) (uint64, error) {
	val, err := p.redis.Get(ctx, redisKey("algod", "round")).Uint64()
	if err != nil {
		return 0, fmt.Errorf("calling GetLastRound: %w", err)
	}
	return val, nil
}

func (p *persistor) SetLastRound(ctx context.Context, round uint64) error {
	err := p.redis.Set(ctx, redisKey("algod", "round"), strconv.FormatUint(round, 10), 0).Err()
	if err != nil {
		return fmt.Errorf("calling SetLastRound: %w", err)
	}
	return nil
}

func (p *persistor) GetWatchedAccounts(ctx context.Context) ([]string, error) {
	val, err := p.redis.SMembers(ctx, redisKey("account", "watched")).Result()
	if err != nil {
		return nil, fmt.Errorf("calling GetWatchedAccounts: %w", err)
	}
	return val, nil
}

func (p *persistor) WatchAccounts(ctx context.Context, addresses ...string) error {
	redisStrings := make([]interface{}, len(addresses))
	for i := range addresses {
		redisStrings[i] = addresses[i]
	}
	err := p.redis.SAdd(ctx, redisKey("accounts", "watched"), redisStrings...).Err()
	if err != nil {
		return fmt.Errorf("calling WatchAccounts: %w", err)
	}
	return nil
}

func (p *persistor) GetAssetInfo(ctx context.Context, assetID uint64) (AssetInformation, error) {
	panic("implement me")
}

func (p *persistor) SetAssetInfo(ctx context.Context, asset AssetInformation) error {
	panic("implement me")
}
