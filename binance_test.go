package cxtgo

// func TestNewBinance(t *testing.T) {
// 	type args struct {
// 		config *exchange.Config
// 		opts   []
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *Binance
// 	}{
// 		{
// 			name: "test with empty arguments",
// 			args: args{
// 				config: nil,
// 				opts:   nil,
// 			},
// 			want: nil,
// 		},
// 		{
// 			name: "test config arguments",
// 			args: args{
// 				config: &exchange.Config{
// 					APISecret: "",
// 					RateLimit: false,
// 					Websocket: true,
// 				},
// 				opts: nil,
// 			},
// 			want: &Binance{
// 				Exchange: &exchange.Exchange{
// 					Name: "binance",
// 					Config: exchange.Config{
// 						APISecret: "",
// 						RateLimit: false,
// 						Websocket: true,
// 					},
// 				},
// 				client: binance.NewClient("", ""),
// 			},
// 		},
// 		{
// 			name: "test with opts arguments",
// 			args: args{
// 				config: &exchange.Config{
// 					APISecret: "",
// 					RateLimit: false,
// 					Websocket: true,
// 				},
// 				opts: []BinanceOptFunc{
// 					BinanceWithHTTPClient(http.DefaultClient),
// 				},
// 			},
// 			want: &Binance{
// 				Exchange: &exchange.Exchange{
// 					Name: "binance",
// 					Config: exchange.Config{
// 						APISecret: "",
// 						RateLimit: false,
// 						Websocket: true,
// 					},
// 				},
// 				client: binance.NewClient("", ""),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewBinance(tt.args.config, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewBinance() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
