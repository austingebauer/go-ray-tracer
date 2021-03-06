package ray

import (
	"github.com/austingebauer/go-ray-tracer/maths"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"testing"

	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/stretchr/testify/assert"
)

func TestNewIntersection(t *testing.T) {
	type args struct {
		t      float64
		object *sphere.Sphere
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
				object: sphere.NewUnitSphere("testID"),
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
					NewIntersection(1, sphere.NewUnitSphere("testID")),
					NewIntersection(2, sphere.NewUnitSphere("testID")),
				},
			},
			want: []*Intersection{
				NewIntersection(1, sphere.NewUnitSphere("testID")),
				NewIntersection(2, sphere.NewUnitSphere("testID")),
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
					NewIntersection(1, sphere.NewUnitSphere("testID")),
					NewIntersection(2, sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(1, sphere.NewUnitSphere("testID")),
		},
		{
			name: "hit when some intersections have negative t",
			args: args{
				intersections: Intersections(
					NewIntersection(-1, sphere.NewUnitSphere("testID")),
					NewIntersection(1, sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(1, sphere.NewUnitSphere("testID")),
		},
		{
			name: "hit when some intersections have negative t and a zero T",
			args: args{
				intersections: Intersections(
					NewIntersection(-1, sphere.NewUnitSphere("testID")),
					NewIntersection(-9, sphere.NewUnitSphere("testID")),
					NewIntersection(0, sphere.NewUnitSphere("testID")),
					NewIntersection(10, sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(0, sphere.NewUnitSphere("testID")),
		},
		{
			name: "hit when all intersections have negative t",
			args: args{
				intersections: Intersections(
					NewIntersection(-7, sphere.NewUnitSphere("testID")),
					NewIntersection(-3, sphere.NewUnitSphere("testID")),
					NewIntersection(-2, sphere.NewUnitSphere("testID")),
				),
			},
			want: nil,
		},
		{
			name: "hit is always the lowest non-negative intersect",
			args: args{
				intersections: Intersections(
					NewIntersection(5, sphere.NewUnitSphere("testID")),
					NewIntersection(7, sphere.NewUnitSphere("testID")),
					NewIntersection(-3, sphere.NewUnitSphere("testID")),
					NewIntersection(2, sphere.NewUnitSphere("testID")),
				),
			},
			want: NewIntersection(2, sphere.NewUnitSphere("testID")),
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
		r *Ray
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
				r: &Ray{
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
				OverPoint: point.NewPoint(0, 0, -1.00001),
				EyeVec:    vector.NewVector(0, 0, -1),
				NormalVec: vector.NewVector(0, 0, -1),
				Inside:    false,
			},
			wantErr: false,
		},
		{
			name: "Computing the state of an intersection when it occurs on the inside of the object",
			args: args{
				i: &Intersection{
					T:      1,
					Object: sphere.NewUnitSphere("testID"),
				},
				r: &Ray{
					Origin:    point.NewPoint(0, 0, 0),
					Direction: vector.NewVector(0, 0, 1),
				},
			},
			want: &IntersectionComputations{
				Intersection: Intersection{
					T:      1,
					Object: sphere.NewUnitSphere("testID"),
				},
				Point:     point.NewPoint(0, 0, 1),
				OverPoint: point.NewPoint(0, 0, 0.99999),
				EyeVec:    vector.NewVector(0, 0, -1),
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

func TestOverPoint(t *testing.T) {
	// Assert that the hit offsets the over point field to avoid shadow acne
	r := NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1))
	shape := sphere.NewUnitSphere("shape")
	shape.SetTransform(matrix.NewTranslationMatrix(0, 0, 1))
	i := NewIntersection(5, shape)
	comps, err := PrepareComputations(i, r)
	assert.NoError(t, err)

	// The over point should be less (further along negative z-axis) than -1 * Epsilon / 2
	assert.Less(t, comps.OverPoint.Z, -1*maths.Epsilon/2)

	// The original point of intersection should be greater than the over point
	assert.Greater(t, comps.Point.Z, comps.OverPoint.Z)
}

func TestRaySphereIntersect(t *testing.T) {
	type args struct {
		sphere    *sphere.Sphere
		ray       *Ray
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
				ray:    NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1)),
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
			name: "ray is tangent to the sphere at one Point of t",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(*point.NewPoint(0, 1, -5), *vector.NewVector(0, 0, 1)),
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
				ray:    NewRay(*point.NewPoint(0, 2, -5), *vector.NewVector(0, 0, 1)),
			},
			want: []*Intersection{},
		},
		{
			name: "ray originates Inside the sphere",
			args: args{
				sphere: sphere.NewUnitSphere("testID"),
				ray:    NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(0, 0, 1)),
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
				ray:    NewRay(*point.NewPoint(0, 0, 5), *vector.NewVector(0, 0, 1)),
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
		ray       *Ray
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
				ray:       NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1)),
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
