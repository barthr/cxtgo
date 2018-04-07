package precision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		input  float64
		places int
	}
	tests := []struct {
		name        string
		args        args
		wantRounded float64
	}{
		{
			name:        "test with empty input",
			args:        args{},
			wantRounded: 0,
		},
		{
			name: "test with input decimal and low places",
			args: args{
				input:  0.1111,
				places: 1,
			},
			wantRounded: 0.1,
		},
		{
			name: "test with input decimal and high places",
			args: args{
				input:  5.3253,
				places: 4,
			},
			wantRounded: 5.3253,
		},
		{
			name: "test with input negative decimal",
			args: args{
				input:  -5.3253,
				places: 4,
			},
			wantRounded: -5.3253,
		},
		{
			name: "test with input above .5",
			args: args{
				input:  5.55,
				places: 1,
			},
			wantRounded: 5.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.wantRounded, Round(tt.args.input, tt.args.places), tt.name)
		})
	}
}
