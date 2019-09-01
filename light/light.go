// Package light represents difference types of light sources.
package light

import (
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/point"
)

// PointLight represents a point light source that exists at a single point in 3D space.
// The point light source has an intensity which describes the color of the light source
// and how bright it is.
type PointLight struct {
	position  point.Point
	intensity color.Color
}

// NewPointLight returns a new PointLight having the passed position and intensity.
func NewPointLight(position point.Point, intensity color.Color) *PointLight {
	return &PointLight{
		position:  position,
		intensity: intensity,
	}
}
