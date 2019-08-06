package ray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

func TestNewRay(t *testing.T) {
	type args struct {
		origin    *point.Point
		direction *vector.Vector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ray has origin and direction",
			args: args{
				origin:    point.NewPoint(1, 2, 3),
				direction: vector.NewVector(4, 5, 6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRay(tt.args.origin, tt.args.direction)
			assert.Equal(t, tt.args.origin, r.Origin)
			assert.Equal(t, tt.args.direction, r.Direction)
		})
	}
}

func TestPosition(t *testing.T) {
	type args struct {
		ray *Ray
		t   float64
	}
	tests := []struct {
		name string
		args args
		want *point.Point
	}{
		{
			name: "compute point that lies distance t along the ray 1",
			args: args{
				ray: NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
				t:   0,
			},
			want: point.NewPoint(2, 3, 4),
		},
		{
			name: "compute point that lies distance t along the ray 2",
			args: args{
				ray: NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
				t:   1,
			},
			want: point.NewPoint(3, 3, 4),
		},
		{
			name: "compute point that lies distance t along the ray 3",
			args: args{
				ray: NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
				t:   -1,
			},
			want: point.NewPoint(1, 3, 4),
		},
		{
			name: "compute point that lies distance t along the ray 4",
			args: args{
				ray: NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0)),
				t:   2.5,
			},
			want: point.NewPoint(4.5, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Position(tt.args.ray, tt.args.t))
		})
	}
}
