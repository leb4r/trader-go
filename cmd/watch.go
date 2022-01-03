package cmd

import (
	"errors"
	"fmt"

	"github.com/leb4r/trader-go/internal"
	"github.com/leb4r/trader-go/internal/models"
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch ticker feed",
	RunE:  watchAction,
}

var (
	defaultCoinbaseProductIds = []string{"BTC-USD"}
	defaultCoinbaseChannels   = []string{"heartbeat", "ticker"}
)

func init() {
	// coinbase
	watchCmd.PersistentFlags().StringArray("coinbase-product-id", defaultCoinbaseProductIds, "Product ID to watch")
	if err := viper.BindPFlag("coinbaseProductIds", watchCmd.PersistentFlags().Lookup("coinbase-product-id")); err != nil {
		internal.ThrowError(err)
	}

	watchCmd.PersistentFlags().StringArray("coinbase-channel", defaultCoinbaseChannels, "The channel to subscribe to")
	if err := viper.BindPFlag("coinbaseChannels", watchCmd.PersistentFlags().Lookup("coinbase-channel")); err != nil {
		internal.ThrowError(err)
	}

	watchCmd.PersistentFlags().Bool("coinbase-sandbox", false, "Whether or not to use Coinbase's Sandbox")
	if err := viper.BindPFlag("coinbaseSandbox", watchCmd.PersistentFlags().Lookup("coinbase-sandbox")); err != nil {
		internal.ThrowError(err)
	}

	rootCmd.AddCommand(watchCmd)
}

func watchAction(cmd *cobra.Command, args []string) error {
	// get products, error if it's empty
	products := viper.GetStringSlice("coinbaseProductIds")

	if len(products) == 0 {
		return errors.New("there are no products to watch")
	}

	// get channels for the products
	channels := viper.GetStringSlice("coinbaseChannels")

	if len(channels) == 0 {
		return errors.New("there are no channels to watch")
	}

	// determine which environment to use
	var feedUrl string
	if viper.GetBool("coinbaseSandbox") {
		fmt.Println("Using the Coinbase Sandbox!")
		feedUrl = internal.DefaultCoinbaseSandboxWsURL
	} else {
		feedUrl = internal.DefaultCoinbaseWsURL
	}

	// create the websockets connection subscribed to specific channels
	wsConn, err := internal.CreateCoinbaseWebsocketConn(feedUrl, channels, products)
	if err != nil {
		return err
	}

	// watch for messages, printing them out when received
	if err := internal.HandleWebsocketMessages(*wsConn, watchMessageHandler); err != nil {
		return err
	}

	return nil
}

func watchMessageHandler(message coinbasepro.Message) error {

	switch message.Type {
	case "ticker":
		var p = &models.Price{
			Pair:   message.ProductID,
			Amount: message.Price,
		}
		if err := p.Create(dbHandler); err != nil {
			return err
		}
	case "heartbeat":
		sugar.Infow("New heartbeat message received")
	case "error":
		return errors.New(message.Message)
	}

	return nil
}
