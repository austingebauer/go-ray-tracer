// Package intersect encapsulates an intersect of a ray with an object.
package intersect

import (
	"testing"

	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
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

func TestSortIntersectionsAsc(t *testing.T) {
	type args struct {
		intersections []*Intersection
	}
	tests := []struct {
		name string
		args args
		want []*Intersection
	}{
		{
			name: "sort intersections by t values ascending",
			args: args{
				intersections: []*Intersection{
					{
						T:      10,
						Object: nil,
					},
					{
						T:      0,
						Object: nil,
					},
					{
						T:      -10,
						Object: nil,
					},
				},
			},
			want: []*Intersection{
				{
					T:      -10,
					Object: nil,
				},
				{
					T:      0,
					Object: nil,
				},
				{
					T:      10,
					Object: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortIntersectionsAsc(tt.args.intersections)
			assert.Equal(t, tt.want, tt.args.intersections)
		})
	}
}

func TestSortIntersectionsDesc(t *testing.T) {
	type args struct {
		intersections []*Intersection
	}
	tests := []struct {
		name string
		args args
		want []*Intersection
	}{
		{
			name: "sort intersections by t values ascending",
			args: args{
				intersections: []*Intersection{
					{
						T:      10,
						Object: nil,
					},
					{
						T:      0,
						Object: nil,
					},
					{
						T:      -10,
						Object: nil,
					},
				},
			},
			want: []*Intersection{
				{
					T:      10,
					Object: nil,
				},
				{
					T:      0,
					Object: nil,
				},
				{
					T:      -10,
					Object: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortIntersectionsDesc(tt.args.intersections)
			assert.Equal(t, tt.want, tt.args.intersections)
		})
	}
}

func TestPrepareComputations(t *testing.T) {
	type args struct {
		i *Intersection
		r *ray.Ray
	}
	tests := []struct {
		name    string
		args    args
		want    *IntersectionComputations
		wantErr bool
	}{
		{
			name: "Computing the state of an intersection when it occurs on the outside of the object",
			args: args{
				i: &Intersection{
					T:      4,
					Object: sphere.NewUnitSphere("testID"),
				},
				r: &ray.Ray{
					Origin:    point.NewPoint(0, 0, -5),
					Direction: vector.NewVector(0, 0, 1),
				},
			},
			want: &IntersectionComputations{
				Intersection: Intersection{
					T:      4,
					Object: sphere.NewUnitSphere("testID"),
				},
				Point:     point.NewPoint(0, 0, -1),
				EyeVec:    vector.NewVector(0, 0, -1),
				NormalVec: vector.NewVector(0, 0, -1),
				Inside:    false,
			},
			wantErr: false,
		},
		{
			name: "Computing the state of an intersection when it occurs on the Inside of the object",
			args: args{
				i: &Intersection{
					T:      1,
					Object: sphere.NewUnitSphere("testID"),
				},
				r: &ray.Ray{
					Origin:    point.NewPoint(0, 0, 0),
					Direction: vector.NewVector(0, 0, 1),
				},
			},
			want: &IntersectionComputations{
				Intersection: Intersection{
					T:      1,
					Object: sphere.NewUnitSphere("testID"),
				},
				Point:  point.NewPoint(0, 0, 1),
				EyeVec: vector.NewVector(0, 0, -1),
				// normal would've been <0,0,1>, but inverted since ray is Inside the object
				NormalVec: vector.NewVector(0, 0, -1),
				Inside:    true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrepareComputations(tt.args.i, tt.args.r)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
