package cxtgo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE(t *testing.T) {
	t.Parallel()

	var tt = map[string]struct {
		in       []interface{}
		expected error
	}{
		"test with invalid arguments": {
			in:       []interface{}{10},
			expected: fmt.Errorf("unknown type %T, value %v in error call", 10, 10),
		},
		"test with invalid arguments and valid": {
			in:       []interface{}{10, "binance"},
			expected: fmt.Errorf("unknown type %T, value %v in error call", 10, 10),
		},
		"test with one valid argument": {
			in:       []interface{}{"binance"},
			expected: &Error{Exchange: ExchangeName("binance")},
		},
		"test with two valid arguments": {
			in:       []interface{}{"binance", "public.Order"},
			expected: &Error{Exchange: ExchangeName("binance"), Op: Op("public.Order")},
		},
		"test with three valid arguments": {
			in:       []interface{}{"binance", "public.Order", SymbolNotFound},
			expected: &Error{Exchange: ExchangeName("binance"), Op: Op("public.Order"), Kind: SymbolNotFound},
		},
		"test with three valid arguments and 1 invalid": {
			in:       []interface{}{"binance", "public.Order", SymbolNotFound, 10},
			expected: fmt.Errorf("unknown type %T, value %v in error call", 10, 10),
		},
		"test with all arguments": {
			in:       []interface{}{"binance", "public.Order", errors.New("test"), SymbolNotFound},
			expected: &Error{Exchange: ExchangeName("binance"), Op: Op("public.Order"), Kind: SymbolNotFound, Err: errors.New("test")},
		},
		"test with typed arguments": {
			in:       []interface{}{ExchangeName("binance"), Op("public.Order"), errors.New("test"), SymbolNotFound},
			expected: &Error{Exchange: ExchangeName("binance"), Op: Op("public.Order"), Kind: SymbolNotFound, Err: errors.New("test")},
		},
		"test with nested error": {
			in:       []interface{}{ExchangeName("binance"), Op("public.Order"), &Error{Exchange: exchanges[0]}, SymbolNotFound},
			expected: &Error{Exchange: ExchangeName("binance"), Op: Op("public.Order"), Kind: SymbolNotFound, Err: &Error{Exchange: exchanges[0]}},
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			output := E(test.in...)
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	var tt = map[string]struct {
		in       *Error
		expected string
	}{
		"test with empty error": {
			in:       &Error{},
			expected: "no error",
		},
		"test with only exchange": {
			in: &Error{
				Exchange: exchanges[0],
			},
			expected: "exchange binance",
		},
		"test with exchange and operation": {
			in: &Error{
				Exchange: exchanges[0],
				Op:       Op("test"),
			},
			expected: "(test): exchange binance",
		},
		"test with exchange, operation, kind": {
			in: &Error{
				Exchange: exchanges[0],
				Op:       Op("test"),
				Kind:     SymbolNotFound,
			},
			expected: "(test): exchange binance: symbol not found",
		},
		"test with exchange, operation, kind, error": {
			in: &Error{
				Exchange: exchanges[0],
				Op:       Op("test"),
				Kind:     SymbolNotFound,
				Err:      errors.New("BTC/USD"),
			},
			expected: "(test): exchange binance: symbol not found: BTC/USD",
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			output := test.in.Error()
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestIs(t *testing.T) {
	t.Parallel()
	var tt = map[string]struct {
		in1      ErrorKind
		in2      error
		expected bool
	}{
		"test with nil error": {
			in1:      SymbolNotFound,
			in2:      nil,
			expected: false,
		},
		"test with invalid type error": {
			in1:      SymbolNotFound,
			in2:      errors.New("test"),
			expected: false,
		},
		"test with valid type error of different kind": {
			in1:      SymbolNotFound,
			in2:      &Error{},
			expected: false,
		},
		"test with valid type error of same kind": {
			in1: SymbolNotFound,
			in2: &Error{
				Kind: SymbolNotFound,
			},
			expected: true,
		},
		"test with nested valid type error of same kind": {
			in1: SymbolNotFound,
			in2: &Error{
				Kind: Other,
				Err: &Error{
					Kind: SymbolNotFound,
				},
			},
			expected: true,
		},
		"test with nested valid type error of other kind": {
			in1: SymbolNotFound,
			in2: &Error{
				Kind: Other,
				Err: &Error{
					Kind: NotSupported,
				},
			},
			expected: false,
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			output := Is(test.in1, test.in2)
			assert.Equal(t, test.expected, output)
		})
	}
}
