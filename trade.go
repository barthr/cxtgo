package cxtgo

// Trade represents an trade from the exchange.
type Trade struct {
}

// TradeStreamer is a streamer interface for the trades
type TradeStreamer interface {
	StreamTrades(s Symbol, onUpdate func(t Trade), onError func(err error)) error
}
