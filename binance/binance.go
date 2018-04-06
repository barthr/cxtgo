package binance

import (
	"strings"
	"sync"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo"
	"go.uber.org/ratelimit"
)

const binanceReqPerMin = 1200

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	test   bool
	base   *cxtgo.Base
	client *binance.Client
	once   *sync.Once

	rl ratelimit.Limiter
}

// NewBinance returns an instance of the binance exchange
func NewBinance(opts ...cxtgo.Opt) *Binance {
	binanceOpts := []cxtgo.Opt{
		cxtgo.WithName("Binance"),
		cxtgo.WithUserAgent("cxtgo/0.1"),
	}
	binanceOpts = append(binanceOpts, opts...)

	ex := cxtgo.NewBase(binanceOpts...)
	b := &Binance{
		base:   ex,
		client: binance.NewClient(ex.APIKEY, ex.APISecret),
		once:   &sync.Once{},
		rl:     ratelimit.New(binanceReqPerMin / 60),
	}

	return b
}

// Info returns the base info for the binance exchange
func (b *Binance) Info() cxtgo.Base {
	return *b.base
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
