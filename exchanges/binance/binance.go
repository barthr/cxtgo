package binance

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/barthr/cxtgo"
	"github.com/go-resty/resty"

	"github.com/sirupsen/logrus"

	"go.uber.org/ratelimit"
)

const (
	binanceReqPerMin  = 1200
	defaultRecvWindow = 5000
	baseURL           = "https://api.binance.com"
)

var (
	errSymbolNotFound = func(s cxtgo.Symbol) error {
		return fmt.Errorf("symbol %v is not found", s.String())
	}
	errMarketsNotSet = errors.New("markets not set")
)

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
	invalidSignature        errorCode = -1022
)

// Binance is the binance implementation for cxtgo interface
type Binance struct {
	base cxtgo.Base

	http   *resty.Client
	logger *logrus.Entry
}

// New returns an instance of the binance exchange, with some defaults set.
func New(opts ...cxtgo.BaseOpt) *Binance {
	binanceOpts := []cxtgo.BaseOpt{
		cxtgo.WithName("Binance"),
		cxtgo.WithUserAgent("cxtgo/0.1"),
		cxtgo.WithRatelimit(ratelimit.New(binanceReqPerMin / 60)),
		cxtgo.WithDebuglogger(os.Stdout),
		cxtgo.WithBaseURL(baseURL),
	}
	// Extend the base options with opts passed in
	// The original opts can be overriden because they are executed as last in the chain
	binanceOpts = append(binanceOpts, opts...)
	ex := cxtgo.NewBase(binanceOpts...)
	recvWindow, ok := ex.CustomParams.GetInt("recvWindow")
	if !ok {
		recvWindow = defaultRecvWindow
	}
	logger := logrus.New()
	if ex.Debug {
		logger.SetLevel(logrus.DebugLevel)
		logger.Out = ex.DebugLog
	} else {
		logger.SetLevel(logrus.PanicLevel)
		logger.Out = ioutil.Discard
	}

	// Create default binance struct with some default fields set
	b := &Binance{
		base: ex,
		http: resty.New().
			SetDebug(ex.Debug).
			SetLogPrefix("cxtgo.binance").
			SetLogger(ex.DebugLog).
			SetTimeout(time.Second * 10).
			SetHostURL(ex.BaseURL).
			SetError(&apiError{}).
			SetHeaders(
				map[string]string{
					"User-Agent":   ex.UserAgent,
					"Content-Type": "application/json",
				},
			).
			SetQueryParams(map[string]string{
				"recvWindow":   strconv.Itoa(recvWindow),
				"X-MBX-APIKEY": ex.APIKEY,
			}),
		logger: logger.WithField("cxtgo", "binance"),
	}

	// Set proxy for the request if one is provided as argument
	if ex.Proxy != "" {
		b.http.SetProxy(ex.Proxy)
	}

	return b
}

// Info returns the base info for the binance exchange
func (b *Binance) Info() cxtgo.Base {
	return b.base
}

type apiError struct {
	Code    errorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

func handleError(err interface{}, op string) error {
	v, ok := err.(*apiError)
	if !ok {
		return cxtgo.Err(cxtgo.ExchangeName("binance"), cxtgo.Op(op), cxtgo.ExchangeNotAvailable, err)
	}
	var cause error
	errKind := cxtgo.Other
	switch v.Code {
	case disconnected, timeOut:
		errKind = cxtgo.Network
	case unauthorized:
		errKind = cxtgo.Authentication
	case tooManyRequests:
		errKind = cxtgo.Ratelimited
	}
	return cxtgo.Err(cxtgo.ExchangeName("binance"), cxtgo.Op(op), cause, errKind)
}
