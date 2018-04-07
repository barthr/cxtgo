package binance

import (
	"os"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/resync"
	"github.com/myesui/uuid"
	"github.com/pkg/errors"
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
		cxtgo.WithID(uuid.NewV4().String()),
		cxtgo.WithName("Binance"),
		cxtgo.WithUserAgent("cxtgo/0.1"),
		cxtgo.WithRatelimit(ratelimit.New(binanceReqPerMin / 60)),
		cxtgo.WithDebuglogger(os.Stdout),
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

// Reset resets the resync.Once, this allows the exchange to reload the related symbol info.
func (b *Binance) Reset() {
	b.once.Reset()
}

// AmountToLots converts the given amount to the lot sized amount.
// Returns the zero value of float64 when the symbol is not found in the marketinfo.
func (b *Binance) AmountToLots(s cxtgo.Symbol, amount float64) (float64, error) {
	if err := b.initMarkets(); err != nil {
		return 0, errors.WithStack(err)
	}
	info, ok := b.base.Market[s]
	if !ok {
		return 0, errors.New("symbol not found")
	}
	return binance.AmountToLotSize(info.Lot, info.Precision.Amount, amount), nil
}

type binanceAdapter struct {
}
