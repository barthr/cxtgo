package binance

import (
	"strings"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/resync"
	"go.uber.org/ratelimit"
)

const binanceReqPerMin = 1200

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	test   bool
	base   cxtgo.Base
	client *binance.Client
	once   resync.Once
}

// New returns an instance of the binance exchange, with some defaults set.
func New(opts ...cxtgo.Opt) *Binance {
	binanceOpts := []cxtgo.Opt{
		cxtgo.WithName("Binance"),
		cxtgo.WithUserAgent("cxtgo/0.1"),
		cxtgo.WithRatelimit(ratelimit.New(binanceReqPerMin / 60)),
	}
	binanceOpts = append(binanceOpts, opts...)

	ex := cxtgo.NewBase(binanceOpts...)
	b := &Binance{
		base:   ex,
		client: binance.NewClient(ex.APIKEY, ex.APISecret),
		once:   resync.Once{},
	}

	return b
}

// Info returns the base info for the binance exchange
func (b *Binance) Info() cxtgo.Base {
	return b.base
}

func (b *Binance) Reset() {
	b.once.Reset()
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
