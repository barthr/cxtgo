package cxtgo

import "github.com/barthr/cxtgo/exchange"

type OrderBookResponse struct {
	Response
	Orderbook exchange.OrderBook
}

type TickerResponse struct {
	Response
}

type TickersResponse struct {
	Response
}

type MarketResponse struct {
	Response
}
