package validates

import (
	"fmt"
	"sort"
	"strings"
)

var validCurrencies = map[string]bool{
	"brl": true,
	"usd": true,
}

var validCoins = map[string]bool{
	"bitcoin":       true,
	"ethereum":      true,
	"aave":          true,
	"dogecoin":      true,
	"solana":        true,
	"litecoin":      true,
	"ripple":        true,
	"cardano":       true,
	"polkadot":      true,
	"chainlink":     true,
	"uniswap":       true,
	"avalanche":     true,
	"cosmos":        true,
	"stellar":       true,
	"monero":        true,
	"tron":         	true,
	"algorand":     true,
	"vechain":      true,
	"filecoin":     true,
	"bitcoin-cash": true,
	"dash":         true,
	"zcash":        true,
	"tezos":        true,
	"neo":          true,
	"elrond":       true,
	"hedera":       true,
	"fantom":       true,
	"decentraland": true,
	"the-sandbox":  true,
	"axie-infinity": true,
	"chiliz":       true,
	"enjin-coin":   true,
	"theta":        true,
}

func GetValidCoins() map[string]bool {
	return validCoins
}

func ValidateInput(coin, currency string) error {
	if !validCurrencies[currency] {
		return fmt.Errorf("invalid currency: %s. Valid options are: brl, usd", currency)
	}

	if !validCoins[coin] {
		var coins []string
		for c := range validCoins {
			coins = append(coins, c)
		}
		sort.Strings(coins)

		return fmt.Errorf("invalid coin: %s. Available cryptocurrencies:\n%s", coin, strings.Join(coins, "\n"))
	}

	return nil
}
