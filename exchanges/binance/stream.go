package binance

import (
	"github.com/barthr/cxtgo"
)

const (
	streamURL = "wss://stream.binance.com:9443"
)

func (b *Binance) StreamTicker(onUpdate func(t cxtgo.Ticker), onError func(err error), opts ...cxtgo.StreamOpt) error {
	return cxtgo.Err(cxtgo.ExchangeName("binance"), cxtgo.Op("streamingAPI.StreamTicker"), cxtgo.NotSupported)
}

func (b *Binance) StreamOrderbook(onUpdate func(s cxtgo.Summary), onError func(err error), opts ...cxtgo.StreamOpt) error {
	return cxtgo.Err(cxtgo.ExchangeName("binance"), cxtgo.Op("streamingAPI.StreamOrderbook"), cxtgo.NotSupported)
}

func (b *Binance) StreamTrades(onUpdate func(t cxtgo.Trade), onError func(err error), opts ...cxtgo.StreamOpt) error {
	return cxtgo.Err(cxtgo.ExchangeName("binance"), cxtgo.Op("streamingAPI.StreamTrades"), cxtgo.NotSupported)
}
