package price

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	"github.com/aureliomalheiros/coinprice/internal/validates"
)

type CoinPrice struct {
	BRL float64 `json:"brl"`
	USD float64 `json:"usd"`
}

type CoinGeckoResponse map[string]CoinPrice

func GetCryptoPrice(coin, currency string) (float64, error) {
	if err := validates.ValidateInput(coin, currency); err != nil {
		return 0, err
	}

	if coin == "" {
		coin = "bitcoin"
	}

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", coin, currency)

	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error fetching data from CoinGecko: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error: received status code %d from CoinGecko", resp.StatusCode)
	}

	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("error decoding JSON: %w", err)
	}

	coinData, exists := result[coin]
	if !exists {
		return 0, fmt.Errorf("coin '%s' not found in response", coin)
	}

	switch currency {
	case "brl":
		return coinData.BRL, nil
	case "usd":
		return coinData.USD, nil
	default:
		return 0, fmt.Errorf("unsupported currency: %s", currency)
	}
}
