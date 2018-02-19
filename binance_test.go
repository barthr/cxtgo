package cxtgo

import (
	"net/http"
	"reflect"
	"testing"

	binance "github.com/adshao/go-binance"
	"github.com/barthr/cxtgo/base"
)

func TestNewBinance(t *testing.T) {
	type args struct {
		config *base.Config
		opts   []BinanceOptFunc
	}
	tests := []struct {
		name string
		args args
		want *Binance
	}{
		{
			name: "test with empty arguments",
			args: args{
				config: nil,
				opts:   nil,
			},
			want: nil,
		},
		{
			name: "test config arguments",
			args: args{
				config: &base.Config{
					APISecret: "",
					RateLimit: false,
					Websocket: true,
				},
				opts: nil,
			},
			want: &Binance{
				Exchange: &base.Exchange{
					Name: "binance",
					Config: base.Config{
						APISecret: "",
						RateLimit: false,
						Websocket: true,
					},
				},
				client: binance.NewClient("", ""),
			},
		},
		{
			name: "test with opts arguments",
			args: args{
				config: &base.Config{
					APISecret: "",
					RateLimit: false,
					Websocket: true,
				},
				opts: []BinanceOptFunc{
					BinanceWithHTTPClient(http.DefaultClient),
				},
			},
			want: &Binance{
				Exchange: &base.Exchange{
					Name: "binance",
					Config: base.Config{
						APISecret: "",
						RateLimit: false,
						Websocket: true,
					},
				},
				client: binance.NewClient("", ""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinance(tt.args.config, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinance() = %v, want %v", got, tt.want)
			}
		})
	}
}
