package cxtgo

// Side defines the side for the order (f.e. Sell or Buy).
type Side int

// String returns the string value of s.
func (s Side) String() string {
	switch s {
	case Sell:
		return "sell"
	case Buy:
		return "buy"
	default:
		return "unknown"
	}
}

const (
	// Unknown side for the zero value of side.
	Unknown Side = iota
	// Sell defines the sell side.
	Sell
	// Buy defines the buy side.
	Buy
)
