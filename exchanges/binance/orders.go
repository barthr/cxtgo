package binance

import (
	"context"

	"github.com/barthr/cxtgo"
)

func (b *Binance) FetchBalance(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) CreateOrder(ctx context.Context) (cxtgo.Response, error) {
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

func (b *Binance) FetchOrder(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOrders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOpenOrders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchClosedOrders(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchMyTrades(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) Deposit(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) Withdraw(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}
