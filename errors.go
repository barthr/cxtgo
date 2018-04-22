package cxtgo

import "fmt"

type ExchangeError struct {
	Exchange string
	Cause    error
}

func (ee ExchangeError) Error() string {
	return fmt.Sprintf("[%s] failed because %v", ee.Exchange, ee.Cause)
}

// NewError instantiates a new exchange error
func NewError(exchange string, cause error) ExchangeError {
	return ExchangeError{
		Exchange: exchange,
		Cause:    cause,
	}
}

// NetworkError defines a network error from the exchange
type NetworkError struct{ ExchangeError }

type ConversionError struct{ ExchangeError }

type NotSupported struct{ ExchangeError }

type AuthenticationError struct{ ExchangeError }

type InsufficientFunds struct{ ExchangeError }

type InvalidOrder struct{ ExchangeError }

type OrderNotFound struct{ ExchangeError }

type ExchangeNotAvailable struct{ ExchangeError }
