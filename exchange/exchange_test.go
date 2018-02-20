package exchange

import (
	"reflect"
	"testing"
)

func TestWithOpts(t *testing.T) {
	type args struct {
		options []Opt
	}
	tests := []struct {
		name string
		args args
		want *Base
	}{
		{
			name: "empty options",
			args: args{
				options: nil,
			},
			want: &Base{
				ID:     "unknown",
				Name:   "unnamed exchange",
				Market: map[Symbol]MarketInfo{},
			},
		},
		{
			name: "with 1 options",
			args: args{
				options: []Opt{
					WithAPIKey("test"),
				},
			},
			want: &Base{
				ID:     "unknown",
				Name:   "unnamed exchange",
				APIKEY: "test",
				Market: map[Symbol]MarketInfo{},
			},
		},
		{
			name: "with multiple options",
			args: args{
				options: []Opt{
					WithAPIKey("test"),
					WithAPISecret("test"),
					WithCountries("test"),
					WithID("test"),
					WithURLs(map[string]string{"doc": "test"}),
					WithUserAgent("cxtgo"),
					WithName("test"),
				},
			},
			want: &Base{
				ID:        "test",
				Name:      "test",
				APIKEY:    "test",
				APISecret: "test",
				URLs:      map[string]string{"doc": "test"},
				UserAgent: "cxtgo",
				Countries: []string{"test"},
				Market:    map[Symbol]MarketInfo{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBase(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCountries() = %v, want %v", got, tt.want)
			}
		})
	}
}
