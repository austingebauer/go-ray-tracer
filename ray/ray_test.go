package ray

import (
	"github.com/austingebauer/go-ray-tracer/intersection"
	"testing"

	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/stretchr/testify/assert"
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

func TestIntersect(t *testing.T) {
	type args struct {
		sphere *sphere.Sphere
		ray    *Ray
	}
	tests := []struct {
		name string
		args args
		want []*intersection.Intersection
	}{
		{
			name: "ray intersects with a sphere at two positive points. sphere is ahead of ray origin.",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1)),
			},
			want: []*intersection.Intersection{
				{
					T:      4.0,
					Object: sphere.NewUnitSphere("testID"),
				},
				{
					T:      6.0,
					Object: sphere.NewUnitSphere("testID"),
				},
			},
		},
		{
			name: "ray is tangent to the sphere at one point of t",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(point.NewPoint(0, 1, -5), vector.NewVector(0, 0, 1)),
			},
			want: []*intersection.Intersection{
				{
					T:      5.0,
					Object: sphere.NewUnitSphere("testID"),
				},
				{
					T:      5.0,
					Object: sphere.NewUnitSphere("testID"),
				},
			},
		},
		{
			name: "ray misses the sphere",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(point.NewPoint(0, 2, -5), vector.NewVector(0, 0, 1)),
			},
			want: []*intersection.Intersection{},
		},
		{
			name: "ray originates inside the sphere",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(point.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1)),
			},
			want: []*intersection.Intersection{
				{
					T:      -1.0,
					Object: sphere.NewUnitSphere("testID"),
				},
				{
					T:      1.0,
					Object: sphere.NewUnitSphere("testID"),
				},
			},
		},
		{
			name: "ray intersects with a sphere at two negative points. sphere is behind ray origin.",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(point.NewPoint(0, 0, 5), vector.NewVector(0, 0, 1)),
			},
			want: []*intersection.Intersection{
				{
					T:      -6.0,
					Object: sphere.NewUnitSphere("testID"),
				},
				{
					T:      -4.0,
					Object: sphere.NewUnitSphere("testID"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intersections := Intersect(tt.args.sphere, tt.args.ray)

			assert.Equal(t, len(tt.want), len(intersections))
			assert.Equal(t, tt.want, intersections)
		})
	}
}
