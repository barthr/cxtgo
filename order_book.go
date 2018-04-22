package cxtgo

import (
	"sync"
)

// BookOrder defines an offer from the order book
type BookOrder struct {
	Price  float64
	Amount float64
}

// Bid represents a bid offer from the order book
type Bid BookOrder

// Ask represents an ask offer from the order book
type Ask BookOrder

// Orderbook is a definition for an orderbook
type Orderbook interface {
	Symbol() Symbol
	Head(n int) Summary
	Spread() float64
	Depth() Summary
	Error() error
}

// Summary defines a view of the order book.
// The bids are sorted descending and the ask ascending.
type Summary struct {
	Bids []Bid
	Asks []Ask
}

type ConcurrentOrderbook struct {
	Symbol
	sync.RWMutex
}

func (co *ConcurrentOrderbook) Head(n int) Summary {
	return Summary{}
}
func (co *ConcurrentOrderbook) Spread() float64 {
	return -1.0
}
func (co *ConcurrentOrderbook) Depth() Summary {
	return Summary{}
}
