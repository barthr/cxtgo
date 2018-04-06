package binance

import (
	"context"
	"strconv"
	"strings"

	"github.com/barthr/cxtgo"
	"github.com/pkg/errors"
)

// LoadMarkets loads all the markets from binance
func (b *Binance) LoadMarkets(ctx context.Context) (map[cxtgo.Symbol]cxtgo.MarketInfo, error) {
	b.base.Ratelimit.Take()
	info, err := b.client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, cxtgo.NetworkError{cxtgo.ExchangeError{"binance", err}}
	}

	marketInfos := map[cxtgo.Symbol]cxtgo.MarketInfo{}
	for _, symbol := range info.Symbols {
		internalSymbol := cxtgo.NewSymbol(symbol.BaseAsset, symbol.QuoteAsset)

		rawFilters := []string{
			symbol.Filters[0]["minPrice"],
			symbol.Filters[0]["maxPrice"],
			symbol.Filters[1]["stepSize"],
			symbol.Filters[1]["minQty"],
			symbol.Filters[1]["maxQty"],
			symbol.Filters[2]["minNotional"],
		}
		conversions := [6]float64{}
		for i, rf := range rawFilters {
			conversion, err := strconv.ParseFloat(rf, 64)
			if err != nil {
				return nil, cxtgo.ConversionError{
					ExchangeError: cxtgo.ExchangeError{
						Exchange: "binance",
						Cause:    errors.WithStack(err),
					},
				}
			}
			conversions[i] = conversion
		}
		marketInfos[internalSymbol] = cxtgo.MarketInfo{
			ID:     strings.ToLower(symbol.Symbol),
			Base:   symbol.BaseAsset,
			Quote:  symbol.QuoteAsset,
			Symbol: internalSymbol,
			Maker:  0.001,
			Taker:  0.001,
			Active: true,
			Precision: cxtgo.MarketPrecision{
				Base:   symbol.BaseAssetPrecision,
				Quote:  symbol.QuotePrecision,
				Price:  precisionFromString(symbol.Filters[0]["minPrice"], "."),
				Amount: precisionFromString(symbol.Filters[1]["minQty"], "."),
			},
			Lot: conversions[2],
			Limits: cxtgo.MarketLimit{
				Price: cxtgo.MinMax{
					Min: conversions[0],
					Max: conversions[1],
				},
				Amount: cxtgo.MinMax{
					Min: conversions[3],
					Max: conversions[4],
				},
				Cost: cxtgo.MinMax{
					Min: conversions[5],
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

func (b *Binance) FetchMarkets(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	// If not called then call!

	panic("not implemented")
}

func (b *Binance) FetchTicker(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchTickers(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOrderBook(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchOHLCV(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}

func (b *Binance) FetchTrades(ctx context.Context) (cxtgo.Response, error) {
	if err := b.initMarkets(); err != nil {
		return cxtgo.Response{}, err
	}
	panic("not implemented")
}
