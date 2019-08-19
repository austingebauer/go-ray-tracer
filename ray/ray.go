// Package ray represents a ray, or line, which has an origin and direction.
package ray

import (
	"errors"
	"github.com/austingebauer/go-ray-tracer/intersection"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
	"math"
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

// Position returns the Point that lies any distance t along the passed ray.
func Position(ray *Ray, t float64) *point.Point {
	return ray.Origin.Add(*ray.Direction.Scale(t))
}

// Intersect intersects the passed ray with the passed sphere.
//
// It returns the t values (i.e., intersection units +/- away from the origin of the Ray)
// where the Ray intersects with the sphere.
//
// If the ray intersects with the sphere at two points, then two different intersection t values are returned.
// If the ray intersects with the sphere at a single, tangent point, then two equal t values are returned.
// If the ray does not intersect with the sphere, then an empty slice is returned.
func Intersect(sphere *sphere.Sphere, ray *Ray) []*intersection.Intersection {
	// Details on calculation: https://en.wikipedia.org/wiki/Line%E2%80%93sphere_intersection

	// Transform the ray by the inverse of the transformation associated with the sphere
	// in order to use unit sphere. Moving the ray makes for more simple math and
	// same intersection results.
	sphereTransformInverse, _ := matrix.Inverse(sphere.Transform)
	transformedRay, _ := Transform(ray, sphereTransformInverse)

	// The vector from the sphere origin to the ray origin.
	sphereToRayVec := point.Subtract(transformedRay.Origin, sphere.Origin)

	// Compute the discriminant to tell whether the ray intersects with the sphere at all.
	a := vector.DotProduct(*transformedRay.Direction, *transformedRay.Direction)
	b := 2 * vector.DotProduct(*transformedRay.Direction, *sphereToRayVec)
	c := vector.DotProduct(*sphereToRayVec, *sphereToRayVec) - 1
	discriminant := math.Pow(b, 2) - 4*a*c

	// If the discriminant is negative, then the ray misses the sphere and no intersections occur.
	if discriminant < 0 {
		return []*intersection.Intersection{}
	}

	// Compute the t values.
	t1 := ((-1 * b) - math.Sqrt(discriminant)) / (2 * a)
	t2 := ((-1 * b) + math.Sqrt(discriminant)) / (2 * a)

	// Return the intersection t values and object in increasing order
	return []*intersection.Intersection{
		{
			T:      t1,
			Object: sphere,
		},
		{
			T:      t2,
			Object: sphere,
		},
	}
}

// Transform applies the passed 4x4 transformation Matrix to the passed Ray.
// Returns a new Ray with the transformed origin and direction.
func Transform(ray *Ray, m *matrix.Matrix) (*Ray, error) {
	if m.GetRows() != 4 || m.GetCols() != 4 {
		return nil, errors.New("order of matrix m must be 4x4")
	}

	// transform the ray origin
	originMatrix, _ := matrix.Multiply(m, point.ToMatrix(ray.Origin))
	transformedOriginPoint, _ := point.ToPoint(originMatrix)

	// transform the ray direction
	directionMatrix, _ := matrix.Multiply(m, vector.ToMatrix(ray.Direction))
	transformedDirectionVector, _ := vector.ToVector(directionMatrix)

	return NewRay(transformedOriginPoint, transformedDirectionVector), nil
}
