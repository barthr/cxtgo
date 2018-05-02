package binance

import (
	"context"
	"strconv"
	"time"

	"github.com/barthr/cxtgo"
)

func (b *Binance) LimitOrder(ctx context.Context, symbol cxtgo.Symbol, side cxtgo.Side, offer cxtgo.Offer, params ...cxtgo.Params) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	timeInForce := "GTC"
	if len(params) > 0 {
		val, ok := params[0].GetString("timeInForce")
		if ok {
			timeInForce = val
		}
	}
	_, err := b.http.R().SetContext(ctx).SetQueryParams(map[string]string{
		"symbol":      symbol.String(),
		"side":        side.String(),
		"type":        cxtgo.LimitOrder.String(),
		"timeInForce": timeInForce,
		"price":       strconv.FormatFloat(offer.Price, 'f', -1, 64),
		"quantity":    strconv.FormatFloat(offer.Amount, 'f', -1, 64),
		"timestamp":   strconv.FormatInt(time.Now().UnixNano(), 10),
	}).Post("/api/v3/order")

	if err != nil {
		return cxtgo.ExchangeNotAvailableError{
			BaseError: cxtgo.NewError("binance", err),
		}
	}
	return nil
}

func (b *Binance) MarketOrder(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) CancelOrder(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) CancelAllOrders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) Order(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) Orders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) OpenOrders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}

func (b *Binance) ClosedOrders(ctx context.Context) error {
	if err := b.initMarkets(); err != nil {
		return err
	}
	panic("not implemented")
}
