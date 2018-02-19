package base

// Config defines the basic configuration for the exchange
type Config struct {
	APIKEY, APISecret string
	RateLimit         bool
	Websocket         bool
}
