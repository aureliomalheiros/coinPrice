package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aureliomalheiros/coinprice/internal/price"
	"github.com/spf13/cobra"
)

var (
	usdFlag  bool
	brlFlag  bool
	jsonFlag bool
	coinFlag string
)

var rootCmd = &cobra.Command{
	Use:   "coinprice [coin]",
	Short: "Get current Crypto price in BRL and USD",
	Run: func(cmd *cobra.Command, args []string) {

		if !brlFlag && !usdFlag {
			usdFlag = true
		}

		coin := "bitcoin"
		if coinFlag != "" {
			coin = coinFlag
		}

		result := make(map[string]float64)

		if brlFlag {
			priceBRL, err := price.GetCryptoPrice(coin, "brl")
			if err != nil {
				log.Fatalf("Error fetching BRL price for %s: %v", coin, err)
			}
			result["brl"] = priceBRL
		}

		if usdFlag {
			priceUSD, err := price.GetCryptoPrice(coin, "usd")
			if err != nil {
				log.Fatalf("Error fetching USD price for %s: %v", coin, err)
			}
			result["usd"] = priceUSD
		}

		if jsonFlag {
			jsonData, err := json.Marshal(result)
			if err != nil {
				log.Fatalf("Error generating JSON: %v", err)
			}
			fmt.Println(string(jsonData))
		} else {
			for currency, value := range result {
				switch currency {
				case "brl":
					fmt.Printf("%s price in BRL: R$ %.2f\n", coin, value)
				case "usd":
					fmt.Printf("%s price in USD: $ %.2f\n", coin, value)
				}
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&usdFlag, "usd", "u", false, "Get price in USD")
	rootCmd.Flags().BoolVarP(&brlFlag, "brl", "b", false, "Get price in BRL")
	rootCmd.Flags().BoolVarP(&jsonFlag, "json", "j", false, "Output in JSON format")
	rootCmd.Flags().StringVarP(&coinFlag, "coin", "c", "", "Specify cryptocurrency  (Read documentation) - default: bitcoin")
}
