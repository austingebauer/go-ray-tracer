// Package point represents a point in a left-handed 3D coordinate system.
package point

import (
	"github.com/austingebauer/go-ray-tracer/maths"
	"github.com/austingebauer/go-ray-tracer/vector"
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
	return maths.Float64Equals(pt.X, ptQ.X, maths.Epsilon) &&
		maths.Float64Equals(pt.Y, ptQ.Y, maths.Epsilon) &&
		maths.Float64Equals(pt.Z, ptQ.Z, maths.Epsilon)
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
func (pt *Point) Add(vec *vector.Vector) *Point {
	pt.X = pt.X + vec.X
	pt.Y = pt.Y + vec.Y
	pt.Z = pt.Z + vec.Z
	return pt
}

// Add returns a new Point with components equal to the sum
// of the corresponding components in the passed Point and Vector.
func Add(pt *Point, vec *vector.Vector) *Point {
	return NewPoint(pt.X+vec.X, pt.Y+vec.Y, pt.Z+vec.Z)
}

// Subtract modifies each component of this Point by setting each of them
// to the difference of the components in this Point and the passed Vector.
func (pt *Point) Subtract(vec *vector.Vector) *Point {
	pt.X = pt.X - vec.X
	pt.Y = pt.Y - vec.Y
	pt.Z = pt.Z - vec.Z
	return pt
}

// Subtract returns a new Vector with components equal to the
// difference of the corresponding components in the passed Points.
func Subtract(pt1 Point, pt2 Point) *vector.Vector {
	return &vector.Vector{
		X: pt1.X - pt2.X,
		Y: pt1.Y - pt2.Y,
		Z: pt1.Z - pt2.Z,
	}
}
