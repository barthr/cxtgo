package base

// Orderbook is a definition for an orderbook
type Orderbook interface {
	Symbol()
	Get()
	Head()
}
