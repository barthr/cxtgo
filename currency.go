package cxtgo

// Currency represents a currency from the exchange.
type Currency struct {
	ID   string
	Code string
}

// Currencies represent multiple currencies from the exchange.
type Currencies map[string]Currency
