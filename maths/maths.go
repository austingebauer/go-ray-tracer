// Package maths contains math utility functions.
package maths

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

// ToRadians returns the conversion of the passed degree value into radians.
func ToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// ToDegrees returns the conversion of the passed radian value into degrees.
func ToDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
