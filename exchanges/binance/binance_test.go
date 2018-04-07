package binance

import (
	"testing"

	"github.com/barthr/cxtgo"
	"github.com/stretchr/testify/assert"
)

func TestBinance_AmountToLots(t *testing.T) {
	assert := assert.New(t)

	binance := New()

	val, err := binance.AmountToLots(cxtgo.NewSymbol("ETH", "BTC"), 0.1234)
	assert.NoError(err)
	assert.Equal(0.123, val)

	val, err = binance.AmountToLots(cxtgo.NewSymbol("ETH", "BTC"), 0.1001)
	assert.NoError(err)
	assert.Equal(0.100, val)

	// test with non existing symbol
	_, err = binance.AmountToLots(cxtgo.NewSymbol("X", "Y"), .0)
	assert.Error(err)
}
