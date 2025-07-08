package list

import (
	"fmt"
	"sort"
	"strings"
	
	"github.com/aureliomalheiros/coinprice/internal/validates"
)

func PrintCoinList() {
	coins := GetCoinList()

	fmt.Println("Available cryptocurrencies:")
	fmt.Println(strings.Repeat("─", 30))
	for i, coin := range coins {
		fmt.Printf("%2d. %s\n", i+1, coin)
	}
	fmt.Println("\nUse --coin <nameCoin> to get prices")
	fmt.Println("\nWhen using --list all commands are ignored, including --coin")
}

func GetCoinList() []string {
	validCoins := validates.GetValidCoins()
	
	var coins []string
	for coin := range validCoins {
		coins = append(coins, coin)
	}
	sort.Strings(coins)
	return coins
}
