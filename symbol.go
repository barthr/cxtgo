package cxtgo

// Symbol represents a combination of two currencies (BTCUSD)
// this type is immutable
type Symbol struct {
	delim  string // optional
	first  string
	second string
}

// First returns the first currency from the Symbol
func (p Symbol) First() string {
	return p.first
}

// Second returns the second currency from the Symbol
func (p Symbol) Second() string {
	return p.second
}

// Delim returns the delim (this is optional to use the symbol)
func (p Symbol) Delim() string {
	return p.delim
}

// Reverse reverse the currencies (fe. BTCUSD to USDBTC or BTC/USD to USD/BTC)
func (p Symbol) Reverse() Symbol {
	return Symbol{
		first:  p.second,
		second: p.first,
		delim:  p.delim,
	}
}

// String returns the string version of the Symbol
func (p Symbol) String() string {
	return p.first + p.delim + p.second
}

// NewSymbol create's a new Pair from the given string
func NewSymbol(from, to string, delim ...string) Symbol {
	delimiter := ""
	if len(delim) != 0 {
		delimiter = delim[0]
	}
	return Symbol{
		delim:  delimiter,
		first:  from,
		second: to,
	}
}
