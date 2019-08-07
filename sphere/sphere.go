// Sphere represents a sphere object with an origin and radius.
package sphere

import (
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/hashicorp/go-uuid"
)

// Sphere is a sphere object with an origin and radius.
type Sphere struct {
	id     string
	Origin *point.Point
	Radius float64
}

// NewUnitSphere returns a new Sphere with origin (0,0,0) and a radius of 1.
func NewUnitSphere() *Sphere {
	return NewSphere(point.NewPoint(0, 0, 0), 1.0)
}

// NewSphere returns a new Sphere with the passed origin and radius.
func NewSphere(origin *point.Point, radius float64) *Sphere {
	// Each sphere as a unique ID
	id, _ := uuid.GenerateUUID()
	return &Sphere{
		id:     id,
		Origin: origin,
		Radius: radius,
	}
}
