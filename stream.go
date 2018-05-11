package cxtgo

import (
	"context"
)

// StreamConfig defines the configuration options for the stream
type StreamConfig struct {
	context.Context
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

// WithStreamContext set's the context for the stream
func WithStreamContext(ctx context.Context) StreamOpt {
	return func(sc *StreamConfig) {
		sc.Context = ctx
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

// StreamType defines which type of stream it is. This is helpfull for debuggin errors
type StreamType int

const (
	// UnknownStream indicates a stream which isn't known
	UnknownStream StreamType = iota
	// TradeStream indicates a trade stream
	TradeStream
	// TickerStream indicates a ticker stream
	TickerStream
	// OrderbookStream indicates a orderbook stream type
	OrderbookStream
)
