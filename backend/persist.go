package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Persistor interface {
	Close() error

	GetLastRound(context.Context) (uint64, error)
	SetLastRound(ctx context.Context, round uint64) error
	GetWatchedAccounts(context.Context) ([]string, error)
	GetWatchedAccountMatches(context.Context, []string) ([]string, error)
	WatchAccounts(ctx context.Context, addresses ...string) error
	GetAccount(ctx context.Context, address string) (*Account, error)
	UpdateAccount(ctx context.Context, account *Account) error
	GetAssetInfo(ctx context.Context, assetID uint64) (*AssetInformation, error)
	SetAssetInfo(ctx context.Context, assetID uint64, asset *AssetInformation) error
	//AddWatches(context.Context, )
	//GetWatchedAccounts() []*Account
	//GetAccount(address string) (Account, error)
}

func initPersistance(ctx context.Context, log *log.Logger) *persistor {
	log.Printf("connecting to redis host:%s", os.Getenv("ALGODEX_REDIS_ADDR"))
	ret := &persistor{
		redis: redis.NewClient(
			&redis.Options{
				Addr: os.Getenv("ALGODEX_REDIS_ADDR"),
			},
		),
	}
	log.Printf(
		"conecting to mysql host:%s:%s db:%s", os.Getenv("ALGODEX_DB_HOST"), os.Getenv("ALGODEX_DB_PORT"),
		os.Getenv("ALGODEX_DB_NAME"),
	)
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
	//ret.sql = sqlConn

	var round uint64
	sqlConn.Get(&round, "SELECT round FROM config")
	log.Println("round from mysql is:", round)

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
	val, err := p.redis.SMembers(ctx, redisKey("accounts", "watched")).Result()
	if err != nil {
		return nil, fmt.Errorf("calling GetWatchedAccounts: %w", err)
	}
	return val, nil
}

func (p *persistor) GetWatchedAccountMatches(ctx context.Context, doesMatch []string) ([]string, error) {
	// AWS doesn't support SMIsMember yet, so just walk them all manually
	matches := make([]string, 0, len(doesMatch))
	for _, address := range doesMatch {
		member, err := p.redis.SIsMember(ctx, redisKey("accounts", "watched"), address).Result()
		if err != nil {
			return nil, fmt.Errorf("calling GetWatchedAccountMatches: %w", err)
		}
		if member {
			matches = append(matches, address)
		}
	}
	// Switch to this once AWS supports Redis 6.2
	//redisStrings := make([]interface{}, len(doesMatch))
	//for i := range doesMatch {
	//	redisStrings[i] = doesMatch[i]
	//}
	//// We get back indexed true/false for whether each element was member of set (our watched accounts)
	//members, err := p.redis.SMIsMember(ctx, redisKey("accounts", "watched"), redisStrings...).Result()
	//if err != nil {
	//	return nil, fmt.Errorf("calling GetWatchedAccountMatches: %w", err)
	//}
	//matches := make([]string, 0, len(redisStrings))
	//for i, isMember := range members {
	//	if isMember {
	//		matches = append(matches, doesMatch[i])
	//	}
	//}
	return matches, nil
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

func (p *persistor) GetAccount(ctx context.Context, address string) (*Account, error) {
	val, err := p.redis.Get(ctx, redisKey("account", address)).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("calling GetWatchedAccounts: %w", err)
	}
	return AccountFromJSON(val)
}

func (p *persistor) UpdateAccount(ctx context.Context, account *Account) error {
	b, err := account.ToJSON()
	if err != nil {
		return fmt.Errorf("error in UpdateAccount json marshal: %w", err)
	}
	err = p.redis.Set(ctx, redisKey("account", account.Address), b, 7*24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("error in UpdateAccount: %w", err)
	}
	return nil
}

func (p *persistor) GetAssetInfo(ctx context.Context, assetID uint64) (*AssetInformation, error) {
	val, err := p.redis.Get(ctx, redisKey("asset:info", strconv.FormatUint(assetID, 10))).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("calling GetAssetInfo: %w", err)
	}
	return AssetInformationFromJSON(val)
}

func (p *persistor) SetAssetInfo(ctx context.Context, assetID uint64, asset *AssetInformation) error {
	b, err := asset.ToJSON()
	if err != nil {
		return fmt.Errorf("error in SetAssetInfo json marshal: %w", err)
	}
	err = p.redis.Set(ctx, redisKey("asset:info", strconv.FormatUint(assetID, 10)), b, 1*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("error in SetAssetInfo: %w", err)
	}
	return nil
}
