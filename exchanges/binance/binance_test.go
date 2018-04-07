package binance

import (
	"testing"

	"github.com/barthr/cxtgo"
	"github.com/stretchr/testify/assert"
)

func TestBinance_AmountToLots(t *testing.T) {
	assert := assert.New(t)

	binance := New()

	assert.Equal(0.123, binance.AmountToLots(cxtgo.NewSymbol("ETH", "BTC"), 0.1234))
	assert.Equal(0.100, binance.AmountToLots(cxtgo.NewSymbol("ETH", "BTC"), 0.1001))
}
