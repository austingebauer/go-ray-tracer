// Package world represents a collection of all objects that make up a scene.
package world

import (
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorld(t *testing.T) {
	tests := []struct {
		name string
		want *World
	}{
		{
			name: "creating a new world",
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
