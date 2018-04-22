package cxtgo

// MarketPrecision defines the precision for a given pair, how many decimals for the amount and prices
type MarketPrecision struct {
	Base   int
	Quote  int
	Amount int
	Price  int
}

// MinMax defines the minimum and maximum value
type MinMax struct {
	Min float64
	Max float64
}

// MarketLimit defines the limits for a market, what min and max amounts/prices and costs.
type MarketLimit struct {
	Amount MinMax
	Price  MinMax
	Cost   MinMax
}

// MarketInfo defines all the info for a market (given pair). Things likes what kind of maker and taker fee etc.
type MarketInfo struct {
	ID        string
	Base      string
	Quote     string
	Symbol    Symbol
	Active    bool
	Precision MarketPrecision
	Limits    MarketLimit
	Lot       float64
	Taker     float64
	Maker     float64
	Raw       []byte
}

// MarketInfos defines all the info for a market for every symbol.
type MarketInfos map[Symbol]MarketInfo
