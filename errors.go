package cxtgo

import (
	"fmt"
	"time"
)

// BaseError is the base error class for errors from cxtgo
type BaseError struct {
	Exchange string
	Cause    error
}

func (ee BaseError) Error() string {
	return fmt.Sprintf("[%s] failed because %v", ee.Exchange, ee.Cause)
}

// WrapError wraps an exchange error
func WrapError(parent error, exchange string, cause error) error {
	switch parent.(type) {
	case SymbolNotFoundError:
		return SymbolNotFoundError{BaseError{exchange, cause}}
	case NetworkError:
		return NetworkError{BaseError{exchange, cause}}
	case ConversionError:
		return ConversionError{BaseError{exchange, cause}}
	case NotSupportedError:
		return NotSupportedError{BaseError{exchange, cause}}
	case AuthenticationError:
		return AuthenticationError{BaseError{exchange, cause}}
	case InsufficientFundsError:
		return InsufficientFundsError{BaseError{exchange, cause}}
	case InvalidOrderError:
		return InvalidOrderError{BaseError{exchange, cause}}
	case OrderNotFoundError:
		return OrderNotFoundError{BaseError{exchange, cause}}
	case ExchangeNotAvailableError:
		return ExchangeNotAvailableError{BaseError{exchange, cause}}
	}
	return BaseError{exchange, cause}
}

type (
	// SymbolNotFoundError defines an error for when executing an action on the exchange for a symbol which is not found.
	SymbolNotFoundError struct{ BaseError }
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

type (
	// StreamClosedByExchangeError represents an error when the stream is closed by the exchange.
	StreamClosedByExchangeError struct{ StreamError }
	// StreamUnavailableError represents an error when the stream is (currently) unavailable.
	StreamUnavailableError struct{ StreamError }
)

// StreamError is the base error in a stream
type StreamError struct {
	StreamType
	BaseError
}

// StreamMaintenanceError represents an error when the stream is under maintenance.
type StreamMaintenanceError struct {
	StreamError
	time.Duration
}
