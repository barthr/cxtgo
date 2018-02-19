package cxtgo

import (
	"net/http"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo/base"
)

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	base.Exchange

	client *binance.Client
}

// BinanceWithHTTPClient sets the http client for binance to use
func BinanceWithHTTPClient(client *http.Client) func(*Binance) {
	return func(b *Binance) {
		b.client.HTTPClient = client
	}
}

// BinanceOptFunc is the option function for binance
// This can be used to define settings for the binance exchange
type BinanceOptFunc func(*Binance)

// NewBinance returns an instance of the binance exchange
func NewBinance(config *base.Config, opts ...BinanceOptFunc) *Binance {
	if config == nil {
		return nil
	}
	b := &Binance{
		Exchange: base.Exchange{
			Name:   "binance",
			Config: *config,
		},
		client: binance.NewClient(config.APIKEY, config.APISecret),
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}
