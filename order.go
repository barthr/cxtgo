package cxtgo

import "time"

// Side defines the side for the order (f.e. Sell or Buy).
type Side uint8

const (
	// Unknown side for the zero value of side.
	Unknown Side = iota
	// Sell defines the sell side.
	Sell
	// Buy defines the buy side.
	Buy
)

// String returns the string value of s.
func (s Side) String() string {
	sidesString := map[Side]string{
		Unknown: "unknown",
		Sell:    "sell",
		Buy:     "buy",
	}
	return sidesString[s]
}

// OrderStatus represents the state of an order.
type OrderStatus uint8

const (
	// UnknownStatus status for the zero value of order status.
	UnknownStatus OrderStatus = iota
	// OrderOpen represents the open order status.
	OrderOpen
	// OrderClosed represents the closed order status.
	OrderClosed
	// OrderCanceled represents the canceled order status.
	OrderCanceled
)

// String returns the string value of os
func (os OrderStatus) String() string {
	statusString := map[OrderStatus]string{
		UnknownStatus: "unknown",
		OrderOpen:     "open",
		OrderClosed:   "closed",
		OrderCanceled: "canceled",
	}
	return statusString[os]
}

// OrderType represents the type of the order.
type OrderType uint8

const (
	// UnknownOrderType represents the zero value of the order type.
	UnknownOrderType OrderType = iota
	// MarketOrder defines the market order type.
	MarketOrder
	// LimitOrder defines the limit order type.
	LimitOrder
)

// String returns the string value of ot.
func (ot OrderType) String() string {
	orderTypeString := map[OrderType]string{
		UnknownOrderType: "unknown",
		MarketOrder:      "market",
		LimitOrder:       "limit",
	}
	return orderTypeString[ot]
}

// Order represents an order at an exchange.
type Order struct {
	Symbol    Symbol
	ID        string
	Timestamp time.Time
	Status    OrderStatus
	Type      OrderType
	Price     float64
	Cost      float64
	Amount    float64
	Filled    float64
	Remaining float64
	Fee       float64
	Raw       []byte
}
