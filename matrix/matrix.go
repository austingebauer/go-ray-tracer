// Package matrix represents a matrix of a given dimension.
// It provides commonly used matrix operations.
package matrix

import (
	"errors"
	"math"

	"github.com/austingebauer/go-ray-tracer/maths"
)

// Matrix represents an n-dimensional grid of floating point numbers.
type Matrix struct {
	rows, cols uint
	data       []float64
}

// NewMatrix returns a new Matrix having the passed row and column lengths.
func NewMatrix(rows, cols uint) *Matrix {
	m := Matrix{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}

	return &m
}

// NewIdentityMatrix returns a new identity Matrix having
// row and column length equal to the passed size.
func NewIdentityMatrix(size uint) *Matrix {
	m := NewMatrix(size, size)
	for i := uint(0); i < m.rows; i++ {
		m.setValue(i, i, 1)
	}

	return m
}

// NewTranslationMatrix returns a new 4x4 translation Matrix.
//
// The translation Matrix returned has the form:
//   | 1 0 0 x |
//   | 0 1 0 y |
//   | 0 0 1 z |
//   | 0 0 0 1 |
func NewTranslationMatrix(x, y, z float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 3, x)
	m.setValue(1, 3, y)
	m.setValue(2, 3, z)
	return m
}

// NewScalingMatrix returns a new 4x4 scaling Matrix.
//
// The scaling Matrix returned has the form:
//   | x 0 0 0 |
//   | 0 y 0 0 |
//   | 0 0 z 0 |
//   | 0 0 0 1 |
func NewScalingMatrix(x, y, z float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 0, x)
	m.setValue(1, 1, y)
	m.setValue(2, 2, z)
	return m
}

// NewXRotationMatrix returns a new 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the X axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | 1 0      0       0 |
//   | 0 cos(r) -sin(r) 0 |
//   | 0 sin(r) cos(r)  0 |
//   | 0 0      0       1 |
func NewXRotationMatrix(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(1, 1, math.Cos(radians))
	m.setValue(1, 2, -1*math.Sin(radians))
	m.setValue(2, 1, math.Sin(radians))
	m.setValue(2, 2, math.Cos(radians))
	return m
}

// NewYRotationMatrix returns a new 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the Y axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | cos(r)  0 sin(r) 0 |
//   | 0       1 0      0 |
//   | -sin(r) 0 cos(r) 0 |
//   | 0       0 0      1 |
func NewYRotationMatrix(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 0, math.Cos(radians))
	m.setValue(0, 2, math.Sin(radians))
	m.setValue(2, 0, -1*math.Sin(radians))
	m.setValue(2, 2, math.Cos(radians))
	return m
}

// NewZRotationMatrix returns a new 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the z-axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | cos(r) -sin(r) 0 0 |
//   | sin(r) cos(r)  0 0 |
//   | 0      0       1 0 |
//   | 0      0       0 1 |
func NewZRotationMatrix(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 0, math.Cos(radians))
	m.setValue(0, 1, -1*math.Sin(radians))
	m.setValue(1, 0, math.Sin(radians))
	m.setValue(1, 1, math.Cos(radians))
	return m
}

// NewShearingMatrix returns a new 4x4 shearing Matrix that can be used for a
// shear transformation of a Point.
//
// The shearing Matrix returned has the form:
//   | 1  xy xz 0 |
//   | yx 1  yz 0 |
//   | zx zy 1  0 |
//   | 0  0  0  1 |
func NewShearingMatrix(xy, xz, yx, yz, zx, zy float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 1, xy)
	m.setValue(0, 2, xz)
	m.setValue(1, 0, yx)
	m.setValue(1, 2, yz)
	m.setValue(2, 0, zx)
	m.setValue(2, 1, zy)
	return m
}

// GetRows returns the number of rows that the Matrix has.
func (m *Matrix) GetRows() uint {
	return m.rows
}

// GetCols returns the number of columns that the Matrix has.
func (m *Matrix) GetCols() uint {
	return m.cols
}

// SetValue sets the passed value at the passed row and column in the Matrix.
func (m *Matrix) SetValue(row, col uint, val float64) error {
	err := CheckInBounds(m, row, col)
	if err != nil {
		return err
	}

	m.setValue(row, col, val)
	return nil
}

