package binance

import (
	"context"
	"testing"

	"github.com/barthr/cxtgo"
	"github.com/stretchr/testify/assert"
)

func TestBinance_LoadMarkets(t *testing.T) {
	assert := assert.New(t)

	binance := New()
	info, err := binance.LoadMarkets(context.Background())

	assert.NoError(err, "err should be empty when loading markets")
	assert.NotNil(info, "info should be filled when loading markets")
	assert.Contains(info, cxtgo.NewSymbol("ETH", "BTC"), "info should contain eth btc")
}

func TestBinance_AmountToLots(t *testing.T) {
	assert := assert.New(t)

	binance := New()

	lotSizedAmount := binance.AmountToLots(cxtgo.NewSymbol("ETH", "BTC"), 0.1234)
	assert.Equal(0.123, lotSizedAmount)
}
