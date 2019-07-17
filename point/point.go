// Package point represents a point in a left-handed 3D coordinate system.
package point

import (
	"github.com/austingebauer/go-ray-tracer/math_utils"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/vector"
)

const (
	pointW = 1
)

// Point represents a point in a left-handed 3D coordinate system
type Point struct {
	// X, Y, and Z represent components in a left-handed 3D coordinate system
	X, Y, Z float64
}

// NewPoint returns a new Point that has the passed X, Y, and Z values.
func NewPoint(x, y, z float64) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}

// Equals returns true if the passed Point is equal to this Point.
// Two Points are equal if their X, Y, Z components are equal.
func (pt *Point) Equals(ptQ *Point) bool {
	return math_utils.Float64Equals(pt.X, ptQ.X, math_utils.Epsilon) &&
		math_utils.Float64Equals(pt.Y, ptQ.Y, math_utils.Epsilon) &&
		math_utils.Float64Equals(pt.Z, ptQ.Z, math_utils.Epsilon)
}

// Negate multiplies each of this Point's components by -1.
func (pt *Point) Negate() *Point {
	return pt.Scale(-1)
}

// Scale multiplies each of this Point's components by the passed scalar value.
func (pt *Point) Scale(scalar float64) *Point {
	pt.X = pt.X * scalar
	pt.Y = pt.Y * scalar
	pt.Z = pt.Z * scalar
	return pt
}

// Add modifies each component of this Point by setting each of them
// to the sum of the components in this Point and the passed Vector.
func (pt *Point) Add(vec vector.Vector) *Point {
	pt.X = pt.X + vec.X
	pt.Y = pt.Y + vec.Y
	pt.Z = pt.Z + vec.Z
	return pt
}

// Add returns a new Point with components equal to the sum
// of the corresponding components in the passed Point and Vector.
func Add(pt Point, vec vector.Vector) Point {
	return Point{
		X: pt.X + vec.X,
		Y: pt.Y + vec.Y,
		Z: pt.Z + vec.Z,
	}
}

// Subtract modifies each component of this Point by setting each of them
// to the difference of the components in this Point and the passed Vector.
func (pt *Point) Subtract(vec vector.Vector) *Point {
	pt.X = pt.X - vec.X
	pt.Y = pt.Y - vec.Y
	pt.Z = pt.Z - vec.Z
	return pt
}

// Subtract returns a new Vector with components equal to the
// difference of the corresponding components in the passed Points.
func Subtract(pt1 Point, pt2 Point) vector.Vector {
	return vector.Vector{
		X: pt1.X - pt2.X,
		Y: pt1.Y - pt2.Y,
		Z: pt1.Z - pt2.Z,
	}
}

// ToMatrix returns a 4x1 Matrix that represents the passed Point.
func ToMatrix(pt Point) *matrix.Matrix {
	m := matrix.NewMatrix(4, 1)
	_ := m.SetValue(0, 0, pt.X)
	_ = m.SetValue(1, 0, pt.Y)
	_ = m.SetValue(2, 0, pt.Z)
	_ = m.SetValue(3, 0, pointW)

	return m
}

// ToPoint returns a Point representation of the passed Matrix.
func ToPoint(m matrix.Matrix) *Point {
	// TODO: error check on 4x1 matrix

	x, _ := m.GetValue(0, 0)
	y, _ := m.GetValue(1, 0)
	z, _ := m.GetValue(2, 0)
	return NewPoint(x, y, z)
}
