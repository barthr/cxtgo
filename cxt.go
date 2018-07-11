package cxtgo

import (
	"context"
	"math"
)

const (
	// Version indicates cxtgo version
	Version = 0.1
)

var (
	exchanges = [...]ExchangeName{
		"binance",
		"bitmex",
		"hitbtc",
	}
)

// AmountToLots converts an amount to a lot sized amount according to the precisions in `MarketInfo`.
func AmountToLots(info MarketInfo, amount float64) float64 {
	return math.Trunc(math.Floor(amount/info.Lot)*info.Lot*math.Pow10(info.Precision.Amount)) / math.Pow10(info.Precision.Amount)
}

// PublicAPI are the public available calls for an exchange.
type PublicAPI interface {
	// Markets return the market information for an exchange.
	Markets(ctx context.Context, params ...Params) (MarketInfos, error)
	// Currencies return the currencies used by the exchange.
	Currencies(ctx context.Context, params ...Params) (Currencies, error)
	// Ticker returns the ticker information for a given symbol.
	Ticker(ctx context.Context, symbol Symbol, params ...Params) (Ticker, error)
	// Tickers returns all the ticker information for the given symbols.
	Tickers(ctx context.Context, symbols Symbols, params ...Params) (Tickers, error)
	// Orderbook returns order book information for a given symbol.
	OrderBook(ctx context.Context, symbol Symbol, params ...Params) (Orderbook, error)
	// OHLCV returns the Open high low close volume infromation for a symbol.
	OHLCV(ctx context.Context, params ...Params) error
	// Trades return all the trades for a given symbol.
	Trades(ctx context.Context, params ...Params) ([]Trade, error)
}

// AccountAPI are the private user api calls for an exchange.
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
	Deposit(ctx context.Context, params ...Params) error
	// Currently unused
	Withdraw(ctx context.Context, params ...Params) error
}

// OrderAPI are all the calls for creating updating and fetching orders.
type OrderAPI interface {
	// LimitOrder creates a limit order at the exchange.
	LimitOrder(ctx context.Context, symbol Symbol, side Side, offer Offer, params ...Params) error
	// MarketOrder creates an market order at the exchange.
	MarketOrder(ctx context.Context, symbol Symbol, side Side, amount float64, params ...Params) error
	// CancelOrder cancels the order for the given id.
	CancelOrder(ctx context.Context, ID string, symbol *Symbol, params ...Params) error
	// Order returns the order information for the given order id.
	Order(ctx context.Context, params ...Params) (Order, error)
	// Orders returns the order information for all the orders for that symbol.
	Orders(ctx context.Context, params ...Params) ([]Order, error)
	// OpenOrders is like orders but only returning orders which are open.
	OpenOrders(ctx context.Context, params ...Params) ([]Order, error)
	// ClosedOrders is like orders but only returning orders which are closed.
	ClosedOrders(ctx context.Context, params ...Params) ([]Order, error)
}

// StreamingAPI defines the streaming api endpoints for an exchange.
type StreamingAPI interface {
	OrderbookStreamer
	TickerStreamer
	TradeStreamer
}

// Exchange defines all the api calls for an exchange.
type Exchange interface {
	Info() Base

	PublicAPI
	AccountAPI
	OrderAPI
}
