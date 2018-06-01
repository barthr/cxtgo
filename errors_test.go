package cxtgo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE(t *testing.T) {
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
