package cxtgo

import (
	"fmt"
	"strings"
)

// Error is the type that implements the error interface.
// It contains a number of fields, each of different type.
// An Error value may leave some values unset.
// Adapted from https://github.com/upspin/upspin/blob/master/errors/errors.go
type Error struct {
	// Exchange is the name of the exchange being used.
	Exchange ExchangeName
	// Op is the operation being performed, usually the name of the method
	Op Op
	// Kind is the class of error, as defined by the error kinds.
	// Other is used if its class is unknown or irrelevant.
	Kind ErrorKind
	// The underlying error that triggered this one, if any.
	Err error
}

func (e *Error) Error() string {
	b := new(strings.Builder)
	if e.Op != "" {
		pad(b, ": ")
		b.WriteString("(")
		b.WriteString(string(e.Op))
		b.WriteString(")")
	}
	if e.Exchange != "" {
		pad(b, ": ")
		b.WriteString("exchange ")
		b.WriteString(string(e.Exchange))
	}
	if e.Kind != 0 {
		pad(b, ": ")
		b.WriteString(e.Kind.String())
	}
	if e.Err != nil {
		pad(b, ": ")
		b.WriteString(e.Err.Error())
	}
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
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

// E is the standard error function for creating errors in cxtgo.
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to cxtgo.E(...) with no arguments")
	}
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case ExchangeName:
			e.Exchange = arg
		case Op:
			e.Op = arg
		case ErrorKind:
			e.Kind = arg
		case string:
			var found bool
			// check available exchanges and pick the correct one
			for _, exchange := range exchanges {
				if ExchangeName(arg) == exchange {
					e.Exchange = exchange
					found = true
				}
			}
			if !found {
				e.Op = Op(arg)
			}
		case *Error:
			// Make a copy
			copy := *arg
			e.Err = &copy
		case error:
			e.Err = arg
		default:
			return fmt.Errorf("unknown type %T, value %v in error call", arg, arg)
		}
	}
	return e
}

// pad appends str to the buffer if the buffer already has some data.
func pad(b *strings.Builder, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}

// Is reports whether err is an *Error of the given Kind.
// If err is nil then Is returns false.
func Is(kind ErrorKind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.Kind != Other {
		return e.Kind == kind
	}
	if e.Err != nil {
		return Is(kind, e.Err)
	}
	return false
}
