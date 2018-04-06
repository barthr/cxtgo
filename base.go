package cxtgo

import "go.uber.org/ratelimit"

// Opt mutates the settings for the exchange
type Opt func(*Base)

// WithName sets the name for the exchange
func WithName(name string) Opt {
	return func(b *Base) {
		b.Name = name
	}
}

// WithRatelimit sets a rate limit to use for the api calls to the exchange
func WithRatelimit(rl ratelimit.Limiter) Opt {
	return func(b *Base) {
		b.Ratelimit = rl
	}
}

// WithID sets the id for the exchange
func WithID(id string) Opt {
	return func(b *Base) {
		b.ID = id
	}
}

// WithUserAgent sets the user agent for the exchange
func WithUserAgent(userAgent string) Opt {
	return func(b *Base) {
		b.UserAgent = userAgent
	}
}

// WithAPIKey sets the api key for the exchange
func WithAPIKey(key string) Opt {
	return func(b *Base) {
		b.APIKEY = key
	}
}

// WithAPISecret sets the api secret for the exchange
func WithAPISecret(secret string) Opt {
	return func(b *Base) {
		b.APISecret = secret
	}
}

// NewBase returns a new base exchange with the given opts applied
func NewBase(opts ...Opt) Base {
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

// Base is the base information and methods for an exchange
type Base struct {
	ID        string
	Name      string
	UserAgent string
	APIKEY    string
	APISecret string
	Ratelimit ratelimit.Limiter
	Market    map[Symbol]MarketInfo
}
