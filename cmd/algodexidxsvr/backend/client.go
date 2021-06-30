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
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

var algoClient *algod.Client

const (
	ALGO_ID     = 1
	apiEndpoint = "https://api.testnet.algoexplorer.io"
)

//const apiEndpoint = "https://api.algoexplorer.io"

type Holding struct {
	AssetID uint64
	Amount  uint64
}
type holdingsMap map[uint64]Holding

func InitAlgoClient(dataDir string, log *log.Logger) {
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
		apiURL = apiEndpoint
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

func getAccountHoldings(ctx context.Context, account string) (holdingsMap, error) {
	info, err := accountInformation(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("get account holdings:%s : %w", account, err)
	}
	holdings := make(holdingsMap, len(info.Assets)+1)
	holdings[ALGO_ID] = Holding{
		AssetID: ALGO_ID,
		Amount:  info.Amount,
	}
	for _, asset := range info.Assets {
		holdings[asset.AssetId] = Holding{
			AssetID: asset.AssetId,
			Amount:  asset.Amount,
		}
	}
	return holdings, nil
}

func accountInformation(ctx context.Context, account string) (models.Account, error) {
	accountInfo, err := algoClient.AccountInformation(account).Do(ctx)
	if err != nil {
		return models.Account{}, fmt.Errorf("fetching account info for:%s : %w", account, err)
	}
	return accountInfo, nil
}
