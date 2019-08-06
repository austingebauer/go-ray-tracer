// Package ray represents a ray, or line, which has an origin and direction.
package ray

import (
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

// Ray is a ray, or line, which has an origin and direction.
type Ray struct {
	Origin    *point.Point
	Direction *vector.Vector
}

// NewRay returns a new Ray having the passed origin and direction.
func NewRay(origin *point.Point, direction *vector.Vector) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}
