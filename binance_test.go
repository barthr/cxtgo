package cxtgo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinance_LoadMarkets(t *testing.T) {
	assert := assert.New(t)

	binance := NewBinance()
	info, err := binance.LoadMarkets(context.Background())

	assert.NoError(err, "err should be empty when loading markets")
	assert.NotNil(info, "info should be filled when loading markets")
	assert.Contains(info, NewSymbol("ETH", "BTC"), "info should contain eth btc")

	assert.Equal(MarketInfo{
		ID:     "ethbtc",
		Base:   "ETH",
		Quote:  "BTC",
		Symbol: NewSymbol("ETH", "BTC"),
		Maker:  0.001,
		Taker:  0.001,
		Active: true,
		Precision: MarketPrecision{
			Base:   8,
			Quote:  8,
			Price:  6,
			Amount: 3,
		},
		Lot: 0.00100000,
		Limits: MarketLimit{
			Price: MinMax{
				Min: 0.00000100,
				Max: 100000.00000000,
			},
			Amount: MinMax{
				Min: 0.00100000,
				Max: 100000.00000000,
			},
			Cost: MinMax{
				Min: 0.00100000,
			},
		},
	}, info[NewSymbol("ETH", "BTC")], "given should be equal expected")
}
