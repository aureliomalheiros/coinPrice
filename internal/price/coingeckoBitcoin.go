package price

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type coingeckoResponse struct {
	Bitcoin struct {
		BRL float64 `json:"brl"`
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func GetBitcoinPriceBRL(currency string) (float64, error) {

	validCurrency := map[string]bool {
		"brl": true,
		"usd": true,
	}

	if !validCurrency[currency] {
		return 0, fmt.Errorf("Invalid currency: %s. Valid options are: brl, usd", currency)
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=%s", currency)

	resp, err := http.Get(url)
	if err != nil  {
		return 0, fmt.Errorf("Error fetching data from CoinGecko: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Error: received status code %d from CoinGecko", resp.StatusCode)
	}
	
	var result coingeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("Error decoded JSON: %w", err)
	}


	switch currency {
		case "brl":
			return result.Bitcoin.BRL, nil
		case "usd":
			return result.Bitcoin.USD, nil
		default:
			return 0, fmt.Errorf("Unsupported currency: %s", currency)
	}

}
