package cmd

import (
	"fmt"

	"github.com/leb4r/trader-go/internal"
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/spf13/cobra"
)

var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "make a trade",
	RunE:  tradeAction,
}

func init() {
	rootCmd.AddCommand(tradeCmd)
}

func tradeAction(cmd *cobra.Command, args []string) error {
	config := coinbasepro.ClientConfig{}

	client := internal.CreateCoinbaseApiClient(config)
	accounts, err := client.GetAccounts()
	if err != nil {
		return err
	}

	for _, a := range accounts {
		fmt.Println(a.Balance)
	}
	return nil
}
