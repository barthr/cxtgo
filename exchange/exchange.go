package exchange

import (
	"time"
)

// Opt mutates the settings for the exchange
type Opt func(*Base)

// WithName sets the name for the exchange
func WithName(name string) Opt {
	return func(b *Base) {
		b.Name = name
	}
}

// WithURLs sets the urls for the exchange
// Valid urls are things like the documentation, website etc.
func WithURLs(urls map[string]string) Opt {
	return func(b *Base) {
		b.URLs = urls
	}
}

// WithID sets the id for the exchange
func WithID(id string) Opt {
	return func(b *Base) {
		b.ID = id
	}
}

// WithVersion sets the version for the exchange
func WithVersion(version string) Opt {
	return func(b *Base) {
		b.Version = version
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

// WithCountries sets the api countries for the exchange
func WithCountries(countries ...string) Opt {
	return func(b *Base) {
		b.Countries = countries
	}
}

// NewBase returns a new base exchange with the given opts applied
func NewBase(opts ...Opt) *Base {
	b := &Base{
		ID:   "unknown",
		Name: "unnamed exchange",
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// Base is the base information and methods for an exchange
type Base struct {
	ID         string
	Name       string
	Version    string
	UserAgent  string
	APIKEY     string
	APISecret  string
	Countries  []string
	URLs       map[string]string
	Has        map[string]bool
	TimeFrames map[string]time.Duration
	Market     map[string]MarketInfo
	RateLimit  time.Duration
}
