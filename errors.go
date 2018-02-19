package cxtgo

import "fmt"

type ExchangeError struct {
	exchange string
	cause    error
}

func (ee ExchangeError) Error() string {
	return fmt.Sprintf("[%s] failed because %v", ee.exchange, ee.cause)
}

// NetworkError defines a network error from the exchange
type NetworkError struct{ ExchangeError }
