package internal

import (
	ws "github.com/gorilla/websocket"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

var (
	DefaultCoinbaseApiURL        = "https://api-public.exchange.coinbase.com"
	DefaultCoinbaseSandboxApiURL = "https://api-public.sandbox.exchange.coinbase.com"

	DefaultCoinbaseWsURL        = "wss://ws-feed.exchange.coinbase.com"
	DefaultCoinbaseSandboxWsURL = "wss://ws-feed-public.sandbox.exchange.coinbase.com"
)

// CreateCoinbaseProClient returns a client that can be used to communicate with coinbase pro
func CreateCoinbaseApiClient(config coinbasepro.ClientConfig) *coinbasepro.Client {
	client := coinbasepro.NewClient()

	// use default url if BaseURL is not set
	// if config.BaseURL == "" {
	// 	config.BaseURL = DefaultCoinbaseApiURL
	// }

	// client.UpdateConfig(&config)
	return client
}

// CreateCoinbaseWebsocketConn returns a websocket connection subscribed to a set of channels for each productId
func CreateCoinbaseWebsocketConn(feedUrl string, channelNames, productIds []string) (*ws.Conn, error) {
	var wsDialer ws.Dialer

	// create dialer using feedUrl
	wsConn, _, err := wsDialer.Dial(feedUrl, nil)
	if err != nil {
		return wsConn, err
	}

	// create channels slice
	channels := getChannels(channelNames, productIds)

	// set which channels to subscribe to
	subscribe := coinbasepro.Message{
		Type:     "subscribe",
		Channels: channels,
	}

	// subscribe to channels
	if err := wsConn.WriteJSON(subscribe); err != nil {
		return wsConn, err
	}

	return wsConn, nil
}

// HandleWebsocketMessages loops continuously executing a specified function to handle the message
func HandleWebsocketMessages(wsConn ws.Conn, onMessage func(message coinbasepro.Message) error) error {
	for {
		message := coinbasepro.Message{}
		if err := wsConn.ReadJSON(&message); err != nil {
			return err
		}

		if err := onMessage(message); err != nil {
			return err
		}
	}
}

// GetCoinbaseWsUrl returns the endpoint to be used for websocket clients, pass sandbox true to return sandbox endpoints
func GetCoinbaseWsUrl(sandbox bool) string {
	if sandbox {
		return DefaultCoinbaseSandboxWsURL
	} else {
		return DefaultCoinbaseWsURL
	}
}

func getChannels(channelNames, productIds []string) []coinbasepro.MessageChannel {
	channels := []coinbasepro.MessageChannel{}

	for i := range channelNames {
		channels = append(channels, coinbasepro.MessageChannel{
			Name:       channelNames[i],
			ProductIds: productIds,
		})
	}

	return channels
}
