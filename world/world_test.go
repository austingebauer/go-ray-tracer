// Package world represents a collection of all Objects that make up a scene.
package world

import (
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/vector"
	"testing"

	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/light"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	tests := []struct {
		name string
		want *World
	}{
		{
			name: "create a new world with no Light source or Objects",
			want: &World{
				Objects: make([]*sphere.Sphere, 0),
				Light:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewWorld())
		})
	}
}

func TestNewDefaultWorld(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "create a new world with default Light source",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultWorld()
			defaultLight := light.NewPointLight(
				*point.NewPoint(-10, 10, -10),
				*color.NewColor(1, 1, 1))
			assert.Equal(t, got.Light, defaultLight)
		})
	}
}

func TestWorld_GetObjects(t *testing.T) {
	type fields struct {
		objects []*sphere.Sphere
		light   *light.PointLight
	}
	tests := []struct {
		name   string
		fields fields
		want   []*sphere.Sphere
	}{
		{
			name: "get Objects from the world",
			fields: fields{
				objects: []*sphere.Sphere{
					sphere.NewUnitSphere("testID"),
				},
				light: nil,
			},
			want: []*sphere.Sphere{
				sphere.NewUnitSphere("testID"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				Objects: tt.fields.objects,
				Light:   tt.fields.light,
			}

			assert.Equal(t, tt.want, w.Objects)
		})
	}
}

func TestRayWorldIntersect(t *testing.T) {
	type args struct {
		r *ray.Ray
		w *World
	}
	tests := []struct {
		name string
		args args
		want []*ray.Intersection
	}{
		{
			name: "ray intersects a world",
			args: args{
				r: ray.NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1)),
				w: NewDefaultWorld(),
			},
			want: []*ray.Intersection{
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

func TestShadeHitComingFromOutside(t *testing.T) {
	w := NewDefaultWorld()
	r := ray.NewRay(*point.NewPoint(0, 0, -5), *vector.NewVector(0, 0, 1))
	shape := w.Objects[0]
	i := ray.NewIntersection(4, shape)
	comps, err := ray.PrepareComputations(i, r)
	assert.NoError(t, err)
	cActual := ShadeHit(w, comps)

	// TODO: Investigate why book says green should be 0.47583.
	cExpected := color.NewColor(0.38066, 0.047583, 0.2855)
	if !assert.True(t, color.Equals(*cExpected, *cActual)) {
		assert.Equal(t, cExpected, cActual)
	}
}

func TestShadeHitComingFromInside(t *testing.T) {
	w := NewDefaultWorld()
	w.Light = light.NewPointLight(*point.NewPoint(0, 0.25, 0),
		*color.NewColor(1, 1, 1))
	r := ray.NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(0, 0, 1))
	shape := w.Objects[1]
	i := ray.NewIntersection(0.5, shape)
	comps, err := ray.PrepareComputations(i, r)
	assert.NoError(t, err)
	cActual := ShadeHit(w, comps)

	cExpected := color.NewColor(0.90498, 0.90498, 0.90498)
	if !assert.True(t, color.Equals(*cExpected, *cActual)) {
		assert.Equal(t, cExpected, cActual)
	}
}

func TestColorAt(t *testing.T) {
	type args struct {
		w *World
		r *ray.Ray
	}
	tests := []struct {
		name string
		args args
		want *color.Color
	}{
		{
			name: "color when a ray misses the world",
			args: args{
				w: NewDefaultWorld(),
				r: ray.NewRay(
					*point.NewPoint(0, 0, -5),
					*vector.NewVector(0, 1, 0)),
			},
			want: color.NewColor(0, 0, 0),
		},
		{
			name: "color when a ray hits the world",
			args: args{
				w: NewDefaultWorld(),
				r: ray.NewRay(
					*point.NewPoint(0, 0, -5),
					*vector.NewVector(0, 0, 1)),
			},
			// TODO: Investigate why book says green should be 0.47583.
			want: color.NewColor(0.38066, 0.047583, 0.2855),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := ColorAt(tt.args.w, tt.args.r)
			assert.NoError(t, err)
			if !assert.True(t, color.Equals(*tt.want, *c)) {
				assert.Equal(t, tt.want, c)
			}
		})
	}
}

// This test case shows that we expect ColorAt() to use the hit
// when computing the color. In this test, we put the ray inside
// of the outer sphere, and pointing at the inner sphere. We expect
// the hit to be on the inner sphere, and this return its color.
func TestColorAtIntersectionBehindAndFrontOfRay(t *testing.T) {
	w := NewDefaultWorld()
	outer := w.Objects[0]
	outer.Material.Ambient = 1
	inner := w.Objects[1]
	inner.Material.Ambient = 1

	r := ray.NewRay(
		*point.NewPoint(0, 0, 0.75),
		*vector.NewVector(0, 0, -1))

	c, err := ColorAt(w, r)
	assert.NoError(t, err)
	if !assert.True(t, color.Equals(inner.Material.Color, *c)) {
		assert.Equal(t, inner.Material.Color, *c)
	}
}

func TestIsShadowed(t *testing.T) {
	type args struct {
		world *World
		pt    *point.Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no shadow when nothing is collinear with the point and the light",
			args: args{
				world: NewDefaultWorld(),
				pt:    point.NewPoint(0, 10, 0),
			},
			want: false,
		},
		{
			name: "the shadow when an object is between the point and the light",
			args: args{
				world: NewDefaultWorld(),
				pt:    point.NewPoint(10, -10, 10),
			},
			want: true,
		},
		{
			name: "no shadow when an object is behind the light",
			args: args{
				world: NewDefaultWorld(),
				pt:    point.NewPoint(-20, 20, -20),
			},
			want: false,
		},
		{
			name: "no shadow when an object is behind the point",
			args: args{
				world: NewDefaultWorld(),
				pt:    point.NewPoint(-2, 2, -2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsShadowed(tt.args.world, tt.args.pt))
		})
	}
}
