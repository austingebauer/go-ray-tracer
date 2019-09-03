package maths

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64Equals(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "float64s are equal",
			args: args{
				a: 1.1234001,
				b: 1.1234001,
				c: Epsilon,
			},
			want: true,
		},
		{
			name: "float64s are barely equal",
			args: args{
				a: 1.0,
				b: 0.99999,
				c: Epsilon,
			},
			want: true,
		},
		{
			name: "float64s are not equal",
			args: args{
				a: 2.1234000,
				b: -1.12,
				c: Epsilon,
			},
			want: false,
		},
		{
			name: "float64s are barely not equal",
			args: args{
				a: 1.0,
				b: 0.99998,
				c: Epsilon,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Float64Equals(tt.args.a, tt.args.b, tt.args.c))
		})
	}
}

func TestToRadians(t *testing.T) {
	type args struct {
		degrees float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "30 degrees to radians",
			args: args{
				degrees: 30,
			},
			want: math.Pi / 6,
		},
		{
			name: "90 degrees to radians",
			args: args{
				degrees: 90,
			},
			want: math.Pi / 2,
		},
		{
			name: "180 degrees to radians",
			args: args{
				degrees: 180,
			},
			want: math.Pi,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rad := ToRadians(tt.args.degrees)
			assert.True(t, Float64Equals(tt.want, rad, Epsilon))
		})
	}
}

func TestToDegrees(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "pi/6 radians to degrees",
			args: args{
				radians: math.Pi / 6,
			},
			want: 30,
		},
		{
			name: "pi/2 radians to degrees",
			args: args{
				radians: math.Pi / 2,
			},
			want: 90,
		},
		{
			name: "pi radians to degrees",
			args: args{
				radians: math.Pi,
			},
			want: 180,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deg := ToDegrees(tt.args.radians)
			assert.True(t, Float64Equals(tt.want, deg, Epsilon))
		})
	}
}
