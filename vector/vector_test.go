package vector

import (
	"math"
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
			name: "vectors are equal",
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
			name: "vectors aren't equal",
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
			name: "vector has magnitude",
			fields: fields{
				2,
				2,
				1,
			},
			want: 3,
		},
		{
			name: "negative component vector has same magnitude",
			fields: fields{
				-2,
				-2,
				-1,
			},
			want: 3,
		},
		{
			name: "unit vector x",
			fields: fields{
				1,
				0,
				0,
			},
			want: 1,
		},
		{
			name: "unit vector y",
			fields: fields{
				0,
				1,
				0,
			},
			want: 1,
		},
		{
			name: "unit vector z",
			fields: fields{
				0,
				0,
				1,
			},
			want: 1,
		},
		{
			name: "vector magnitude equals square root",
			fields: fields{
				1,
				2,
				3,
			},
			want: math.Sqrt(14),
		},
		{
			name: "negative vector magnitude equals square root",
			fields: fields{
				-1,
				-2,
				-3,
			},
			want: math.Sqrt(14),
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
		{
			name: "negate a vector positive",
			fields: fields{
				1,
				2,
				3,
			},
			want: &Vector{
				-1,
				-2,
				-3,
			},
		},
		{
			name: "negate a vector negative",
			fields: fields{
				-1,
				-2,
				-3,
			},
			want: &Vector{
				1,
				2,
				3,
			},
		},
		{
			name: "negate a zero vector",
			fields: fields{
				0,
				0,
				0,
			},
			want: &Vector{
				0,
				0,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Negate())
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
		{
			name: "scale the vector",
			fields: fields{
				0,
				2,
				-3,
			},
			args: args{
				scalar: float64(2),
			},
			want: &Vector{
				0,
				4,
				-6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Scale(tt.args.scalar))
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
		{
			name: "vector add vector method",
			fields: fields{
				0,
				-1,
				1,
			},
			args: args{
				&Vector{
					2,
					3,
					0,
				},
			},
			want: &Vector{
				2,
				2,
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Add(tt.args.vec2))
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
		{
			name: "vector subtract vector method",
			fields: fields{
				0,
				-1,
				2,
			},
			args: args{
				&Vector{
					0,
					2,
					3,
				},
			},
			want: &Vector{
				0,
				-3,
				-1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Subtract(tt.args.vec2))
		})
	}
}
