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

func InitBackend(ctx context.Context, log *log.Logger, network string) {
	initAlgoClient(os.Getenv("ALGORAND_DATA"), log, network) // will currently crash hard if can't initialize
	go accountWatcher(ctx, log)
}

func initAlgoClient(dataDir string, log *log.Logger, network string) {
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
		log.Fatal("failed to parse url:", apiURL, "error:", err)
	}

	algoClient, err = algod.MakeClient(serverAddr.String(), apiKey)
	if err != nil {
		log.Fatal("failed to make algod client (url:", serverAddr.String(), ")", err)
	}
}
