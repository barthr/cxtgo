package cxtgo

// OrderBook is a definition for an orderbook
type OrderBook interface {
	Symbol() Symbol
	Get()
	Head()
}
