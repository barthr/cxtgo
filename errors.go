package cxtgo

import "fmt"

type ExchangeError struct {
	Exchange string
	Cause    error
}

func (ee ExchangeError) Error() string {
	return fmt.Sprintf("[%s] failed because %v", ee.Exchange, ee.Cause)
}

// NetworkError defines a network error from the exchange
type NetworkError struct{ ExchangeError }

type ConversionError struct{ ExchangeError }
