package cxtgo

// Balance represents a balance for a currency
type Balance struct {
	Asset string
	Free  float64
	Used  float64
	Total float64
}

// Balances represent all the balances from an exchange.
// Currency mapped to the actual balance.
type Balances map[string]Balance
