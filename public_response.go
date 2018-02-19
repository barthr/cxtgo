package cxtgo

import "github.com/barthr/cxtgo/base"

type OrderBookResponse struct {
	Response
	Orderbook base.OrderBook
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
