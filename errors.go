package cxtgo

// Error is the type that implements the error interface.
// It contains a number of fields, each of different type.
// An Error value may leave some values unset.
// Adapted from https://github.com/upspin/upspin/blob/master/errors/errors.go
type Error struct {
	// Exchange is the name of the exchange being used.
	Exchange string
	// Op is the operation being performed, usually the name of the method
	Op Op
	// Kind is the class of error, as defined by the error kinds.
	// Other is used if its class is unknown or irrelevant.
	Kind ErrorKind
	// The underlying error that triggered this one, if any.
	Err error
}

// IsZero returns if the error is a zero error.
func (e *Error) IsZero() bool {
	return e.Exchange == "" && e.Kind == 0 && e.Err == nil
}

// Op describes an operation, usually as the package and method,
// such as "exchanges/binance.Ticker".
type Op string

// ErrorKind defines the kind of error this is.
type ErrorKind uint8

// Kinds of errors.
//
// These are the different error kinds from cxtgo.
// Do not reorder this list or remove any
// items since that will change their values.
// New items must be added only to the end.
const (
	Other                  ErrorKind = iota // Unclassified error. This value is not printed in the error message.
	SymbolNotFound                          // Kind for when executing an action on the exchange for a symbol which is not found.
	Network                                 // Permission denied.
	Conversion                              // External I/O error such as network failure.
	NotSupported                            // Item already exists.
	Authentication                          // Item does not exist.
	InsufficientFunds                       // Item is a directory.
	InvalidOrder                            // Item is not a directory.
	OrderNotFound                           // Directory not empty.
	ExchangeNotAvailable                    // Information withheld.
	StreamClosedByExchange                  // Internal error or inconsistency.
	StreamUnavailable                       // No wrapped key for user with read access.
)

func (ek ErrorKind) String() string {
	kind, ok := map[ErrorKind]string{
		Other:                  "other error",
		SymbolNotFound:         "symbol not found",
		Network:                "network error",
		Conversion:             "conversion error",
		NotSupported:           "operation not supported",
		Authentication:         "authentication failure",
		InsufficientFunds:      "insufficient funds",
		InvalidOrder:           "order is invalid",
		OrderNotFound:          "order was not found",
		ExchangeNotAvailable:   "error exchange is not available",
		StreamClosedByExchange: "stream is closed by exchange",
		StreamUnavailable:      "stream is currently unavailable",
	}[ek]
	if !ok {
		kind = "unknown error kind"
	}
	return kind
}

// // BaseError is the base error class for errors from cxtgo
// type BaseError struct {
// 	Exchange string
// 	Cause    error
// }

// func (ee BaseError) Error() string {
// 	return fmt.Sprintf("[%s] failed because %v", ee.Exchange, ee.Cause)
// }

// // WrapError wraps an exchange error
// func WrapError(parent error, exchange string, cause error) error {
// 	switch parent.(type) {
// 	case SymbolNotFoundError:
// 		return SymbolNotFoundError{
// 			BaseError: BaseError{exchange, cause},
// 		}
// 	case NetworkError:
// 		return NetworkError{BaseError{exchange, cause}}
// 	case ConversionError:
// 		return ConversionError{BaseError{exchange, cause}}
// 	case NotSupportedError:
// 		return NotSupportedError{BaseError{exchange, cause}}
// 	case AuthenticationError:
// 		return AuthenticationError{BaseError{exchange, cause}}
// 	case InsufficientFundsError:
// 		return InsufficientFundsError{BaseError{exchange, cause}}
// 	case InvalidOrderError:
// 		return InvalidOrderError{BaseError{exchange, cause}}
// 	case OrderNotFoundError:
// 		return OrderNotFoundError{BaseError{exchange, cause}}
// 	case ExchangeNotAvailableError:
// 		return ExchangeNotAvailableError{BaseError{exchange, cause}}
// 	}
// 	return BaseError{exchange, cause}
// }

// type (
// 	// SymbolNotFoundError defines an error for when executing an action on the exchange for a symbol which is not found.
// 	SymbolNotFoundError struct {
// 		BaseError
// 		Symbol
// 	}

// 	// NetworkError defines a network error from the exchange.
// 	NetworkError struct{ BaseError }

// 	// ConversionError defines a conversion error from the exchange.
// 	ConversionError struct{ BaseError }

// 	// NotSupportedError defines a not supported error from the exchange.
// 	NotSupportedError struct{ BaseError }

// 	// AuthenticationError defines a authentication error from the exchange.
// 	AuthenticationError struct{ BaseError }

// 	// InsufficientFundsError defines a error indicating that there are not sufficient funds for the operation.
// 	InsufficientFundsError struct{ BaseError }

// 	// InvalidOrderError defines an error indicating that creating an order failed because it's invalid.
// 	InvalidOrderError struct{ BaseError }

// 	// OrderNotFoundError defines an error when the requested order is not found.
// 	OrderNotFoundError struct {
// 		BaseError
// 	}

// 	// ExchangeNotAvailableError defines an error for when the exchange is not available.
// 	ExchangeNotAvailableError struct{ BaseError }
// )

// type (
// 	// StreamClosedByExchangeError represents an error when the stream is closed by the exchange.
// 	StreamClosedByExchangeError struct{ StreamError }
// 	// StreamUnavailableError represents an error when the stream is (currently) unavailable.
// 	StreamUnavailableError struct{ StreamError }
// )

// // StreamError is the base error in a stream
// type StreamError struct {
// 	StreamType
// 	BaseError
// }

// // StreamMaintenanceError represents an error when the stream is under maintenance.
// type StreamMaintenanceError struct {
// 	StreamError
// 	time.Duration
// }
