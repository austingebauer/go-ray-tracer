// Package sphere represents a sphere object with an origin and radius.
package sphere

import (
	"github.com/austingebauer/go-ray-tracer/intersect"
	"github.com/austingebauer/go-ray-tracer/material"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/vector"
	"math"
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

// RaySphereIntersect intersects the passed ray with the passed sphere.
//
// It returns the t values (i.e., intersection units +/- away from the origin of the Ray)
// where the Ray intersects with the sphere.
//
// If the ray intersects with the sphere at two points, then two different intersection t values are returned.
// If the ray intersects with the sphere at a single, tangent Point, then two equal t values are returned.
// If the ray does not intersect with the sphere, then an empty slice is returned.
func RaySphereIntersect(r *ray.Ray, s *Sphere) []*intersect.Intersection {
	// Details on calculation: https://en.wikipedia.org/wiki/Line%E2%80%93sphere_intersection

	// Transform the r by the inverse of the transformation associated with the s
	// in order to use unit s. Moving the r makes for more simple math and
	// same intersection results.
	sphereTransformInverse, _ := matrix.Inverse(s.Transform)
	transformedRay, _ := ray.Transform(r, sphereTransformInverse)

	// The vector from the s origin to the r origin.
	sphereToRayVec := point.Subtract(*transformedRay.Origin, *s.Origin)

	// Compute the discriminant to tell whether the r intersects with the s at all.
	a := vector.DotProduct(*transformedRay.Direction, *transformedRay.Direction)
	b := 2 * vector.DotProduct(*transformedRay.Direction, *sphereToRayVec)
	c := vector.DotProduct(*sphereToRayVec, *sphereToRayVec) - 1
	discriminant := math.Pow(b, 2) - 4*a*c

	// If the discriminant is negative, then the r misses the s and no intersections occur.
	if discriminant < 0 {
		return []*intersect.Intersection{}
	}

	// Compute the t values.
	t1 := ((-1 * b) - math.Sqrt(discriminant)) / (2 * a)
	t2 := ((-1 * b) + math.Sqrt(discriminant)) / (2 * a)

	// Return the intersection t values and object in increasing order
	return []*intersect.Intersection{
		{
			T:      t1,
			Object: s,
		},
		{
			T:      t2,
			Object: s,
		},
	}
}
