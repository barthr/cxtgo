package precision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			assert.Equal(tt.want, FromString(tt.args.input, tt.args.splitter), tt.name)
		})
	}
}
