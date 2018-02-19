package base

import (
	"time"
)

// ExchangeOpt mutates the settings for the exchange
type ExchangeOpt func(*Exchange)

// WithName sets the name for the exchange
func WithName(name string) ExchangeOpt {
	return func(ex *Exchange) {
		ex.Name = name
	}
}

// WithVersion sets the version for the exchange
func WithVersion(version string) ExchangeOpt {
	return func(ex *Exchange) {
		ex.Version = version
	}
}

// WithUserAgent sets the user agent for the exchange
func WithUserAgent(userAgent string) ExchangeOpt {
	return func(ex *Exchange) {
		ex.UserAgent = userAgent
	}
}

// WithAPIKey sets the api key for the exchange
func WithAPIKey(key string) ExchangeOpt {
	return func(ex *Exchange) {
		ex.APIKEY = key
	}
}

// WithAPISecret sets the api secret for the exchange
func WithAPISecret(secret string) ExchangeOpt {
	return func(ex *Exchange) {
		ex.APISecret = secret
	}
}

// NewExchange returns a new exchange with the given opts applied
func NewExchange(opts ...ExchangeOpt) *Exchange {
	ex := &Exchange{
		ID:   "unknown",
		Name: "unnamed exchange",
	}
	for _, opt := range opts {
		opt(ex)
	}
	return ex
}

// Exchange is the base information and methods for an exchange
type Exchange struct {
	ID         string
	Name       string
	Version    string
	UserAgent  string
	APIKEY     string
	APISecret  string
	Countries  []string
	URLs       map[string]string
	Endpoints  map[string]string
	Has        map[string]bool
	TimeFrames map[string]time.Duration
	RateLimit  time.Duration
	Config     Config
}
