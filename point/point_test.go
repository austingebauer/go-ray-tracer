package point

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/austingebauer/go-ray-tracer/vector"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "NewPoint sets x, y, and z",
			args: args{
				x: float64(1),
				y: float64(-2),
				z: float64(0),
			},
			want: &Point{
				X: 1,
				Y: -2,
				Z: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPoint(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestPoint_Equals(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		ptQ *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "vectors are equal",
			fields: fields{
				1,
				2,
				3,
			},
			args: args{
				&Point{
					1,
					2,
					3,
				},
			},
			want: true,
		},
		{
			name: "vectors aren't equal",
			fields: fields{
				1,
				2,
				3,
			},
			args: args{
				&Point{
					1,
					-2,
					3,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, pt.Equals(tt.args.ptQ))
		})
	}
}

func TestPoint_Negate(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Point
	}{
		{
			name: "negate a point positive",
			fields: fields{
				1,
				2,
				3,
			},
			want: &Point{
				-1,
				-2,
				-3,
			},
		},
		{
			name: "negate a point negative",
			fields: fields{
				-1,
				-2,
				-3,
			},
			want: &Point{
				1,
				2,
				3,
			},
		},
		{
			name: "negate an origin zero point",
			fields: fields{
				0,
				0,
				0,
			},
			want: &Point{
				0,
				0,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := pt.Negate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Scale(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		scalar float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		{
			name: "scale the point",
			fields: fields{
				0,
				2,
				-3,
			},
			args: args{
				scalar: float64(2),
			},
			want: &Point{
				0,
				4,
				-6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, pt.Scale(tt.args.scalar))
		})
	}
}

func TestPoint_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec *vector.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		{
			name: "point add vector method",
			fields: fields{
				0,
				-1,
				1,
			},
			args: args{
				&vector.Vector{
					X: 2,
					Y: 3,
					Z: -1,
				},
			},
			want: &Point{
				2,
				2,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}

			assert.Equal(t, tt.want, pt.Add(tt.args.vec))
		})
	}
}

func TestPoint_Subtract(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec *vector.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		{
			name: "point subtract vector method",
			fields: fields{
				0,
				-1,
				1,
			},
			args: args{
				&vector.Vector{
					X: 2,
					Y: 3,
					Z: -1,
				},
			},
			want: &Point{
				-2,
				-4,
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}

			assert.Equal(t, tt.want, pt.Subtract(tt.args.vec))
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		pt1 *Point
		pt2 *Point
	}
	tests := []struct {
		name string
		args args
		want *vector.Vector
	}{
		{
			name: "subtract two points for vector function",
			args: args{
				pt1: &Point{
					-1,
					0,
					1,
				},
				pt2: &Point{
					-2,
					1,
					0,
				},
			},
			want: &vector.Vector{
				1,
				-1,
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Subtract(tt.args.pt1, tt.args.pt2))
		})
	}
}
