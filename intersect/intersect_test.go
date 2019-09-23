// Package intersect encapsulates an intersect of a ray with an object.
package intersect

import (
	"testing"

	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/austingebauer/go-ray-tracer/world"
	"github.com/stretchr/testify/assert"
)

func TestNewIntersection(t *testing.T) {
	type args struct {
		t      float64
		object sphere.Sphere
	}
	tests := []struct {
		name string
		args args
		want *Intersection
	}{
		{
			name: "intersect encapsulates t and an object",
			args: args{
				t:      -5,
				object: *sphere.NewUnitSphere("testID"),
			},
			want: &Intersection{
				T:      -5,
				Object: sphere.NewUnitSphere("testID"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewIntersection(tt.args.t, tt.args.object))
		})
	}
}

func TestIntersections(t *testing.T) {
	type args struct {
		intersections []*Intersection
	}
	tests := []struct {
		name string
		args args
		want []*Intersection
	}{
		{
			name: "aggregating intersections",
			args: args{
				intersections: []*Intersection{
					NewIntersection(1, *sphere.NewUnitSphere("testID")),
					NewIntersection(2, *sphere.NewUnitSphere("testID")),
				},
			},
			want: []*Intersection{
				NewIntersection(1, *sphere.NewUnitSphere("testID")),
				NewIntersection(2, *sphere.NewUnitSphere("testID")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Intersections(tt.args.intersections...))
		})
	}
}

func TestHit(t *testing.T) {
	type args struct {
		intersections []*Intersection
	}
	tests := []struct {
		name string
		args args
		want *Intersection
	}{
		{
			name: "hit when there are no intersections",
			args: args{
				intersections: Intersections(),
			},
			want: nil,
		},
		{
			name: "hit when all intersections have positive t",
			args: args{
				intersections: Intersections(
					NewIntersection(1, *sphere.NewUnitSphere("testID")),
					NewIntersection(2, *sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(1, *sphere.NewUnitSphere("testID")),
		},
		{
			name: "hit when some intersections have negative t",
			args: args{
				intersections: Intersections(
					NewIntersection(-1, *sphere.NewUnitSphere("testID")),
					NewIntersection(1, *sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(1, *sphere.NewUnitSphere("testID")),
		},
		{
			name: "hit when some intersections have negative t and a zero T",
			args: args{
				intersections: Intersections(
					NewIntersection(-1, *sphere.NewUnitSphere("testID")),
					NewIntersection(-9, *sphere.NewUnitSphere("testID")),
					NewIntersection(0, *sphere.NewUnitSphere("testID")),
					NewIntersection(10, *sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(0, *sphere.NewUnitSphere("testID")),
		},
		{
			name: "hit when all intersections have negative t",
			args: args{
				intersections: Intersections(
					NewIntersection(-7, *sphere.NewUnitSphere("testID")),
					NewIntersection(-3, *sphere.NewUnitSphere("testID")),
					NewIntersection(-2, *sphere.NewUnitSphere("testID")),
				),
			},
			want: nil,
		},
		{
			name: "hit is always the lowest non-negative intersect",
			args: args{
				intersections: Intersections(
					NewIntersection(5, *sphere.NewUnitSphere("testID")),
					NewIntersection(7, *sphere.NewUnitSphere("testID")),
					NewIntersection(-3, *sphere.NewUnitSphere("testID")),
					NewIntersection(2, *sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(2, *sphere.NewUnitSphere("testID")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Hit(tt.args.intersections))
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		sphere    *sphere.Sphere
		ray       *ray.Ray
		transform *matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want []*Intersection
	}{
		{
			name: "ray intersects with a sphere at two positive points. sphere is ahead of ray origin.",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    ray.NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1)),
			},
			want: []*Intersection{
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
				ray:    ray.NewRay(*point.NewPoint(0, 1, -5), *vector.NewVector(0, 0, 1)),
			},
			want: []*Intersection{
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
				ray:    ray.NewRay(*point.NewPoint(0, 2, -5), *vector.NewVector(0, 0, 1)),
			},
			want: []*Intersection{},
		},
		{
			name: "ray originates inside the sphere",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    ray.NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(0, 0, 1)),
			},
			want: []*Intersection{
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
				ray:    ray.NewRay(*point.NewPoint(0, 0, 5), *vector.NewVector(0, 0, 1)),
			},
			want: []*Intersection{
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
			intersections := RaySphereIntersect(tt.args.ray, tt.args.sphere)

			assert.Equal(t, len(tt.want), len(intersections))
			assert.Equal(t, tt.want, intersections)
		})
	}
}

func TestIntersectWithSphereTransform(t *testing.T) {
	type args struct {
		sphere    *sphere.Sphere
		ray       *ray.Ray
		transform *matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want []*Intersection
	}{
		{
			name: "intersecting a scaled unit sphere with a ray",
			args: args{
				sphere:    sphere.NewUnitSphere("testID"),
				ray:       ray.NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1)),
				transform: matrix.NewScalingMatrix(2, 2, 2),
			},
			want: []*Intersection{
				{
					T:      3.0,
					Object: sphere.NewUnitSphere("testID"),
				},
				{
					T:      7.0,
					Object: sphere.NewUnitSphere("testID"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.sphere.SetTransform(tt.args.transform)
			intersections := RaySphereIntersect(tt.args.ray, tt.args.sphere)

			assert.Equal(t, len(tt.want), len(intersections))

			// only check T values of intersections
			assert.Equal(t, tt.want[0].T, intersections[0].T)
			assert.Equal(t, tt.want[1].T, intersections[1].T)
		})
	}
}

func TestRayWorldIntersect(t *testing.T) {
	type args struct {
		r *ray.Ray
		w *world.World
	}
	tests := []struct {
		name string
		args args
		want []*Intersection
	}{
		{
			name: "ray intersects a world",
			args: args{
				r: ray.NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1)),
				w: world.NewDefaultWorld(),
			},
			want: []*Intersection{
				{
					T:      4,
					Object: nil,
				},
				{
					T:      4.5,
					Object: nil,
				},
				{
					T:      5.5,
					Object: nil,
				},
				{
					T:      6,
					Object: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIntersections := RayWorldIntersect(tt.args.r, tt.args.w)
			assert.Equal(t, len(tt.want), len(actualIntersections))

			// Each actual intersection T value matches the expected T value
			for idx, intersection := range actualIntersections {
				assert.Equal(t, tt.want[idx].T, intersection.T)
			}
		})
	}
}
