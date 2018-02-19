package cxtgo

import "context"

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
	LoadMarkets(ctx context.Context) Response
	FetchMarkets(ctx context.Context) Response
	FetchTicker(ctx context.Context) Response
	FetchTickers(ctx context.Context) Response
	FetchOrderBook(ctx context.Context) Response
	FetchOHLCV(ctx context.Context) Response
	FetchTrades(ctx context.Context) Response
}

// PrivateExchange are the private available calls for an exchange
type PrivateExchange interface {
	FetchBalance(ctx context.Context) Response
	CreateOrder(ctx context.Context) Response
	CancelOrder(ctx context.Context) Response
	FetchOrder(ctx context.Context) Response
	FetchOrders(ctx context.Context) Response
	FetchOpenOrders(ctx context.Context) Response
	FetchClosedOrders(ctx context.Context) Response
	FetchMyTrades(ctx context.Context) Response
	Deposit(ctx context.Context) Response
	Withdraw(ctx context.Context) Response
}
