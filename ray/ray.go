// Package ray represents a ray, or line, which has an origin and direction.
package ray

import (
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

// Ray is a ray, or line, which has an origin and direction.
type Ray struct {
	// Origin of the Ray
	Origin *point.Point
	// Direction vector of the Ray
	Direction *vector.Vector
}

// NewRay returns a new Ray having the passed origin and direction.
func NewRay(origin *point.Point, direction *vector.Vector) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}

// Position returns the Point that lies any distance t along the passed ray.
func Position(ray *Ray, t float64) *point.Point {
	return ray.Origin.Add(*ray.Direction.Scale(t))
}
