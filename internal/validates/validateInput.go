package validates

import (
	"fmt"
	"sort"
	"strings"
)

func ValidateInput(coin, currency string) error {
	validCurrencies := map[string]bool{
		"brl": true,
		"usd": true,
	}
	
	validCoin := map[string]bool{
		"bitcoin": true,
		"ethereum": true,
		"aave": true,
		"solana": true,
		"dogecoin": true,
		"litecoin": true,
		"ripple": true,
		"cardano": true,
		"polkadot": true,
		"chainlink": true,
		"uniswap": true,
		"avalanche": true,
		"cosmos": true,
		"stellar": true,
		"monero": true,
		"tron": true,
		"algorand": true,
		"vechain": true,
		"filecoin": true,
		"bitcoin-cash": true,
		"dash": true,
		"zcash": true,
		"tezos": true,
		"neo": true,
		"elrond": true,
		"hedera": true,
		"fantom": true,
		"decentraland": true,
		"the-sandbox": true,
		"axie-infinity": true,
		"chiliz": true,
		"enjin-coin": true,
		"theta": true,			
	}

	if !validCurrencies[currency] {
		return fmt.Errorf("Invalid currency: %s. Valid options are: brl, usd", currency)
	}

	if !validCoin[coin] {
		var coins []string
		for c := range validCoin {
			coins = append(coins, c)
		}
		sort.Strings(coins)

		return fmt.Errorf("Invalid coin: %s. Available cryptocurrencies:\n%s", coin, strings.Join(coins, "\n"))
	}

	return nil
}
