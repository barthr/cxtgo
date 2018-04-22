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
	Markets(ctx context.Context, params ...Params) (MarketInfos, error)
	Currencies(ctx context.Context, params ...Params) (Currencies, error)
	Ticker(ctx context.Context, symbol Symbol, params ...Params) (Ticker, error)
	Tickers(ctx context.Context, symbols []Symbol, params ...Params) (Tickers, error)
	OrderBook(ctx context.Context, symbol Symbol, params ...Params) (Orderbook, error)
	OHLCV(ctx context.Context, params ...Params) (Response, error)
	Trades(ctx context.Context, params ...Params) ([]Trade, error)
}

// AccountAPI are the private user api calls for an exchange
type AccountAPI interface {
	Balances(ctx context.Context, params ...Params) (Balances, error)
	MyTrades(ctx context.Context, params ...Params) ([]Trade, error)
	FreeBalance(ctx context.Context, params ...Params) (PartialBalances, error)
	UsedBalance(ctx context.Context, params ...Params) (PartialBalances, error)
	TotalBalance(ctx context.Context, params ...Params) (PartialBalances, error)
	// Currently unused
	Deposit(ctx context.Context, params ...Params) (Response, error)
	// Currently unused
	Withdraw(ctx context.Context, params ...Params) (Response, error)
}

// OrderAPI are all the calls for creating updating and fetching orders
type OrderAPI interface {
	LimitOrder(ctx context.Context, symbol Symbol, side Side, amount, price float64, params ...Params) (Response, error)
	MarketOrder(ctx context.Context, symbol Symbol, side Side, amount float64, params ...Params) (Response, error)
	CancelOrder(ctx context.Context, params ...Params) (Response, error)
	Order(ctx context.Context, params ...Params) ([]Order, error)
	Orders(ctx context.Context, params ...Params) ([]Order, error)
	OpenOrders(ctx context.Context, params ...Params) ([]Order, error)
	ClosedOrders(ctx context.Context, params ...Params) ([]Order, error)
}

// Exchange defines all the api calls for an exchange
type Exchange interface {
	Info() Base
	Reset()

	PublicAPI
	AccountAPI
	OrderAPI
}
