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
					*point.NewPoint(0,0,-5),
					*vector.NewVector(0,1,0)),
			},
			want: color.NewColor(0,0,0),
		},
		{
			name: "color when a ray hits the world",
			args: args{
				w: NewDefaultWorld(),
				r: ray.NewRay(
					*point.NewPoint(0,0,-5),
					*vector.NewVector(0,0,1)),
			},
			want: color.NewColor(0.38066,0.47583,0.2855),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ColorAt(tt.args.w, tt.args.r)
			if !assert.True(t, color.Equals(*tt.want, *c)) {
				assert.Equal(t, tt.want, ColorAt(tt.args.w, tt.args.r))
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
		*point.NewPoint(0,0,0.75),
		*vector.NewVector(0,0,-1))

	c := ColorAt(w, r)
	if !assert.True(t, color.Equals(inner.Material.Color, *c)) {
		assert.Equal(t, inner.Material.Color, *c)
	}
}
