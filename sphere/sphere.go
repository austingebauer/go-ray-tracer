// Package sphere represents a sphere object with an origin and radius.
package sphere

import (
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

// Sphere is a sphere object with an origin and radius.
type Sphere struct {
	Id        string
	Origin    *point.Point
	Radius    float64
	Transform *matrix.Matrix
}

// NewUnitSphere returns a new Sphere with id, origin (0,0,0), and a radius of 1.
func NewUnitSphere(id string) *Sphere {
	return NewSphere(id, point.NewPoint(0, 0, 0), 1.0)
}

// NewSphere returns a new Sphere with the passed id, origin, and radius.
func NewSphere(id string, origin *point.Point, radius float64) *Sphere {
	return &Sphere{
		Id:        id,
		Origin:    origin,
		Radius:    radius,
		Transform: matrix.NewIdentityMatrix(4),
	}
}

// SetTransform sets the transform of this Sphere.
func (s *Sphere) SetTransform(m *matrix.Matrix) {
	s.Transform = m
}

// NormalAt returns the normal vector on the passed Sphere, at the passed Point.
// The function assumes that the passed Point will always be on the surface of the sphere.
func NormalAt(s *Sphere, p *point.Point) *vector.Vector {
	return point.Subtract(p, s.Origin).Normalize()
}
