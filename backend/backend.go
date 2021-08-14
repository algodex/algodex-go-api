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
)

//var algoClient *algod.Client

type Itf interface {
	WatchAccounts(ctx context.Context, addresses ...string) error
	UnwatchAccounts(ctx context.Context, addresses ...string) error
	GetAccount(ctx context.Context, address string) (*Account, error)
	GetAccounts(ctx context.Context) []*Account
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

func InitBackend(ctx context.Context, log *log.Logger, network string) *backendState {
	var err error
	be := &backendState{log: log}

	be.algoClient, err = initAlgoClient(os.Getenv("ALGORAND_DATA"), log, network)
	if err != nil {
		log.Fatalf("failure in algo client setup: %v", err)
	}
	be.persist = initPersistance(ctx, log)
	// Load all the accounts we've already been told to watch

	// Start the block watcher - giving it persistence interface for getting data/pushing updates...
	be.watcher = newWatcher(log, be.algoClient, be.persist)
	be.watcher.start(ctx)
	return be
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
