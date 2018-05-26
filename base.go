package cxtgo

import (
	"io"

	"go.uber.org/ratelimit"
)

// BaseOpt mutates the settings for the exchange.
type BaseOpt func(*Base)

// WithCustom sets custom parameters for the exchange.
// Additional exchange specific parameters can be passed in here and used in the exchange implemenations.
func WithCustom(custom Params) BaseOpt {
	return func(b *Base) {
		b.CustomParams = custom
	}
}

// WithIncludeRaw sets the toggle to include the raw response from the exchange.
func WithIncludeRaw(toggle bool) BaseOpt {
	return func(b *Base) {
		b.Raw = toggle
	}
}

// WithName sets the name for the exchange.
func WithName(name string) BaseOpt {
	return func(b *Base) {
		b.Name = name
	}
}

// WithRatelimit sets a rate limit to use for the api calls to the exchange.
func WithRatelimit(rl ratelimit.Limiter) BaseOpt {
	return func(b *Base) {
		b.Ratelimit = rl
	}
}

// WithID sets the id for the exchange.
func WithID(id string) BaseOpt {
	return func(b *Base) {
		b.ID = id
	}
}

// WithUserAgent sets the user agent for the exchange.
func WithUserAgent(userAgent string) BaseOpt {
	return func(b *Base) {
		b.UserAgent = userAgent
	}
}

// WithAPIKey sets the api key for the exchange.
func WithAPIKey(key string) BaseOpt {
	return func(b *Base) {
		b.APIKEY = key
	}
}

// WithAPISecret sets the api secret for the exchange.
func WithAPISecret(secret string) BaseOpt {
	return func(b *Base) {
		b.APISecret = secret
	}
}

// WithDebug sets the debug flag for the exchange.
func WithDebug(toggle bool) BaseOpt {
	return func(b *Base) {
		b.Debug = toggle
	}
}

// WithDebuglogger sets the debug logger output to w for the exchange.
func WithDebuglogger(w io.Writer) BaseOpt {
	return func(b *Base) {
		b.DebugLog = w
	}
}

// WithProxyURL sets the proxy to use for the http requests
func WithProxyURL(proxy string) BaseOpt {
	return func(b *Base) {
		b.Proxy = proxy
	}
}

// WithBaseURL sets the base url for the exchange
func WithBaseURL(url string) BaseOpt {
	return func(b *Base) {
		b.BaseURL = url
	}
}

// NewBase returns a new base exchange with the given opts applied.
func NewBase(opts ...BaseOpt) Base {
	b := Base{
		ID:     "unknown",
		Name:   "unnamed exchange",
		Market: map[Symbol]MarketInfo{},
	}
	for _, opt := range opts {
		opt(&b)
	}
	return b
}

// Base is the base information and methods for an exchange.
type Base struct {
	ID           string
	Name         string
	Raw          bool
	Debug        bool
	DebugLog     io.Writer
	UserAgent    string
	BaseURL      string
	Proxy        string
	APIKEY       string
	APISecret    string
	Ratelimit    ratelimit.Limiter
	Market       MarketInfos
	CustomParams Params
}
