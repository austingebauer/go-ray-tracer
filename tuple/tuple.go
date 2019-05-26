// package tuple represents a point or a vector in a left-handed coordinate system.
// The package provides functions used to make calculations on points and vectors in a 3D space.
package tuple

import (
	"errors"
	"log"

	"github.com/austingebauer/go-ray-tracer/math"
)

const (
	vector float64 = 0.0
	point  float64 = 1.0
)

// tuple represents a point or a vector in a left-handed coordinate system
type tuple struct {
	// x, y, and z represent components in a left-handed coordinate system
	x, y, z float64

	// w indicates whether this tuple represents a point or a vector
	// If the value of w is 0.0, then it represents a vector
	// If the value of w is 1.0, then it represents a point
	w float64
}

// NewPoint returns a new tuple that has the passed x, y, and z values.
func NewPoint(x, y, z float64) *tuple {
	tpl, err := newTuple(x, y, z, point)
	if err != nil {
		log.Fatal(err)
	}
	return tpl
}

// NewVector returns a new tuple that has the passed x, y, and z values.
func NewVector(x, y, z float64) *tuple {
	tpl, err := newTuple(x, y, z, vector)
	if err != nil {
		log.Fatal(err)
	}
	return tpl
}

// Equals returns true if the passed tuple is equal to this tuple.
// Two Tuples are equal if their x, y, z, and w members are equal.
func Equals(tpl1, tpl2 *tuple) bool {
	return math.Float64Equals(tpl1.x, tpl2.x, math.Epsilon) &&
		math.Float64Equals(tpl1.y, tpl2.y, math.Epsilon) &&
		math.Float64Equals(tpl1.z, tpl2.z, math.Epsilon) &&
		math.Float64Equals(tpl1.w, tpl2.w, math.Epsilon)
}

// Add returns a new tuple by adding the corresponding components in each of the passed Tuples.
func Add(tpl1, tpl2 *tuple) (*tuple, error) {
	if tpl1.w+tpl2.w == (point + point) {
		return nil, errors.New("error: cannot add two point tuples")
	}

	return newTuple(tpl1.x+tpl2.x, tpl1.y+tpl2.y, tpl1.z+tpl2.z, tpl1.w+tpl2.w)
}

// Subtract returns a new tuple by subtracting the corresponding components in each of the passed Tuples.
func Subtract(tpl1, tpl2 *tuple) (*tuple, error) {
	if tpl1.w-tpl2.w == (vector - point) {
		return nil, errors.New("error: cannot subtract a point from a vector")
	}

	return newTuple(tpl1.x-tpl2.x, tpl1.y-tpl2.y, tpl1.z-tpl2.z, tpl1.w-tpl2.w)
}

// Negate negates the passed Tuple by multiplying each of it's components by -1.
func Negate(tpl *tuple) (*tuple, error) {
	// The w component does not get negated because it's a flag to signify a point or vector
	return newTuple(tpl.x*-1, tpl.y*-1, tpl.z*-1, tpl.w)
}

// newTuple returns a new tuple that has the passed x, y, z, and w values.
// This function is private in order to make the public interface for
// a tuple explicitly a point or a vector.
func newTuple(x, y, z, w float64) (*tuple, error) {
	if w != vector && w != point {
		return nil, errors.New("error: w must be either 0.0 or 1.0")
	}

	return &tuple{
		x: x,
		y: y,
		z: z,
		w: w,
	}, nil
}
