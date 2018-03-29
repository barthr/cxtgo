package cxtgo

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/ratelimit"

	binance "github.com/adshao/go-binance"
)

const binanceReqPerMin = 1200

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	test   bool
	base   *Base
	client *binance.Client
	once   *sync.Once

	rl ratelimit.Limiter
}

// NewBinance returns an instance of the binance exchange
func NewBinance(opts ...Opt) *Binance {
	binanceOpts := []Opt{
		WithName("Binance"),
		WithUserAgent("cxtgo/0.1"),
	}
	binanceOpts = append(binanceOpts, opts...)

	ex := NewBase(binanceOpts...)
	b := &Binance{
		base:   ex,
		client: binance.NewClient(ex.APIKEY, ex.APISecret),
		once:   &sync.Once{},
		rl:     ratelimit.New(binanceReqPerMin / 60),
	}

	return b
}

// Info returns the base info for the binance exchange
func (b *Binance) Info() Base {
	return *b.base
}

// LoadMarkets loads all the markets from binance
func (b *Binance) LoadMarkets(ctx context.Context) (map[Symbol]MarketInfo, error) {
	b.rl.Take()
	info, err := b.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, NetworkError{ExchangeError{"binance", err}}
	}

	marketInfos := map[Symbol]MarketInfo{}
	for _, symbol := range info.Symbols {
		internalSymbol := NewSymbol(symbol.BaseAsset, symbol.QuoteAsset)

		minPrice, err := strconv.ParseFloat(symbol.Filters[0]["minPrice"], 64)
		if err != nil {
			return nil, ConversionError{ExchangeError{"binance", err}}
		}
		maxPrice, err := strconv.ParseFloat(symbol.Filters[0]["maxPrice"], 64)
		if err != nil {
			return nil, ConversionError{ExchangeError{"binance", err}}
		}
		lotSize, err := strconv.ParseFloat(symbol.Filters[1]["stepSize"], 64)
		if err != nil {
			return nil, ConversionError{ExchangeError{"binance", err}}
		}
		minQty, err := strconv.ParseFloat(symbol.Filters[1]["minQty"], 64)
		if err != nil {
			return nil, ConversionError{ExchangeError{"binance", err}}
		}
		maxQty, err := strconv.ParseFloat(symbol.Filters[1]["maxQty"], 64)
		if err != nil {
			return nil, ConversionError{ExchangeError{"binance", err}}
		}
		minNotional, err := strconv.ParseFloat(symbol.Filters[2]["minNotional"], 64)
		if err != nil {
			return nil, ConversionError{ExchangeError{"binance", err}}
		}
		marketInfos[internalSymbol] = MarketInfo{
			ID:     strings.ToLower(symbol.Symbol),
			Base:   symbol.BaseAsset,
			Quote:  symbol.QuoteAsset,
			Symbol: internalSymbol,
			Maker:  0.001,
			Taker:  0.001,
			Active: true,
			Precision: MarketPrecision{
				Base:   symbol.BaseAssetPrecision,
				Quote:  symbol.QuotePrecision,
				Price:  precisionFromString(symbol.Filters[0]["minPrice"], "."),
				Amount: precisionFromString(symbol.Filters[1]["minQty"], "."),
			},
			Lot: lotSize,
			Limits: MarketLimit{
				Price: MinMax{
					Min: minPrice,
					Max: maxPrice,
				},
				Amount: MinMax{
					Min: minQty,
					Max: maxQty,
				},
				Cost: MinMax{
					Min: minNotional,
				},
			},
			// todo raw
		}
	}
	// copy the map but return a unmodifiable version
	for key, value := range marketInfos {
		b.base.Market[key] = value
	}
	// Call dummy function so that once is triggered
	b.once.Do(func() {})
	return marketInfos, nil
}

func (b *Binance) initMarkets() error {
	var err error
	b.once.Do(func() {
		_, err = b.LoadMarkets(context.Background())
	})
	return err
}

func (b *Binance) FetchMarkets(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	// If not called then call!

	panic("not implemented")
}

func (b *Binance) FetchTicker(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchTickers(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOrderBook(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOHLCV(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchTrades(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchBalance(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) CreateOrder(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) CancelOrder(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) CancelAllOrders(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOrder(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOrders(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOpenOrders(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchClosedOrders(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchMyTrades(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) Deposit(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) Withdraw(ctx context.Context) (Response, error) {
	if err := b.initMarkets(); err != nil {
		return Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) AmountToLots(value float64) float64 {
	if err := b.initMarkets(); err != nil {
		return .0
	}
	panic("not implemented")
}

func precisionFromString(input string, splitter string) int {
	parts := strings.Split(input, splitter)
	pricePrecision := 0
	if len(parts) != 2 {
		return 0
	}
	for _, item := range parts[1] {
		pricePrecision++
		if item != '0' {
			break
		}
	}
	return pricePrecision
}
