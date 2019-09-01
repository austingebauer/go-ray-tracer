package color

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

func TestAdd(t *testing.T) {
	type args struct {
		c1 Color
		c2 Color
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			name: "add color to color function",
			args: args{
				c1: Color{
					Red:   1,
					Green: 0,
					Blue:  -1,
				},
				c2: Color{
					Red:   2,
					Green: 3,
					Blue:  4,
				},
			},
			want: Color{
				Red:   3,
				Green: 3,
				Blue:  3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Add(tt.args.c1, tt.args.c2))
		})
	}
}

func TestColor_Subtract(t *testing.T) {
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
			name: "subtract color from color method",
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
				Red:   -1,
				Green: -3,
				Blue:  -5,
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
			assert.Equal(t, tt.want, c.Subtract(tt.args.c2))
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		c1 Color
		c2 Color
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			name: "subtract color from color function",
			args: args{
				c1: Color{
					Red:   1,
					Green: 0,
					Blue:  -1,
				},
				c2: Color{
					Red:   2,
					Green: 3,
					Blue:  4,
				},
			},
			want: Color{
				Red:   -1,
				Green: -3,
				Blue:  -5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Subtract(tt.args.c1, tt.args.c2))
		})
	}
}

func TestColor_Scale(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		scalar float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "scale color by scalar",
			fields: fields{
				Red:   1,
				Green: 0,
				Blue:  -1,
			},
			args: args{
				scalar: -3,
			},
			want: &Color{
				Red:   -3,
				Green: 0,
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
			assert.Equal(t, tt.want, c.Scale(tt.args.scalar))
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		c      Color
		scalar float64
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{
			name: "scale color by scalar for new color",
			args: args{
				c:      *NewColor(1, 0, -1),
				scalar: -3,
			},
			want: NewColor(-3, 0, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Scale(tt.args.c, tt.args.scalar))
		})
	}
}

func TestColor_Multiply(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		c1 Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "multiply colors method",
			fields: fields{
				Red:   1,
				Green: 0.2,
				Blue:  0.4,
			},
			args: args{
				c1: Color{
					Red:   0.9,
					Green: 1,
					Blue:  0.1,
				},
			},
			want: &Color{
				Red:   0.9,
				Green: 0.2,
				Blue:  0.04000000000000001,
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
			assert.Equal(t, tt.want, c.Multiply(tt.args.c1))
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		c1 Color
		c2 Color
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{
			name: "multiply colors method",
			args: args{
				c1: Color{
					Red:   1,
					Green: 0.2,
					Blue:  0.4,
				},
				c2: Color{
					Red:   0.9,
					Green: 1,
					Blue:  0.1,
				},
			},
			want: &Color{
				Red:   0.9,
				Green: 0.2,
				Blue:  0.04000000000000001,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Multiply(tt.args.c1, tt.args.c2))
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		c1 Color
		c2 Color
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
