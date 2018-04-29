package cxtgo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStreamOpt(t *testing.T) {
	r := require.New(t)

	// start with empty config
	config := &StreamConfig{}
	params := Params{"test": "test"}
	symbol := NewSymbol("ETH", "BTC")
	ctx := context.Background()
	// Create some opts
	opts := []StreamOpt{
		WithStreamParams(params),
		WithStreamSymbol(symbol),
		WithStreamContext(ctx),
	}
	// apply opts to config
	for _, opt := range opts {
		opt(config)
	}
	r.Equal(params, config.Params)
	r.Equal(symbol, config.Symbol)
	r.Equal(ctx, config.Context)
}
