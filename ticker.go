package cxtgo

import "time"

// Ticker represents a ticker from an exchange.
type Ticker struct {
	Symbol
	Ask         float64
	Bid         float64
	Basevolume  float64
	Quotevolume float64
	Average     float64
	Change      float64
	Open        float64
	Close       float64
	First       float64
	Last        float64
	High        float64
	Low         float64
	Percentage  float64
	Vwap        float64

	Datetime  time.Time
	Timestamp int64

	Raw []byte
}

// Tickers represents multiple tickers from an exchange.
type Tickers map[Symbol]Tickers

// TickerStreamer is a streamer interface for tickers
type TickerStreamer interface {
	StreamTicker(s Symbol, onUpdate func(t Ticker), onError func(err error)) error
}
