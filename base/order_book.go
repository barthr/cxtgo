package base

// Orderbook is a definition for an orderbook
type OrderBook interface {
	Symbol()
	Get()
	Head()
}
