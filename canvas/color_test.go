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

func TestColor_Add(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		c2 Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "add color to color method",
			fields: fields{
				Red:   0.1,
				Green: 0.3,
				Blue:  3.3,
			},
			args: args{
				c2: Color{
					Red:   0.1,
					Green: 0.1,
					Blue:  0.1,
				},
			},
			want: &Color{
				Red:   0.2,
				Green: 0.4,
				Blue:  3.4,
			},
		},
		{
			name: "add color to color method",
			fields: fields{
				Red:   1,
				Green: 0,
				Blue:  -1,
			},
			args: args{
				c2: Color{
					Red:   2,
					Green: 3,
					Blue:  4,
				},
			},
			want: &Color{
				Red:   3,
				Green: 3,
				Blue:  3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				Red:   tt.fields.Red,
				Green: tt.fields.Green,
				Blue:  tt.fields.Blue,
			}
			assert.Equal(t, tt.want, c.Add(tt.args.c2))
		})
	}
}
