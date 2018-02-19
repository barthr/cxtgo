package cxtgo

import (
	"reflect"
	"testing"
)

func TestNewSymbol(t *testing.T) {
	type args struct {
		from  string
		to    string
		delim []string
	}
	tests := []struct {
		name       string
		args       args
		want       Symbol
		wantString string
	}{
		{
			name: "empty symbol",
			args: args{
				from:  "",
				to:    "",
				delim: nil,
			},
			want: Symbol{
				first:  "",
				second: "",
				delim:  "",
			},
			wantString: "",
		},
		{
			name: "test non empty symbol",
			args: args{
				from:  "BTC",
				to:    "USD",
				delim: nil,
			},
			want: Symbol{
				first:  "BTC",
				second: "USD",
				delim:  "",
			},
			wantString: "BTCUSD",
		},
		{
			name: "test non empty symbol with delim",
			args: args{
				from:  "BTC",
				to:    "USD",
				delim: []string{"/"},
			},
			want: Symbol{
				first:  "BTC",
				second: "USD",
				delim:  "/",
			},
			wantString: "BTC/USD",
		},
		{
			name: "test non empty symbol with multiple delim",
			args: args{
				from:  "BTC",
				to:    "USD",
				delim: []string{"/", "#"},
			},
			want: Symbol{
				first:  "BTC",
				second: "USD",
				delim:  "/",
			},
			wantString: "BTC/USD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSymbol(tt.args.from, tt.args.to, tt.args.delim...); !reflect.DeepEqual(got, tt.want) && got.String() != tt.wantString {
				t.Errorf("NewSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbol_Reverse(t *testing.T) {
	type fields struct {
		delim  string
		first  string
		second string
	}
	tests := []struct {
		name   string
		fields fields
		want   Symbol
	}{
		{
			name: "empty symbol",
			fields: fields{
				first:  "",
				second: "",
				delim:  "",
			},
			want: Symbol{
				first:  "",
				second: "",
				delim:  "",
			},
		},
		{
			name: "non empty symbol",
			fields: fields{
				first:  "BTC",
				second: "USD",
				delim:  "",
			},
			want: Symbol{
				first:  "USD",
				second: "BTC",
				delim:  "",
			},
		},
		{
			name: "non empty symbol with delim",
			fields: fields{
				first:  "BTC",
				second: "USD",
				delim:  "/",
			},
			want: Symbol{
				first:  "USD",
				second: "BTC",
				delim:  "/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Symbol{
				delim:  tt.fields.delim,
				first:  tt.fields.first,
				second: tt.fields.second,
			}
			if got := p.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Symbol.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbolSet(t *testing.T) {
	type args struct {
		data Symbols
	}
	tests := []struct {
		name string
		args args
		want Symbols
	}{
		{
			name: "nil symbols",
			args: args{
				data: nil,
			},
			want: nil,
		},
		{
			name: "empty symbols",
			args: args{
				data: Symbols{},
			},
			want: Symbols{},
		},
		{
			name: "non empty duplicate symbols",
			args: args{
				data: Symbols{
					NewSymbol("BTC", "USD"),
					NewSymbol("BTC", "USD"),
				},
			},
			want: Symbols{
				NewSymbol("BTC", "USD"),
			},
		},
		{
			name: "non empty symbols",
			args: args{
				data: Symbols{
					NewSymbol("BTC", "USD"),
				},
			},
			want: Symbols{
				NewSymbol("BTC", "USD"),
			},
		},
		{
			name: "non empty unique symbols",
			args: args{
				data: Symbols{
					NewSymbol("BTC", "USD"),
					NewSymbol("ETH", "USD"),
				},
			},
			want: Symbols{
				NewSymbol("BTC", "USD"),
				NewSymbol("ETH", "USD"),
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SymbolSet(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SymbolSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
