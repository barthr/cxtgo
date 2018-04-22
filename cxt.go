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
	Market(ctx context.Context, symbol Symbol, params ...Params) (MarketInfo, error)
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
	// Balances returns the balances from the exchange
	Balances(ctx context.Context, params ...Params) (Balances, error)
	// MyTrades returns the trades made by that account
	MyTrades(ctx context.Context, params ...Params) ([]Trade, error)
	// FreeBalance returns the free balance in the account
	FreeBalance(ctx context.Context, params ...Params) (PartialBalances, error)
	// UsedBalance returns the used balance (in trade) in the account
	UsedBalance(ctx context.Context, params ...Params) (PartialBalances, error)
	// TotalBalance returns the total used + free balance in the account
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
	CancelOrder(ctx context.Context, ID string, symbol *Symbol, params ...Params) (Response, error)
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
