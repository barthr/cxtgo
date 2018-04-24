package cxtgo

// Orderbook is a definition for an orderbook
type Orderbook struct {
	// Symbol() Symbol
	// Head(n int) Summary
	// Spread() float64
	// Depth() Summary
	// Error() error
}

// Offer defines an offer from the order book
type Offer struct {
	Price  float64
	Amount float64
}

// Summary defines a view of the order book.
// The bids are sorted descending and the ask ascending.
type Summary struct {
	s    Symbol
	Bids []Bid
	Asks []Ask
}

// Symbol returns the symbol from the summary
func (s Summary) Symbol() Symbol {
	return s.s
}

// Spread calculates the spread between bid and ask.
func (s Summary) Spread() float64 {
	return .1
}

// Head returns n items from the orderbook returning a new summary with the reflected
func (s Summary) Head(n int) (Summary, error) {
	return s, nil
}

// Bid represents a bid offer from the order book
type Bid Offer

// Ask represents an ask offer from the order book
type Ask Offer
