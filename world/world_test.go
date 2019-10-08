// Package world represents a collection of all Objects that make up a scene.
package world

import (
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
