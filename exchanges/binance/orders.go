package binance

import (
	"context"

	"github.com/barthr/cxtgo"
)

func (b *Binance) LimitOrder(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) MarketOrder(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) CancelOrder(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) CancelAllOrders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) Order(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) Orders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) OpenOrders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) ClosedOrders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}
