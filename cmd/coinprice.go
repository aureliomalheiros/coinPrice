package cmd

import (
	"fmt"
	"log"
	"os"
	"encoding/json"

	"github.com/aureliomalheiros/coinprice/internal/price"
	"github.com/spf13/cobra"
)

var (
	usdFlag bool
	brlFlag bool
	jsonFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "coinprice",
	Short: "Get current Bitcoin price in BRL",
	Run: func(cmd *cobra.Command, args []string) {
		if !usdFlag && !brlFlag {
			fmt.Println("Please specify a currency using --usd or --brl flag")
			cmd.Help()
			os.Exit(1)
		}

		result := make(map[string]float64)

		if brlFlag {
			priceBRL, err := price.GetBitcoinPriceBRL("brl")
			if err != nil {
				log.Fatalf("Error fetching BRL price: %v", err)
			}
			result["brl"] = priceBRL
		}

		if usdFlag {
			priceUSD, err := price.GetBitcoinPriceBRL("usd")
			if err != nil {
				log.Fatalf("Error fetching USD price: %v", err)
			}
			result["brl"] = priceUSD
		}

		if brlFlag && usdFlag {
			priceBRL, err := price.GetBitcoinPriceBRL("brl")
			if err != nil {
				log.Fatalf("Error fetching BRL price: %v", err)
			}
			priceUSD, err := price.GetBitcoinPriceBRL("usd")
			if err != nil {
				log.Fatalf("Error fetching USD price: %v", err)
			}
			result["brl"] = priceBRL
			result["usd"] = priceUSD
		}

		if jsonFlag {
			jsonData, _ := json.Marshal(result)
			fmt.Println(string(jsonData))
		} else {
			for k, v := range result {
				if k == "brl" {
					fmt.Printf("Bitcoin price in %s: R$ %.2f\n", k, v)
				} else if k == "usd" {
					fmt.Printf("Bitcoin price in %s: $ %.2f\n", k, v)
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
	rootCmd.Flags().BoolVarP(&usdFlag, "usd", "u", false, "Get Bitcoin price in USD")
	rootCmd.Flags().BoolVarP(&brlFlag, "brl", "b", false, "Get Bitcoin price in BRL")
	rootCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
}
