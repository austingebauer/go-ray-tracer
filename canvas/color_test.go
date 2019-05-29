package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	type args struct {
		red   float64
		green float64
		blue  float64
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{
			name: "create a new color",
			args: args{
				red:   1,
				green: 0,
				blue:  0,
			},
			want: &Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewColor(tt.args.red, tt.args.green, tt.args.blue))
		})
	}
}
