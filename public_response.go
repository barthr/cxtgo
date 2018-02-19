package cxtgo

import "github.com/barthr/cxtgo/base"

type OrderBookResponse struct {
	Response
	Orderbook base.OrderBook
}
