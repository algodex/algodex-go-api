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

var algoClient *algod.Client

type Itf interface {
	WatchAccounts(ctx context.Context, addresses ...string) error
	GetAccount(ctx context.Context, address string) Account
	GetAccounts(ctx context.Context) []Account
}

type backendState struct {
	log        *log.Logger
	algoClient *algod.Client
	persist    Persistor
	//watcher    Watcher
}

func (b *backendState) WatchAccounts(ctx context.Context, addresses ...string) error {
	return b.persist.WatchAccounts(ctx, addresses...)
}

func (b *backendState) GetAccount(ctx context.Context, address string) Account {
	// TODO implement!
	return Account{}
}

func (b *backendState) GetAccounts(ctx context.Context) []Account {
	// TODO implement!
	//accounts := b.persist.GetWatchedAccounts(ctx)
	//for _, account := range accounts {
	//
	//}
	return []Account{}
}

func InitBackend(ctx context.Context, log *log.Logger, network string) *backendState {
	var err error
	be := &backendState{log: log}

	algoClient, err = initAlgoClient(os.Getenv("ALGORAND_DATA"), log, network)
	if err != nil {
		log.Fatalf("failure in algo client setup: %v", err)
	}
	be.algoClient = algoClient
	be.persist = initPersistance(ctx)
	// Start the block watcher - giving it persistence interface for pushing updates...
	newWatcher(log, algoClient, be.persist).start(ctx)
	return be
}

func initAlgoClient(dataDir string, log *log.Logger, network string) (*algod.Client, error) {
	var (
		apiURL     string
		apiKey     string
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
		apiKey = string(apiKeyBytes)
	} else {
		if network == "testnet" {
			apiURL = "https://api.testnet.algoexplorer.io"
		} else if network == "mainnet" {
			apiURL = "https://api.algoexplorer.io"
		}
		if os.Getenv("API_ENDPOINT") != "" {
			apiURL = os.Getenv("API_ENDPOINT")
		}
		apiKey = os.Getenv("API_KEY")
	}
	serverAddr, err = url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url:%v, error:%w", apiURL, err)
	}

	client, err := algod.MakeClient(serverAddr.String(), apiKey)
	if err != nil {
		return nil, fmt.Errorf(`failed to make algod client (url:%s), error:%w`, serverAddr.String(), err)
	}
	return client, err
}
