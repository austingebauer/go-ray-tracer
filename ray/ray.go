// Package ray represents a ray, or line, which has an origin and direction.
package ray

import (
	"errors"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

// Ray is a ray, or line, which has an origin and direction.
type Ray struct {
	Origin    *point.Point
	Direction *vector.Vector
}

// NewRay returns a new Ray having the passed origin and direction.
func NewRay(origin point.Point, direction vector.Vector) *Ray {
	return &Ray{
		Origin:    &origin,
		Direction: &direction,
	}
}

// Position returns the Point that lies any distance t along the passed ray.
func Position(ray *Ray, t float64) *point.Point {
	// Scale the direction vector by t
	scaledDirectionVec := vector.Scale(*ray.Direction, t)

	// Add the scaled vector to the origin and return the point position on the ray
	return point.Add(ray.Origin, scaledDirectionVec)
}

// Transform applies the passed 4x4 transformation Matrix to the passed Ray.
// Returns a new Ray with the transformed origin and direction.
func Transform(ray *Ray, m *matrix.Matrix) (*Ray, error) {
	if m.GetRows() != 4 || m.GetCols() != 4 {
		return nil, errors.New("order of matrix m must be 4x4")
	}

	// transform the ray origin
	originMatrix, _ := matrix.Multiply(m, matrix.PointToMatrix(ray.Origin))
	transformedOriginPoint, _ := matrix.MatrixToPoint(originMatrix)

	// transform the ray direction
	directionMatrix, _ := matrix.Multiply(m, matrix.VectorToMatrix(ray.Direction))
	transformedDirectionVector, _ := matrix.MatrixToVector(directionMatrix)

	return NewRay(*transformedOriginPoint, *transformedDirectionVector), nil
}
