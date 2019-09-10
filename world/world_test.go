// Package world represents a collection of all objects that make up a scene.
package world

import (
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/light"
	"github.com/austingebauer/go-ray-tracer/point"
	"testing"

	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	tests := []struct {
		name string
		want *World
	}{
		{
			name: "create a new world with no light source or objects",
			want: &World{
				objects: make([]*sphere.Sphere, 0),
				light:   nil,
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
			name: "create a new world with default light source",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultWorld()
			defaultLight := light.NewPointLight(
				*point.NewPoint(-10, 10, -10),
				*color.NewColor(1, 1, 1))
			assert.Equal(t, got.light, defaultLight)
		})
	}
}
