package exchange

type MarketPrecision struct {
	Price  int
	Amount int
	Cost   int
}

type MinMax struct {
	Min float64
	Max float64
}

type MarketLimit struct {
	Amount MinMax
	Price  MinMax
}

type MarketInfo struct {
	ID        string
	Base      string
	Quote     string
	Symbol    Symbol
	Active    bool
	Precision MarketPrecision
	Limits    MarketLimit
	Raw       []byte
}
