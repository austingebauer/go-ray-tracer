// Package vector represents a vector in a left-handed 3D coordinate system.
package vector

import (
	"math"

	"github.com/austingebauer/go-ray-tracer/maths"
)

// Vector represents a vector in a left-handed 3D coordinate system
type Vector struct {
	// X, Y, and Z represent components in a left-handed 3D coordinate system
	X, Y, Z float64
}

// NewVector returns a new Vector that has the passed X, Y, and Z values.
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

// Equals returns true if the passed Vector is equal to this Vector.
// Two Vectors are equal if their X, Y, Z components are equal.
func (vec *Vector) Equals(vecQ *Vector) bool {
	return maths.Float64Equals(vec.X, vecQ.X, maths.Epsilon) &&
		maths.Float64Equals(vec.Y, vecQ.Y, maths.Epsilon) &&
		maths.Float64Equals(vec.Z, vecQ.Z, maths.Epsilon)
}

// Magnitude computes and returns the length of this Vector.
// The length is calculated using Pythagoras' theorem.
func (vec *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2) +
		math.Pow(vec.Y, 2) +
		math.Pow(vec.Z, 2))
}

// Negate multiplies each of this Vector's components by -1.
func (vec *Vector) Negate() *Vector {
	return vec.Scale(-1)
}

// Scale multiplies each of this Vector's components by the passed scalar value.
func (vec *Vector) Scale(scalar float64) *Vector {
	vec.X = vec.X * scalar
	vec.Y = vec.Y * scalar
	vec.Z = vec.Z * scalar
	return vec
}

// Scale returns a new Vector that is the result of scaling the passed Vector
// by the passed scalar value.
func Scale(vec Vector, scalar float64) *Vector {
	return NewVector(vec.X*scalar, vec.Y*scalar, vec.Z*scalar)
}

// Normalize normalizes this Vector by converting it to a unit vector.
func (vec *Vector) Normalize() *Vector {
	mag := vec.Magnitude()
	vec.X = vec.X / mag
	vec.Y = vec.Y / mag
	vec.Z = vec.Z / mag
	return vec
}

// Normalize returns a new Vector that is the result of normalizing the passed Vector.
func Normalize(vec Vector) *Vector {
	vec.Normalize()
	return &Vector{
		X: vec.X,
		Y: vec.Y,
		Z: vec.Z,
	}
}

// DotProduct computes and returns the dot product of passed Vectors.
func DotProduct(vec1, vec2 Vector) float64 {
	// The dot product yields the cosine of the angle between the passed Vectors.
	// If it is equal to 0, the vectors are perpendicular to each other.
	// If it is equal to 1, the vectors are parallel to each other.
	return vec1.X*vec2.X +
		vec1.Y*vec2.Y +
		vec1.Z*vec2.Z
}

// CrossProduct computes and returns a new Vector that is the
// cross product of the passed Vectors.
func CrossProduct(vec1, vec2 Vector) Vector {
	return Vector{
		X: (vec1.Y * vec2.Z) - (vec1.Z * vec2.Y),
		Y: (vec1.Z * vec2.X) - (vec1.X * vec2.Z),
		Z: (vec1.X * vec2.Y) - (vec1.Y * vec2.X),
	}
}

// Add modifies each component of this Vector by setting each of them
// to the sum of the components in this Vector and the passed Vector.
func (vec *Vector) Add(vec2 Vector) *Vector {
	vec.X = vec.X + vec2.X
	vec.Y = vec.Y + vec2.Y
	vec.Z = vec.Z + vec2.Z
	return vec
}

// Add returns a new Vector with components equal to the sum
// of the corresponding components in the passed Vectors.
func Add(vec1, vec2 Vector) Vector {
	return Vector{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
		Z: vec1.Z + vec2.Z,
	}
}

// Subtract modifies each component of this Vector by setting each of them
// to the difference of the components in this Vector and the passed Vector.
func (vec *Vector) Subtract(vec2 Vector) *Vector {
	vec.X = vec.X - vec2.X
	vec.Y = vec.Y - vec2.Y
	vec.Z = vec.Z - vec2.Z
	return vec
}

// Subtract returns a new Vector with components equal to the difference
// of the corresponding components in the passed Vectors.
func Subtract(vec1, vec2 Vector) *Vector {
	return NewVector(vec1.X-vec2.X, vec1.Y-vec2.Y, vec1.Z-vec2.Z)
}

// Reflect returns a new Vector that is the results of reflecting
// the passed in Vector around the passed normal Vector.
func Reflect(in, normal Vector) *Vector {
	dp := DotProduct(in, normal)
	nScaled := normal.Scale(2).Scale(dp)
	return Subtract(in, *nScaled)
}
