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

type (
	// NetworkError defines a network error from the exchange.
	NetworkError struct{ ExchangeError }

	// ConversionError defines a conversion error from the exchange.
	ConversionError struct{ ExchangeError }

	// NotSupported defines a not supported error from the exchange.
	NotSupported struct{ ExchangeError }

	// AuthenticationError defines a authentication error from the exchange.
	AuthenticationError struct{ ExchangeError }

	// InsufficientFunds defines a error indicating that there are not sufficient funds for the operation.
	InsufficientFunds struct{ ExchangeError }

	InvalidOrder struct{ ExchangeError }

	OrderNotFound struct{ ExchangeError }

	ExchangeNotAvailable struct{ ExchangeError }
)
