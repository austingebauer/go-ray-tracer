// Package canvas represents a rectangular grid of Pixels.
package canvas

import (
	"bytes"
	"io"
	"testing"

	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		width  int
		height int
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
			assert.Equal(t, color.NewColor(0, 0, 0), canvas.Pixels[0][0])
		})
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	type args struct {
		x     int
		y     int
		color *color.Color
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *color.Color
		wantError bool
	}{
		{
			name: "write color to pixel on the canvas",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x:     9,
				y:     19,
				color: color.NewColor(1, 0, 0),
			},
			want:      color.NewColor(1, 0, 0),
			wantError: false,
		},
		{
			name: "write color to pixel out of bounds x",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x:     100,
				y:     2,
				color: color.NewColor(1, 0, 0),
			},
			want:      color.NewColor(1, 0, 0),
			wantError: true,
		},
		{
			name: "write color to pixel out of bounds y",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x:     8,
				y:     20,
				color: color.NewColor(1, 0, 0),
			},
			want:      color.NewColor(1, 0, 0),
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
		Width  int
		Height int
	}
	type args struct {
		x     int
		y     int
		color *color.Color
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *color.Color
		wantErr bool
	}{
		{
			name: "pixel at x and y on the canvas",
			fields: fields{
				Width:  10,
				Height: 20,
			},
			args: args{
				x:     2,
				y:     2,
				color: color.NewColor(1, 0, 0),
			},
			want: &color.Color{
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
				x:     20,
				y:     0,
				color: color.NewColor(1, 0, 0),
			},
			want: &color.Color{
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
				x:     0,
				y:     2,
				color: color.NewColor(1, 0, 0),
			},
			want: &color.Color{
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
			errWrite := c.WritePixel(tt.args.x, tt.args.y, tt.args.color)
			clr, errPixelAt := c.PixelAt(tt.args.y, tt.args.x)
			if tt.wantErr {
				assert.Error(t, errWrite)
				assert.Error(t, errPixelAt)
				assert.Nil(t, clr)
			} else {
				assert.NoError(t, errWrite)
				assert.NoError(t, errPixelAt)
				assert.Equal(t, tt.want, clr)
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
			name: "canvas to portable pixmap (PPM) file 10x2",
			c:    NewCanvas(10, 2),
			wantWriter: "P3\n" +
				"10 2\n" +
				"255\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n" +
				"255 204 153\n",
			wantErr: false,
		},
		{
			name: "canvas to portable pixmap (PPM) file 0x0",
			c:    NewCanvas(0, 0),
			wantWriter: "P3\n" +
				"0 0\n" +
				"255\n",
			wantErr: false,
		},
		{
			name: "canvas to portable pixmap (PPM) file 1x0",
			c:    NewCanvas(1, 0),
			wantWriter: "P3\n" +
				"1 0\n" +
				"255\n",
			wantErr: false,
		},
		{
			name: "canvas to portable pixmap (PPM) file 0x1",
			c:    NewCanvas(0, 1),
			wantWriter: "P3\n" +
				"0 1\n" +
				"255\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < int(tt.c.Height); i++ {
				for j := 0; j < int(tt.c.Width); j++ {
					err := tt.c.WritePixel(
						j,
						i,
						color.NewColor(1, 0.8, 0.6))
					assert.NoError(t, err)
				}
			}

			writer := &bytes.Buffer{}
			if tt.wantErr {
				badTemplate := "{ does not .work template }"
				err := tt.c.ToPPM(writer, badTemplate)
				assert.Error(t, err)
			} else {
				err := tt.c.ToPPM(writer, PixelMapTemplate)
				assert.NoError(t, err)
				str := writer.String()
				assert.Equal(t, tt.wantWriter, str)
			}
		})
	}
}

func TestCanvas_ToPPMError(t *testing.T) {
	tests := []struct {
		name       string
		c          *Canvas
		writer     io.Writer
		goTemplate string
	}{
		{
			name:       "canvas to portable pixmap (PPM) nil writer",
			c:          NewCanvas(1, 1),
			writer:     nil,
			goTemplate: PixelMapTemplate,
		},
		{
			name:       "canvas to portable pixmap (PPM) template parse error",
			c:          NewCanvas(1, 1),
			writer:     &bytes.Buffer{},
			goTemplate: "{{not a . good .template.}",
		},
		{
			name:       "canvas to portable pixmap (PPM) execution error",
			c:          NewCanvas(1, 1),
			writer:     &bytes.Buffer{},
			goTemplate: "{{ .PPMIdentifier }} and not real fields {{ .MadeUpField }}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.ToPPM(tt.writer, tt.goTemplate)
			assert.Error(t, err)
		})
	}
}

func TestCanvas_ValidateInCanvasBounds(t *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "x and y pixels in bounds of canvas",
			fields: fields{
				Width:  100,
				Height: 100,
			},
			args: args{
				x: 0,
				y: 0,
			},
			wantErr: false,
		},
		{
			name: "x and y pixels in bounds of canvas",
			fields: fields{
				Width:  100,
				Height: 100,
			},
			args: args{
				x: 99,
				y: 99,
			},
			wantErr: false,
		},
		{
			name: "x and y pixels out of bounds of canvas positive width",
			fields: fields{
				Width:  100,
				Height: 100,
			},
			args: args{
				x: 100,
				y: 0,
			},
			wantErr: true,
		},
		{
			name: "x and y pixels out of bounds of canvas positive height",
			fields: fields{
				Width:  100,
				Height: 100,
			},
			args: args{
				x: 0,
				y: 100,
			},
			wantErr: true,
		},
		{
			name: "x and y pixels out of bounds of canvas negative width",
			fields: fields{
				Width:  100,
				Height: 100,
			},
			args: args{
				x: -1,
				y: 0,
			},
			wantErr: true,
		},
		{
			name: "x and y pixels out of bounds of canvas negative height",
			fields: fields{
				Width:  100,
				Height: 100,
			},
			args: args{
				x: 0,
				y: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(tt.fields.Width, tt.fields.Height)

			if tt.wantErr {
				assert.Error(t, c.ValidateInCanvasBounds(tt.args.x, tt.args.y))
			} else {
				assert.NoError(t, c.ValidateInCanvasBounds(tt.args.x, tt.args.y))
			}
		})
	}
}
