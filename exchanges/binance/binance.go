package binance

import (
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo"
	"github.com/barthr/cxtgo/resync"
	"github.com/myesui/uuid"
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"
)

const (
	binanceReqPerMin  = 1200
	defaultRecvWindow = 5000
	baseURL           = "https://api.binance.com"
)

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	base   cxtgo.Base
	client *binance.Client
	http   *resty.Client
	once   resync.Once
}

// New returns an instance of the binance exchange, with some defaults set.
func New(opts ...cxtgo.BaseOpt) *Binance {
	binanceOpts := []cxtgo.BaseOpt{
		cxtgo.WithID(uuid.NewV4().String()),
		cxtgo.WithName("Binance"),
		cxtgo.WithUserAgent("cxtgo/0.1"),
		cxtgo.WithRatelimit(ratelimit.New(binanceReqPerMin / 60)),
		cxtgo.WithDebuglogger(os.Stdout),
	}
	binanceOpts = append(binanceOpts, opts...)

	ex := cxtgo.NewBase(binanceOpts...)
	recvWindow, ok := ex.CustomParams.GetInt("recvWindow")
	if !ok {
		recvWindow = defaultRecvWindow
	}
	b := &Binance{
		base:   ex,
		client: binance.NewClient(ex.APIKEY, ex.APISecret),
		once:   resync.Once{},
		http: resty.New().
			SetDebug(ex.Debug).
			SetLogPrefix("cxtgo.binance").
			SetLogger(ex.DebugLog).
			SetTimeout(time.Second*10).
			SetHostURL(baseURL).
			SetError(&apiError{}).
			SetHeader("User-Agent", ex.UserAgent).
			SetHeader("Content-Type", "application/json").
			SetQueryParams(map[string]string{
				"recvWindow":   strconv.Itoa(recvWindow),
				"X-MBX-APIKEY": ex.APIKEY,
			}),
	}

	if ex.Proxy != "" {
		b.http.SetProxy(ex.Proxy)
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
	return cxtgo.AmountToLotSize(info.Lot, info.Precision.Amount, amount), nil
}

type apiError struct {
	Code    errorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

// see: https://github.com/binance-exchange/binance-official-api-docs/blob/master/errors.md
type errorCode int

const (
	unknown                 errorCode = -1000
	disconnected            errorCode = -1001
	unauthorized            errorCode = -1002
	tooManyRequests         errorCode = -1003
	unexpectedResponse      errorCode = -1006
	timeOut                 errorCode = -1007
	invalidMessage          errorCode = -1013
	unknownOrderCompisition errorCode = -1014
	tooManyOrders           errorCode = -1015
	serviceShuttingDown     errorCode = -1016
	unsupportedOperation    errorCode = -1020
	invalidTimestamp        errorCode = -1021
	invalidSIgnature        errorCode = -1022
)
