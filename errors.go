package cxtgo

import "fmt"

// BaseError is the base error class for errors from cxtgo
type BaseError struct {
	Exchange string
	Cause    error
}

func (ee BaseError) Error() string {
	return fmt.Sprintf("[%s] failed because %v", ee.Exchange, ee.Cause)
}

// NewError instantiates a new exchange error
func NewError(exchange string, cause error) BaseError {
	return BaseError{
		Exchange: exchange,
		Cause:    cause,
	}
}

type (
	// NetworkError defines a network error from the exchange.
	NetworkError struct{ BaseError }

	// ConversionError defines a conversion error from the exchange.
	ConversionError struct{ BaseError }

	// NotSupportedError defines a not supported error from the exchange.
	NotSupportedError struct{ BaseError }

	// AuthenticationError defines a authentication error from the exchange.
	AuthenticationError struct{ BaseError }

	// InsufficientFundsError defines a error indicating that there are not sufficient funds for the operation.
	InsufficientFundsError struct{ BaseError }

	// InvalidOrderError defines an error indicating that creating an order failed because it's invalid.
	InvalidOrderError struct{ BaseError }

	// OrderNotFoundError defines an error when the requested order is not found.
	OrderNotFoundError struct{ BaseError }
	// ExchangeNotAvailableError defines an error for when the exchange is not available.
	ExchangeNotAvailableError struct{ BaseError }
)
