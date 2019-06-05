// Package canvas represents a rectangular grid of pixels.
package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		width  uint64
		height uint64
	}
	tests := []struct {
		name string
		args args
		want *Canvas
	}{
		{
			name: "new canvas has width, height, and default color black rgb(0,0,0)",
			args: args{
				width:  10,
				height: 20,
			},
			want: &Canvas{
				width:  10,
				height: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas := NewCanvas(tt.args.width, tt.args.height)
			assert.Equal(t, tt.want.height, canvas.height)
			assert.Equal(t, tt.want.width, canvas.width)
			assert.Equal(t, *NewColor(0, 0, 0), canvas.pixels[0][0])
		})
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	type fields struct {
		Width  uint64
		Height uint64
	}
	type args struct {
		x     uint64
		y     uint64
		color Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
		error  bool
	}{
		{
			name: "write color to pixel on the canvas",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x: 9,
				y: 19,
				color: Color{
					Red:   1,
					Green: 0,
					Blue:  0,
				},
			},
			want: Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
			error: false,
		},
		{
			name: "write color to pixel out of bounds x",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x: 100,
				y: 2,
				color: Color{
					Red:   1,
					Green: 0,
					Blue:  0,
				},
			},
			want: Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
			error: true,
		},
		{
			name: "write color to pixel out of bounds y",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x: 8,
				y: 20,
				color: Color{
					Red:   1,
					Green: 0,
					Blue:  0,
				},
			},
			want: Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
			error: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(tt.fields.Width, tt.fields.Height)
			err := c.WritePixel(tt.args.x, tt.args.y, tt.args.color)
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, c.pixels[tt.args.y][tt.args.x])
			}
		})
	}
}

func TestCanvas_PixelAt(t *testing.T) {
	type fields struct {
		Width  uint64
		Height uint64
	}
	type args struct {
		x     uint64
		y     uint64
		color Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
		error  bool
	}{
		{
			name: "pixel at x and y on the canvas",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x: 2,
				y: 2,
				color: Color{
					Red:   1,
					Green: 0,
					Blue:  0,
				},
			},
			want: &Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
			error: false,
		},
		{
			name: "pixel at x and y out of bounds on x",
			fields: fields{
				Width:  1,
				Height: 2,
			},
			args: args{
				x: 20,
				y: 0,
				color: Color{
					Red:   1,
					Green: 0,
					Blue:  0,
				},
			},
			want: &Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
			error: true,
		},
		{
			name: "pixel at x and y out of bounds on y",
			fields: fields{
				Width:  1,
				Height: 2,
			},
			args: args{
				x: 0,
				y: 2,
				color: Color{
					Red:   1,
					Green: 0,
					Blue:  0,
				},
			},
			want: &Color{
				Red:   1,
				Green: 0,
				Blue:  0,
			},
			error: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(tt.fields.Width, tt.fields.Height)
			err := c.WritePixel(tt.args.x, tt.args.y, tt.args.color)
			color, err := c.PixelAt(tt.args.y, tt.args.x)
			if tt.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, color)
			}
		})
	}
}
