// Package utils contains utility functions.
package math_utils

import "math"

const (
	Epsilon float64 = 0.00001
)

// Float64Equals returns true if the passed float64 values are equal.
// The two passed float64 values are considered equal if the absolute
// values of their difference is less than the passed epsilon.
func Float64Equals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
