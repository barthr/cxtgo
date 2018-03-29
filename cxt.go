package cxtgo

import (
	"context"
)

const (
	// Version indicates cxtgo version
	Version = 0.1
)

// Compile time helpers to check if the implementation implement the desired interfaces
var (
// _ = (FullExchange)(&Binance{})
)

// Lotter is the interface for converting amounts to lot sizes
type Lotter interface {
	AmountToLots(float64) float64
}

// Response is the original response from the exchange api
type Response struct {
	Original []byte
}

// PublicExchange are the public available calls for an exchange
type PublicExchange interface {
	LoadMarkets(ctx context.Context, reload ...bool) (map[Symbol]MarketInfo, error)
	FetchMarkets(ctx context.Context) (Response, error)
	FetchTicker(ctx context.Context) (Response, error)
	FetchTickers(ctx context.Context) (Response, error)
	FetchOrderBook(ctx context.Context) (Response, error)
	FetchOHLCV(ctx context.Context) (Response, error)
	FetchTrades(ctx context.Context) (Response, error)
}

// PrivateExchange are the private available calls for an exchange
type PrivateExchange interface {
	FetchBalance(ctx context.Context) (Response, error)
	CreateOrder(ctx context.Context) (Response, error)
	CancelOrder(ctx context.Context) (Response, error)
	CancelAllOrders(ctx context.Context) (Response, error)
	FetchOrder(ctx context.Context) (Response, error)
	FetchOrders(ctx context.Context) (Response, error)
	FetchOpenOrders(ctx context.Context) (Response, error)
	FetchClosedOrders(ctx context.Context) (Response, error)
	FetchMyTrades(ctx context.Context) (Response, error)
	Deposit(ctx context.Context) (Response, error)
	Withdraw(ctx context.Context) (Response, error)
}

type FullExchange interface {
	Info() Base

	PublicExchange
	PrivateExchange
}

type MarketError struct {
	Exchange string
	Cause    error
}
