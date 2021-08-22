package backend

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/jmoiron/sqlx"
)

//var algoClient *algod.Client

type Itf interface {
	WatchAccounts(ctx context.Context, addresses ...string) error
	UnwatchAccounts(ctx context.Context, addresses ...string) error
	GetAccount(ctx context.Context, address string) (*Account, error)
	GetAccounts(ctx context.Context) []*Account
	IsWatchedAccount(ctx context.Context, accounts []string) ([]string, error)
	Reset(ctx context.Context) error
	GetAssetInfo(ctx context.Context, assetID uint64) (*AssetInformation, error)
	GetRawSQLHandle(ctx context.Context) (*sqlx.DB, error)
}

type backendState struct {
	log        *log.Logger
	algoClient *algod.Client
	persist    Persistor
	watcher    *watcher
}

func (b *backendState) WatchAccounts(ctx context.Context, addresses ...string) error {
	return b.watcher.WatchAccounts(ctx, addresses...)
}

func (b *backendState) UnwatchAccounts(ctx context.Context, addresses ...string) error {
	return b.watcher.UnwatchAccounts(ctx, addresses...)
}

func (b *backendState) GetAccount(ctx context.Context, address string) (*Account, error) {
	return b.watcher.GetAccount(ctx, address)
}

func (b *backendState) GetAccounts(ctx context.Context) []*Account {
	return b.watcher.GetAccounts(ctx)
}

func (b *backendState) IsWatchedAccount(ctx context.Context, accounts []string) ([]string, error) {
	return b.persist.GetWatchedAccountMatches(ctx, accounts)
}

func (b *backendState) Reset(ctx context.Context) error {
	return b.persist.Reset(ctx)
}

func (b *backendState) GetAssetInfo(ctx context.Context, assetID uint64) (*AssetInformation, error) {
	return b.watcher.GetAssetInfo(ctx, assetID)
}

func (b *backendState) GetRawSQLHandle(ctx context.Context) (*sqlx.DB, error) {
	return b.persist.GetRawSQLHandle(ctx)
}

func InitBackend(ctx context.Context, log *log.Logger, network string) (*backendState, error) {
	var err error
	be := &backendState{log: log}

	if err = validateEnvironment(network); err != nil {
		return nil, err
	}

	be.algoClient, err = initAlgoClient(os.Getenv("ALGORAND_DATA"), log, network)
	if err != nil {
		return nil, fmt.Errorf("failure in algo client setup: %w", err)
	}
	be.persist = initPersistance(ctx, log)

	// Start the block watcher - giving it persistence interface for getting data/pushing updates...
	be.watcher = newWatcher(log, be.algoClient, be.persist)
	be.watcher.start(ctx)
	return be, nil
}

func validateEnvironment(network string) error {
	// TODO: Verify all environment variables we need are present
	if network == "" {
		network = os.Getenv("ALGODEX_NETWORK")
	}
	if network != "testnet" && network != "mainnet" {
		return fmt.Errorf("invalid algorand network:%s", network)
	}

	return nil

}

func initAlgoClient(dataDir string, log *log.Logger, network string) (*algod.Client, error) {
	var (
		apiURL     string
		apiToken   string
		serverAddr *url.URL
		err        error
	)
	if dataDir != "" {
		// Read address and token from main-net directory
		netPath, err := ioutil.ReadFile(filepath.Join(dataDir, "algod.net"))
		if err != nil {
			log.Fatal("error reading data-dir file", err)
		}
		apiKeyBytes, err := ioutil.ReadFile(filepath.Join(dataDir, "algod.token"))
		if err != nil {
			log.Fatal("error reading data-dir file", err)
		}
		apiURL = fmt.Sprintf("http://%s", strings.TrimSpace(string(netPath)))
		apiToken = string(apiKeyBytes)
	} else {
		apiURL = os.Getenv("ALGODEX_ALGOD")
		apiToken = os.Getenv("ALGODEX_ALGOD_TOKEN")
		if apiURL == "" {
			if network == "testnet" {
				apiURL = "https://api.testnet.algoexplorer.io"
			} else if network == "mainnet" {
				apiURL = "https://api.algoexplorer.io"
			}
		}
		// Strip off trailing slash if present in url which the Algorand client doesn't handle properly
		apiURL = strings.TrimRight(apiURL, "/")
	}
	serverAddr, err = url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url:%v, error:%w", apiURL, err)
	}
	log.Printf("Connecting to node at:%s", serverAddr.String())

	client, err := algod.MakeClient(serverAddr.String(), apiToken)
	if err != nil {
		return nil, fmt.Errorf(`failed to make algod client (url:%s), error:%w`, serverAddr.String(), err)
	}
	return client, err
}
