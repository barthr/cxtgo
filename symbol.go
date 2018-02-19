package cxtgo

import (
	"sort"

	"github.com/xtgo/set"
)

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

// SymbolSet returns Symbols with the duplicates removed
func SymbolSet(data Symbols) Symbols {
	sort.Sort(data)
	n := set.Uniq(data)
	return data[:n]
}

// Symbols is a container type for multiple symbols
type Symbols []Symbol

// Len returns the length of s
func (s Symbols) Len() int {
	return len(s)
}

// Less defines the order of s for sorting
func (s Symbols) Less(i int, j int) bool {
	return s[i].String() < s[j].String()
}

// Swap swaps elements in s
func (s Symbols) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}
