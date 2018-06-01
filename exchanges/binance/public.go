package binance

import (
	"context"
	"errors"

	"github.com/barthr/cxtgo"
)

// Markets loads all the markets from binance
func (b *Binance) Markets(ctx context.Context, params ...cxtgo.Params) (cxtgo.MarketInfos, error) {
	b.base.Ratelimit.Take()
	return nil, cxtgo.E(cxtgo.ExchangeName("binance"), cxtgo.Op("public.Markets"), errors.New("it's nothing"), cxtgo.ExchangeNotAvailable)
}

func (b *Binance) initMarkets() error {
	var err error
	b.once.Do(func() {
		_, err = b.Markets(context.Background())
	})
	return err
}

// Ticker loads a single ticker for the symbol `s` from binance.
func (b *Binance) Ticker(ctx context.Context, s cxtgo.Symbol) (cxtgo.Ticker, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Ticker{}, err
	}
	panic("not implemented")
}

// Tickers loads all the tickers from binance.
func (b *Binance) Tickers(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) OrderBook(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) OHLCV(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) Trades(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}
