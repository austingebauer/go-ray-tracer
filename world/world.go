// Package world represents a collection of all objects that make up a scene.
package world

import (
	"github.com/austingebauer/go-ray-tracer/light"
	"github.com/austingebauer/go-ray-tracer/sphere"
)

// World represents a collection of all objects that make up a scene.
type World struct {
	objects []*sphere.Sphere
	light   *light.PointLight
}

// NewWorld returns a new World.
func NewWorld() *World {
	return &World{
		objects: make([]*sphere.Sphere, 0),
		light:   nil,
	}
}

// NewDefaultWorld returns a new default world, which contains
// two spheres and a point light source.
func NewDefaultWorld() *World {
	// TODO: add spheres and light source
	return &World{
		objects: make([]*sphere.Sphere, 0),
		light:   nil,
	}
}
