package point

import (
	"reflect"
	"testing"

	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/stretchr/testify/assert"
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
			name: "Vectors are equal",
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
			name: "Vectors aren't equal",
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := pt.Scale(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Scale() = %v, want %v", got, tt.want)
			}
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
		vec vector.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := pt.Add(tt.args.vec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Add() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := pt.Subtract(tt.args.vec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Subtract() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.pt1, tt.args.pt2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
