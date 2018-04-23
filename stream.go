package cxtgo

// StreamConfig defines the configuration options for the stream
type StreamConfig struct {
	Symbol
	Params
}

// StreamOpt defines a function option to modify the streamconfiguration
type StreamOpt func(*StreamConfig)

// WithStreamParams set's the stream parameters to the given params.
func WithStreamParams(params Params) StreamOpt {
	return func(sc *StreamConfig) {
		sc.Params = params
	}
}

// WithStreamSymbol set's the stream symbol the the given symbol.
func WithStreamSymbol(s Symbol) StreamOpt {
	return func(sc *StreamConfig) {
		sc.Symbol = s
	}
}

// TickerStreamer is a streamer interface for tickers
type TickerStreamer interface {
	StreamTicker(onUpdate func(t Ticker), onError func(err error), opts ...StreamOpt) error
}

// TradeStreamer is a streamer interface for the trades
type TradeStreamer interface {
	StreamTrades(onUpdate func(t Trade), onError func(err error), opts ...StreamOpt) error
}

// OrderbookStreamer is a streamer interface for the orderbook
type OrderbookStreamer interface {
	StreamOrderbook(onUpdate func(s Summary), onError func(err error), opts ...StreamOpt) error
}
