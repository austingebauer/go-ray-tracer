// Package sphere represents a sphere object with an origin and radius.
package sphere

import (
	"github.com/austingebauer/go-ray-tracer/material"
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
	Material  *material.Material
}

// NewUnitSphere returns a new Sphere with id, origin (0,0,0), and a radius of 1.
func NewUnitSphere(id string) *Sphere {
	return NewSphere(id, *point.NewPoint(0, 0, 0), 1.0)
}

// NewSphere returns a new Sphere with the passed id, origin, and radius.
func NewSphere(id string, origin point.Point, radius float64) *Sphere {
	return &Sphere{
		Id:        id,
		Origin:    &origin,
		Radius:    radius,
		Transform: matrix.NewIdentityMatrix(4),
		Material:  material.NewDefaultMaterial(),
	}
}

// SetTransform sets the transform of this Sphere.
func (s *Sphere) SetTransform(m *matrix.Matrix) {
	s.Transform = m
}

// NormalAt returns the normal vector on the passed Sphere, at the passed Point.
// The function assumes that the passed Point will always be on the surface of the sphere.
func NormalAt(s *Sphere, worldSpacePoint *point.Point) (*vector.Vector, error) {
	// Get the inverse of the transform applied to the sphere
	inverseTransform, err := matrix.Inverse(s.Transform)
	if err != nil {
		return nil, err
	}

	// Convert the passed point in world space into a point in object space
	objectSpacePointM, err := matrix.Multiply(inverseTransform, point.ToMatrix(worldSpacePoint))
	if err != nil {
		return nil, err
	}
	objectSpacePoint, err := point.ToPoint(objectSpacePointM)
	if err != nil {
		return nil, err
	}

	// Get the normal vector in object space by subtracting the sphere
	// origin (always point(0,0,0)) from the object space point.
	objectSpaceNormal := point.Subtract(*objectSpacePoint, *s.Origin).Normalize()

	// Convert the object space normal vector back to world space by multiplying
	// by the transposed, inverse of the transform applied to the sphere.
	transposedInverseTransform := matrix.Transpose(*inverseTransform)
	worldSpaceNormalM, err := matrix.Multiply(transposedInverseTransform, vector.ToMatrix(objectSpaceNormal))
	if err != nil {
		return nil, err
	}

	// Normalize and return the world space normal vector
	worldSpaceNormalVector, err := vector.ToVector(worldSpaceNormalM)
	if err != nil {
		return nil, err
	}

	return worldSpaceNormalVector.Normalize(), nil
}
