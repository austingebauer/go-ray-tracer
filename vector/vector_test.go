package vector

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "NewPoint sets x, y, and z",
			args: args{
				x: float64(1),
				y: float64(-2),
				z: float64(0),
			},
			want: &Vector{
				1,
				-2,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewVector(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestVector_Equals(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vecQ *Vector
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
				&Vector{
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
				&Vector{
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
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Equals(tt.args.vecQ))
		})
	}
}

func TestVector_Magnitude(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Vector has magnitude",
			fields: fields{
				2,
				2,
				1,
			},
			want: 3,
		},
		{
			name: "Negative component vector has same magnitude",
			fields: fields{
				-2,
				-2,
				-1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Magnitude())
		})
	}
}

func TestVector_Negate(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vec.Negate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Scale(t *testing.T) {
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
		want   *Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vec.Scale(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec2 *Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vec.Add(tt.args.vec2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Subtract(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec2 *Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vec.Subtract(tt.args.vec2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
