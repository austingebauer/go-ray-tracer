// package tuple represents a point or a vector in a left-handed coordinate system.
// The package provides functions used to make calculations on points and vectors in a 3D space.
package tuple

import (
	"errors"

	"github.com/austingebauer/go-ray-tracer/utils"
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
	tpl, _ := newTuple(x, y, z, point)
	return tpl
}

// NewVector returns a new tuple that has the passed x, y, and z values.
func NewVector(x, y, z float64) *tuple {
	tpl, _ := newTuple(x, y, z, vector)
	return tpl
}

// Equals returns true if the passed tuple is equal to this tuple.
// Two Tuples are equal if their x, y, z, and w members are equal.
func Equals(tpl1, tpl2 *tuple) bool {
	return utils.Float64Equals(tpl1.x, tpl2.x, utils.Epsilon) &&
		utils.Float64Equals(tpl1.y, tpl2.y, utils.Epsilon) &&
		utils.Float64Equals(tpl1.z, tpl2.z, utils.Epsilon) &&
		utils.Float64Equals(tpl1.w, tpl2.w, utils.Epsilon)
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

// Negate returns a new tuple by negating the passed Tuple.
// Negating the passed tuple happens by multiplying each of it's components by -1.
func Negate(tpl *tuple) (*tuple, error) {
	return Scale(tpl, -1)
}

// Scale returns a new tuple that is the result of multiplying each
// component of the passed tuple by the passed scalar.
func Scale(tpl *tuple, scalar float64) (*tuple, error) {
	// The w component does not get multipled because it's a flag to signify a point or vector
	return newTuple(tpl.x*scalar, tpl.y*scalar, tpl.z*scalar, tpl.w)
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