// setValue sets the passed value at the passed row and column in the Matrix
// without checking for the values being out-of-bounds.
func (m *Matrix) setValue(row, col uint, val float64) {
	m.data[row*m.cols+col] = val
}

// GetValue sets the passed value at the passed row and column in the Matrix.
func (m *Matrix) GetValue(row, col uint) (float64, error) {
	err := CheckInBounds(m, row, col)
	if err != nil {
		return 0, err
	}

	return m.getValue(row, col), nil
}

// setValue sets the passed value at the passed row and column in the Matrix
// without checking for the values being out-of-bounds.
func (m *Matrix) getValue(row, col uint) float64 {
	return m.data[row*m.cols+col]
}

// Equals returns true if this Matrix has identical rows, columns,
// and elements as the passed Matrix.
func (m *Matrix) Equals(m1 *Matrix) bool {
	if m.rows != m1.rows || m.cols != m1.cols {
		return false
	}

	for r := uint(0); r < m.rows; r++ {
		for c := uint(0); c < m.cols; c++ {
			if !maths.Float64Equals(m.getValue(r, c), m1.getValue(r, c), maths.Epsilon) {
				return false
			}
		}
	}

	return true
}

// Translate translates this Matrix by the passed x, y, and z values.
func (m *Matrix) Translate(x, y, z float64) *Matrix {
	transform := NewTranslationMatrix(x, y, z)
	mRes, _ := Multiply(transform, m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// Scale scales this Matrix by the passed x, y, and z values.
func (m *Matrix) Scale(x, y, z float64) *Matrix {
	transform := NewScalingMatrix(x, y, z)
	mRes, _ := Multiply(transform, m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// RotateX rotates this Matrix around the x-axis by the passed number of radians.
// Rotation happens clockwise when looking down the positive x-axis towards the negative x-axis.
func (m *Matrix) RotateX(radians float64) *Matrix {
	transform := NewXRotationMatrix(radians)
	mRes, _ := Multiply(transform, m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// RotateY rotates this Matrix around the y-axis by the passed number of radians.
// Rotation happens clockwise when looking down the positive y-axis towards the negative y-axis.
func (m *Matrix) RotateY(radians float64) *Matrix {
	transform := NewYRotationMatrix(radians)
	mRes, _ := Multiply(transform, m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// RotateZ rotates this Matrix around the z-axis by the passed number of radians.
// Rotation happens clockwise when looking down the positive z-axis towards the negative z-axis.
func (m *Matrix) RotateZ(radians float64) *Matrix {
	transform := NewZRotationMatrix(radians)
	mRes, _ := Multiply(transform, m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// Shear shears or skews this Matrix in on coordinate relative to another coordinate.
// For example, the parameter xy represents how much to shear x relative to y.
func (m *Matrix) Shear(xy, xz, yx, yz, zx, zy float64) *Matrix {
	transform := NewShearingMatrix(xy, xz, yx, yz, zx, zy)
	mRes, _ := Multiply(transform, m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// Multiply returns a new Matrix that is the result of multiplying the passed matrices.
// If the column length in m1 is not equal to the row length in m2, an error is returned.
func Multiply(m1, m2 *Matrix) (*Matrix, error) {
	// To multiply an m×n matrix by an n×p matrix, the n's must be the same.
	if m1.cols != m2.rows {
		return nil, errors.New("column length of m1 must be equal to the row length of m2")
	}

	// The result is an m×p matrix.
	multM := NewMatrix(m1.rows, m2.cols)

	// Multiply the two matrices
	for m := uint(0); m < multM.rows; m++ {
		for p := uint(0); p < multM.cols; p++ {

			// Compute the dot product over m1 columns and m2 rows for range 0 < n
			var dotProduct float64
			for n := uint(0); n < m1.cols; n++ {
				dotProduct += m1.getValue(m, n) * m2.getValue(n, p)
			}

			// Store dot product in m and p
			multM.setValue(m, p, dotProduct)
		}
	}
	return multM, nil
}

// Transpose returns a new Matrix that is the result of transposing the passed Matrix.
// Transposing a Matrix turns the nth row into the nth column in the resulting Matrix.
func Transpose(m Matrix) *Matrix {
	transM := NewMatrix(m.cols, m.rows)

	for r := uint(0); r < m.rows; r++ {
		for c := uint(0); c < m.cols; c++ {
			transM.setValue(c, r, m.getValue(r, c))
		}
	}

	return transM
}

// Determinant calculate and returns the determinant of the passed Matrix.
// If the passed Matrix is not a square matrix, then an error is returned.
func Determinant(m *Matrix) (float64, error) {
	if m.rows != m.cols {
		return 0, errors.New("m must be a square matrix with equal row and column lengths")
	}

	if m.rows == 0 && m.cols == 0 {
		return 1, nil
	}

	if m.rows == 1 && m.cols == 1 {
		return m.getValue(0, 0), nil
	}

	if m.rows == 2 && m.cols == 2 {
		return (m.getValue(0, 0) * m.getValue(1, 1)) -
			(m.getValue(0, 1) * m.getValue(1, 0)), nil
	}

	// for each column in the selected row
	var det float64
	row := uint(0)
	for col := uint(0); col < m.cols; col++ {
		cofactor, _ := Cofactor(m, uint(row), uint(col))
		det = det + (m.getValue(row, col) * cofactor)
	}

	return det, nil
}

// Submatrix returns a new Matrix that is the result of removing
// the passed row and column index from the passed Matrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Submatrix(m *Matrix, row, col uint) (*Matrix, error) {
	err := CheckInBounds(m, row, col)
	if err != nil {
		return nil, err
	}

	subM := NewMatrix(m.rows-1, m.cols-1)
	for r := uint(0); r < m.rows; r++ {
		for c := uint(0); c < m.cols; c++ {
			if r == row || c == col {
				continue
			}

			// hold current value of r and c
			rowPlacement := r
			colPlacement := c

			// if current row or col are beyond the passed row and col to remove,
			// then their placement in the submatrix will be minus one of their
			// current location.
			if rowPlacement > row {
				rowPlacement--
			}

			if colPlacement > col {
				colPlacement--
			}

			subM.setValue(rowPlacement, colPlacement, m.getValue(r, c))
		}
	}

	return subM, nil
}

// Minor returns the determinant of the submatrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Minor(m *Matrix, row, col uint) (float64, error) {
	subM, err := Submatrix(m, row, col)
	if err != nil {
		return 0, err
	}

	det, err := Determinant(subM)
	if err != nil {
		return 0, err
	}

	return det, nil
}

// Cofactor returns the cofactor of the submatrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Cofactor(m *Matrix, row, col uint) (float64, error) {
	minor, err := Minor(m, row, col)
	if err != nil {
		return 0, err
	}

	// If row plus column is an odd number, then the cofactor is the negated minor.
	// Otherwise, the cofactor is the unmodified minor.
	shouldNegateMinor := (row+col)%2 == 1
	if shouldNegateMinor {
		minor = minor * -1
	}

	return minor, nil
}

// IsInvertible returns true if the passed Matrix is invertible.
// The passed Matrix is invertible if it's determinant is equal to 0.
func IsInvertible(m *Matrix) bool {
	det, err := Determinant(m)
	if err != nil {
		return false
	}
	return det != 0
}

// Inverse returns the inverse of the passed Matrix.
func Inverse(m *Matrix) (*Matrix, error) {
	if !IsInvertible(m) {
		return nil, errors.New("the passed matrix is not invertible")
	}

	mInverted := NewMatrix(m.rows, m.cols)

	// Calculate the determinant of m
	determinantM, _ := Determinant(m)

	// Place the cofactor of each element divided by the determinant into a transposition of m.
	for row := uint(0); row < m.rows; row++ {
		for col := uint(0); col < m.cols; col++ {
			c, err := Cofactor(m, uint(row), uint(col))
			if err != nil {
				return nil, err
			}

			// Note that col and row are reversed in the placement to accomplish transposing
			mInverted.setValue(col, row, c/determinantM)
		}
	}

	return mInverted, nil
}

// CheckInBounds returns an error if either the row or column values
// are out of bounds of the passed Matrix.
func CheckInBounds(m *Matrix, row, col uint) error {
	if row < 0 || row >= m.rows {
		return errors.New("row is out of bounds of the passed matrix")
	}

	if col < 0 || col >= m.cols {
		return errors.New("col is out of bounds of the passed matrix")
	}

	return nil
}
