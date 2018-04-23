package cxtgo

// Orderbook is a definition for an orderbook
type Orderbook struct {
	// Symbol() Symbol
	// Head(n int) Summary
	// Spread() float64
	// Depth() Summary
	// Error() error
}

// BookOrder defines an offer from the order book
type Offer struct {
	Price  float64
	Amount float64
}

// Summary defines a view of the order book.
// The bids are sorted descending and the ask ascending.
type Summary struct {
	Bids []Bid
	Asks []Ask
}

// Bid represents a bid offer from the order book
type Bid Offer

// Ask represents an ask offer from the order book
type Ask Offer
