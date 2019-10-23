// Package light provides different types of light sources and behavior using light.
package light

import (
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/point"
)

// PointLight represents a point light source that exists at a single point in 3D space.
// The point light source has an Intensity which describes the color of the light source
// and how bright it is.
type PointLight struct {
	Position  point.Point
	Intensity color.Color
}

// NewPointLight returns a new PointLight having the passed Position and Intensity.
func NewPointLight(position point.Point, intensity color.Color) *PointLight {
	return &PointLight{
		Position:  position,
		Intensity: intensity,
	}
}
