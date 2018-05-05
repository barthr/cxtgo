package cxtgo

import "time"

// Side defines the side for the order (f.e. Sell or Buy).
type Side int

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
	sidesString := [...]string{
		"unknown",
		"sell",
		"buy",
	}
	return sidesString[s]
}

// OrderStatus represents the state of an order.
type OrderStatus int

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
	statusString := [...]string{
		"unknown",
		"open",
		"closed",
		"canceled",
	}
	return statusString[os]
}

// OrderType represents the type of the order.
type OrderType int

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
	orderTypeString := [...]string{
		"unknown",
		"market",
		"limit",
	}
	return orderTypeString[ot]
}

// Order represents an order at an exchange.
type Order struct {
	Symbol    Symbol
	ID        string
	Timestamp int64
	Datetime  time.Time
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
