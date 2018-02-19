package cxtgo

import (
	"context"
	"strings"

	"go.uber.org/ratelimit"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo/exchange"
)

const binanceReqPerMin = 1200

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	base   *exchange.Base
	client *binance.Client

	rl ratelimit.Limiter
}

// NewBinance returns an instance of the binance exchange
func NewBinance(opts ...exchange.Opt) *Binance {
	binanceOpts := []exchange.Opt{
		exchange.WithName("Binance"),
		exchange.WithUserAgent("cxt/0.1"),
		exchange.WithVersion("v3"),
	}
	binanceOpts = append(binanceOpts, opts...)

	ex := exchange.NewBase(binanceOpts...)
	b := &Binance{
		base:   ex,
		client: binance.NewClient(ex.APIKEY, ex.APISecret),
		rl:     ratelimit.New(binanceReqPerMin / 60),
	}
	return b
}

func (b *Binance) Info() exchange.Base {
	return *b.base
}

func (b *Binance) LoadMarkets(ctx context.Context) (map[exchange.Symbol]exchange.MarketInfo, error) {
	b.rl.Take()
	info, err := b.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, NetworkError{ExchangeError{"binance", err}}
	}

	marketInfos := map[exchange.Symbol]exchange.MarketInfo{}

	for _, symbol := range info.Symbols {
		internalSymbol := exchange.NewSymbol(symbol.BaseAsset, symbol.QuoteAsset)
		marketInfos[internalSymbol] = exchange.MarketInfo{
			ID:     strings.ToLower(symbol.Symbol),
			Base:   symbol.BaseAsset,
			Quote:  symbol.QuoteAsset,
			Symbol: internalSymbol,
			Active: true,
			Precision: exchange.MarketPrecision{
				Price:  8,
				Amount: 8,
				Cost:   8,
			},
			Limits: exchange.MarketLimit{
				Price:  exchange.MinMax{},
				Amount: exchange.MinMax{},
			},
			// todo raw
		}
	}
	// copy the map but return a unmodifiable version
	for key, value := range marketInfos {
		b.base.Market[key] = value
	}
	return marketInfos, nil
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

func (b *Binance) CancelAllOrders(ctx context.Context) (Response, error) {
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
