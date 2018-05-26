package binance

import (
	"fmt"
	"strings"

	"github.com/barthr/cxtgo"
	"github.com/gorilla/websocket"
)

const (
	streamURL = "wss://stream.binance.com:9443"
)

func (b *Binance) StreamTicker(onUpdate func(t cxtgo.Ticker), onError func(err error), opts ...cxtgo.StreamOpt) error {
	config := new(cxtgo.StreamConfig)
	for _, opt := range opts {
		opt(config)
	}
	if config.Symbol.String() == "" {
		return cxtgo.WrapError(cxtgo.SymbolNotFoundError{}, "binance", errSymbolNotFound(config.Symbol))
	}
	endpoint := fmt.Sprintf("%s/ws/%s@kline_%s", streamURL, strings.ToLower(config.Symbol.String()), "1m")
	streamLogger := b.logger.WithField("stream", "ticker").WithField("symbol", config.Symbol)
	maxConnectRetries := 0

	var err error
	var conn *websocket.Conn
	for maxConnectRetries < 5 {
		streamLogger.WithField("endpoint", endpoint).Debugln("starting websocket connection with stream")
		conn, _, err = websocket.DefaultDialer.Dial(endpoint, nil)
		if err != nil {
			streamLogger.WithError(err).Debugln("failed connecting to websocket stream")
			if !config.Reconnect {
				return cxtgo.WrapError(cxtgo.StreamUnavailableError{}, "binance", err)
			}
			maxConnectRetries++
		}
		for {
			msgType, _, err := conn.ReadMessage()
			if err != nil {
				streamLogger.WithError(err).Debugln("failed decoding input from stream")
				err = cxtgo.WrapError(cxtgo.StreamError{StreamType: cxtgo.TickerStream}, "binance", err)
				goto end
			}
			if msgType == websocket.TextMessage {
				// spew.Dump(resp)
			} else {
				fmt.Println("disconnected")
			}
		}
	end:
		conn.Close()
		if !config.Reconnect {
			return cxtgo.WrapError(cxtgo.StreamClosedByExchangeError{}, "binance", err)
		}
	}
	return cxtgo.WrapError(cxtgo.StreamClosedByExchangeError{}, "binance", nil)
}

func (b *Binance) StreamOrderbook(onUpdate func(s cxtgo.Summary), onError func(err error), opts ...cxtgo.StreamOpt) error {
	panic("not implemented")
}

func (b *Binance) StreamTrades(onUpdate func(t cxtgo.Trade), onError func(err error), opts ...cxtgo.StreamOpt) error {
	panic("not implemented")
}
