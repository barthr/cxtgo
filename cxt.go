package cxtgo

import (
	"context"
)

const (
	// Version indicates cxtgo version
	Version = 0.1
)

// Lotter is the interface for converting amounts to lot sizes
type Lotter interface {
	AmountToLots(Symbol, float64) float64
}

// Response is the original response from the exchange api
type Response struct {
	Original []byte
}

// PublicAPI are the public available calls for an exchange
type PublicAPI interface {
	Markets(ctx context.Context) (map[Symbol]MarketInfo, error)
	Ticker(ctx context.Context) (Response, error)
	Tickers(ctx context.Context) (Response, error)
	OrderBook(ctx context.Context, symbol Symbol) (Orderbook, error)
	OHLCV(ctx context.Context) (Response, error)
	Trades(ctx context.Context) (Response, error)
}

// AccountAPI are the private user api calls for an exchange
type AccountAPI interface {
	Balance(ctx context.Context) (Response, error)
	MyTrades(ctx context.Context) (Response, error)
	Deposit(ctx context.Context) (Response, error)
	Withdraw(ctx context.Context) (Response, error)
}

// OrderAPI are all the calls for creating updating and fetching orders
type OrderAPI interface {
	LimitOrder(ctx context.Context, symbol Symbol, side Side, amount, price float64, params Params) (Response, error)
	MarketOrder(ctx context.Context, symbol Symbol, side Side, amount, params Params) (Response, error)
	CancelOrder(ctx context.Context) (Response, error)
	Order(ctx context.Context) (Response, error)
	Orders(ctx context.Context) (Response, error)
	OpenOrders(ctx context.Context) (Response, error)
	ClosedOrders(ctx context.Context) (Response, error)
}

// Exchange defines all the api calls for an exchange
type Exchange interface {
	Info() Base
	Reset()

	PublicAPI
	AccountAPI
	OrderAPI
}
