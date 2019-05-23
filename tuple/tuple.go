package tuple

import (
	"errors"

	"github.com/austingebauer/go-ray-tracer/math"
)

const (
	VectorW float64 = 0.0
	PointW  float64 = 1.0
)

type Tuple struct {
	// X, Y, and Z represent points in a left-handed coordinate system
	X, Y, Z float64

	// W indicates whether this Tuple represents a Point or a Vector
	// If the value of W is 0.0, then it represents a Vector
	// If the value of W is 1.0, then it represents a Point
	W float64
}

// NewTuple returns a new Tuple that has the passed x, y, z, and w values
func NewTuple(x, y, z, w float64) (*Tuple, error) {
	if w != 0.0 && w != 1.0 {
		return nil, errors.New("W must be either 0.0 or 1.0")
	}

	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}, nil
}

// TODO: Consider Equals(tpl1, tpl2 *Tuple) signture instead
// Equals returns true if the passed Tuple is equal to this Tuple.
// Two Tuples are equal if their X, Y, Z, and W members are equal.
func (tpl1 *Tuple) Equals(tpl2 *Tuple) bool {
	return math.Float64Equals(tpl1.X, tpl2.X, math.Epsilon) &&
		math.Float64Equals(tpl1.Y, tpl2.Y, math.Epsilon) &&
		math.Float64Equals(tpl1.Z, tpl2.Z, math.Epsilon) &&
		math.Float64Equals(tpl1.W, tpl2.W, math.Epsilon)
}
