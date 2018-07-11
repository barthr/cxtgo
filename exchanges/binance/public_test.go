package binance

import (
	"context"
	"testing"

	"github.com/barthr/cxtgo"
	"github.com/stretchr/testify/assert"
)

func TestBinance_LoadMarkets(t *testing.T) {
	var tt = map[string]struct {
		params     cxtgo.Params
		wantErr    bool
		wantOutput cxtgo.MarketInfos
	}{
		"Test with erroring exchange": {
			params:     cxtgo.Params{},
			wantErr:    true,
			wantOutput: nil,
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			output, err := New().Markets(context.Background(), test.params)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, output, test.wantOutput)
		})
	}
}
