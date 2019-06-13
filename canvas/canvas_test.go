// Package canvas represents a rectangular grid of Pixels.
package canvas

import (
	"bytes"
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
			name: "new canvas has Width, Height, and default color black rgb(0,0,0)",
			args: args{
				width:  10,
				height: 20,
			},
			want: &Canvas{
				Width:  10,
				Height: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas := NewCanvas(tt.args.width, tt.args.height)
			assert.Equal(t, tt.want.Height, canvas.Height)
			assert.Equal(t, tt.want.Width, canvas.Width)
			assert.Equal(t, *NewColor(0, 0, 0), canvas.Pixels[0][0])
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
		name      string
		fields    fields
		args      args
		want      Color
		wantError bool
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
			wantError: false,
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
			wantError: true,
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
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(tt.fields.Width, tt.fields.Height)
			err := c.WritePixel(tt.args.x, tt.args.y, tt.args.color)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, c.Pixels[tt.args.y][tt.args.x])
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
		name    string
		fields  fields
		args    args
		want    *Color
		wantErr bool
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
			wantErr: false,
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
			wantErr: true,
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(tt.fields.Width, tt.fields.Height)
			err := c.WritePixel(tt.args.x, tt.args.y, tt.args.color)
			color, err := c.PixelAt(tt.args.y, tt.args.x)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, color)
			}
		})
	}
}

func TestCanvas_ToPPM(t *testing.T) {
	tests := []struct {
		name       string
		c          *Canvas
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "canvas to portable pixmap (PPM) file",
			c:          NewCanvas(10, 2),
			wantWriter: "P3\n2 2\n255\n{0 0 0}\n",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < int(tt.c.Height); i++ {
				for j := 0; j < int(tt.c.Width); j++ {
					err := tt.c.WritePixel(uint64(j), uint64(i), *NewColor(1, 0.8, 0.6))
					assert.NoError(t, err)
				}
			}

			writer := &bytes.Buffer{}
			err := tt.c.ToPPM(writer)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantWriter, writer.String())
			}
		})
	}
}
