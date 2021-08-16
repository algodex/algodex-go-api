package backend

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Persistor interface {
	Close() error

	GetLastRound(context.Context) (uint64, error)
	SetLastRound(ctx context.Context, round uint64) error
	GetWatchedAccounts(context.Context) ([]string, error)
	GetWatchedAccountMatches(context.Context, []string) ([]string, error)
	WatchAccounts(ctx context.Context, addresses ...string) error
	UnwatchAccounts(ctx context.Context, addresses ...string) error
	GetAccount(ctx context.Context, address string) (*Account, error)
	UpdateAccount(ctx context.Context, account *Account) error
	GetAssetInfo(ctx context.Context, assetID uint64) (*AssetInformation, error)
	SetAssetInfo(ctx context.Context, assetID uint64, asset *AssetInformation) error
	Reset(ctx context.Context) error
}

func initPersistance(_ context.Context, log *log.Logger) *persistor {
	log.Printf("connecting to redis host:%s", os.Getenv("ALGODEX_REDIS_ADDR"))
	ret := &persistor{
		redis: redis.NewClient(&redis.Options{Addr: os.Getenv("ALGODEX_REDIS_ADDR")}),
		sql:   initSQL(),
	}

	return ret
}

func initSQL() *sqlx.DB {
	log.Printf(
		"conecting to mysql host:%s:%s db:%s", os.Getenv("ALGODEX_DB_HOST"), os.Getenv("ALGODEX_DB_PORT"),
		os.Getenv("ALGODEX_DB_NAME"),
	)
	var (
		// Retry for up to 2? minutes
		sqlConn *sqlx.DB
		err     error
	)
	retryExpiration := time.Now().Add(2 * time.Minute)
	for {
		sqlConn, err = sqlx.Connect(
			"mysql", fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s",
				os.Getenv("ALGODEX_DB_USER"), os.Getenv("ALGODEX_DB_PASS"),
				os.Getenv("ALGODEX_DB_HOST"), os.Getenv("ALGODEX_DB_PORT"),
				os.Getenv("ALGODEX_DB_NAME"),
			),
		)
		if err == nil {
			break
		}
		// We got an error... retry until we hit our expiration time
		if time.Now().After(retryExpiration) {
			log.Panicf("error in opening sql connection: %s", err.Error())
		}
		if errors.Is(err, syscall.ECONNREFUSED) {
			// if the host just isn't available we'll keep retrying
			log.Printf("Host conection refused... retrying")
		} else {
			var sqlErr *mysql.MySQLError
			if errors.As(err, &sqlErr) {
				if sqlErr.Number == 0x419 {
					// unknown database.. keep trying..
					log.Printf("Unknown database - orderbook hasn't created yet... retrying")
				} else {
					log.Panicf("Unexpected error in opening sql connection: %v", sqlErr)
				}
			} else {
				log.Panicf("Unexpected error in opening sql connection: %#v", err)
			}
		}
		time.Sleep(2 * time.Second)
	}
	sqlConn.SetConnMaxLifetime(time.Minute * 3)
	sqlConn.SetMaxOpenConns(10)
	sqlConn.SetMaxIdleConns(10)

	var round uint64
	_ = sqlConn.Get(&round, "SELECT round FROM config")
	log.Println("round from mysql is:", round)
	return sqlConn
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

func (p *persistor) UnwatchAccounts(ctx context.Context, addresses ...string) error {
	redisStrings := make([]interface{}, len(addresses))
	for i := range addresses {
		redisStrings[i] = addresses[i]
	}
	err := p.redis.SRem(ctx, redisKey("accounts", "watched"), redisStrings...).Err()
	if err != nil {
		return fmt.Errorf("calling UnwatchAccounts: %w", err)
	}
	// Go ahead and remove the account record if present.  We could wait for it to TTL out but this may be a transient
	// account and there may be a lot of them so there's no reason to hold onto that space.
	for _, address := range addresses {
		p.redis.Del(ctx, redisKey("account", address))
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

func (p *persistor) Reset(ctx context.Context) error {
	// Remove the only things that don't naturally TTL out which is our 'set' of watched accounts
	return p.redis.Del(ctx, redisKey("accounts", "watched")).Err()
}
