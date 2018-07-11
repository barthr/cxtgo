package cxtgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountToLot(t *testing.T) {
	var tt = map[string]struct {
		info     MarketInfo
		amount   float64
		expected float64
	}{
		"Test with lot of zero and invalid amount": {
			info: MarketInfo{
				Lot: 0.00100000,
				Precision: MarketPrecision{
					Amount: 8,
				},
			},
			amount:   0.00010000,
			expected: 0,
		},
		"Test with lot": {
			info: MarketInfo{
				Lot: 0.00100000,
				Precision: MarketPrecision{
					Amount: 3,
				},
			},
			amount:   1.39,
			expected: 1.389,
		},
		"Test with big decimal": {
			info: MarketInfo{
				Lot: 0.00100000,
				Precision: MarketPrecision{
					Amount: 8,
				},
			},
			amount:   11.31232419283240912834434,
			expected: 11.312,
		},
		"Test with big number": {
			info: MarketInfo{
				Lot: 0.0010000,
				Precision: MarketPrecision{
					Amount: 8,
				},
			},
			amount:   14000.14000,
			expected: 14000.140,
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			output := AmountToLots(test.info, test.amount)
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestFromString(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		input    string
		splitter string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty arguments",
			args: args{},
			want: 0,
		},
		{
			name: "test with input and no splitter",
			args: args{
				input: "000",
			},
			want: 0,
		},
		{
			name: "test with input and comma splitter",
			args: args{
				input:    "0,0",
				splitter: ",",
			},
			want: 1,
		},
		{
			name: "test with input and dot splitter and high precision",
			args: args{
				input:    "0000000000000000000000000000000.00000000000",
				splitter: ".",
			},
			want: 11,
		},
		{
			name: "test with usual input",
			args: args{
				input:    "0.000055",
				splitter: ".",
			},
			want: 5,
		},
		{
			name: "test with input and multiple dot splitters",
			args: args{
				input:    "0000000000000000000000000000000.....00000000000",
				splitter: ".",
			},
			want: 0,
		},
		{
			name: "test with input and multiple dot splitters",
			args: args{
				input:    "000.0.0.0",
				splitter: ".",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.want, Zeros(tt.args.input, tt.args.splitter), tt.name)
		})
	}
}
