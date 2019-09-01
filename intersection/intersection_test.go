// Package intersection encapsulates an intersection of a ray with an object.
package intersection

import (
	"testing"

	"github.com/austingebauer/go-ray-tracer/sphere"
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
			name: "intersection encapsulates t and an object",
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
			name: "hit is always the lowest non-negative intersection",
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
