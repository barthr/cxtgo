package exchange

type MarketPrecision struct {
	Base   int
	Quote  int
	Amount int
	Price  int
}

type MinMax struct {
	Min float64
	Max float64
}

type MarketLimit struct {
	Amount MinMax
	Price  MinMax
	Cost   MinMax
}

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
