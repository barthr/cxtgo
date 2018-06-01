package cxtgo

import (
	"errors"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/require"
)

func TestE(t *testing.T) {
	r := require.New(t)

	err := E("test")
	r.IsType(&Error{}, err)
	{
		err := E(ExchangeName("binance"), Op("public/LimitOrder"), SymbolNotFound, errors.New("BTC/USD"))
		spew.Dump(err.Error())
		// r.Equal(ExchangeName("binance"), err.Exchange)
	}

}
