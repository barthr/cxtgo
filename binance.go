package cxtgo

import (
	"context"
	"net/http"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo/base"
)

// BinanceWithHTTPClient sets the http client for binance to use
func BinanceWithHTTPClient(client *http.Client) func(*Binance) {
	return func(b *Binance) {
		b.client.HTTPClient = client
	}
}

// BinanceOptFunc is the option function for binance
// This can be used to define settings for the binance exchange
type BinanceOptFunc func(*Binance)

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	base.Exchange

	client *binance.Client
}

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

func (b *Binance) LoadMarkets(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchMarkets(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchTicker(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchTickers(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOrderBook(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOHLCV(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchTrades(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchBalance(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) CreateOrder(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) CancelOrder(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOrder(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOrders(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchOpenOrders(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchClosedOrders(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) FetchMyTrades(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) Deposit(ctx context.Context) (Response, error) {
	panic("not implemented")
}

func (b *Binance) Withdraw(ctx context.Context) (Response, error) {
	panic("not implemented")
}
