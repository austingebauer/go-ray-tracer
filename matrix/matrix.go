package matrix

import (
	"errors"
	"math"

	"github.com/austingebauer/go-ray-tracer/math_utils"
)

// Matrix represents an n-dimensional grid of floating point numbers.
type Matrix struct {
	rows, cols uint
	data       [][]float64
}

// NewMatrix returns a new Matrix having the passed row and column lengths.
func NewMatrix(rows, cols uint) *Matrix {
	m := Matrix{
		rows: rows,
		cols: cols,
	}

	d := make([][]float64, m.rows)
	for row := range d {
		d[row] = make([]float64, m.cols)
	}
	m.data = d

	return &m
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
	err := CheckInBounds(*m, row, col)
	if err != nil {
		return err
	}

	m.data[row][col] = val
	return nil
}

// GetValue sets the passed value at the passed row and column in the Matrix.
func (m *Matrix) GetValue(row, col uint) (float64, error) {
	err := CheckInBounds(*m, row, col)
	if err != nil {
		return 0, err
	}

	return m.data[row][col], nil
}

// NewIdentityMatrix returns a new identity Matrix having row and column
// length equal to the passed size.
func NewIdentityMatrix(size uint) *Matrix {
	m := NewMatrix(size, size)
	for i := 0; i < int(m.rows); i++ {
		m.data[i][i] = 1
	}

	return m
}

// Equals returns true if this Matrix has identical rows, columns,
// and elements as the passed Matrix.
func (m *Matrix) Equals(m1 *Matrix) bool {
	if m.rows != m1.rows || m.cols != m1.cols {
		return false
	}

	for r := 0; r < int(m.rows); r++ {
		for c := 0; c < int(m.cols); c++ {
			if !math_utils.Float64Equals(m.data[r][c], m1.data[r][c], math_utils.Epsilon) {
				return false
			}
		}
	}

	return true
}

// Multiply returns a new Matrix that is the result of multiplying the passed matrices.
// If the column length in m1 is not equal to the row length in m2, an error is returned.
func Multiply(m1, m2 Matrix) (*Matrix, error) {
	// To multiply an m×n matrix by an n×p matrix, the n's must be the same.
	if m1.cols != m2.rows {
		return nil, errors.New("column length of m1 must be equal to the row length of m2")
	}

	// The result is an m×p matrix.
	multM := NewMatrix(m1.rows, m2.cols)

	// Multiply the two matrices
	for m := 0; m < int(multM.rows); m++ {
		for p := 0; p < int(multM.cols); p++ {

			// Compute the dot product over m1 columns and m2 rows for range 0 < n
			var dotProduct float64
			for n := 0; n < int(m1.cols); n++ {
				dotProduct += m1.data[m][n] * m2.data[n][p]
			}

			// Store dot product in m and p
			multM.data[m][p] = dotProduct
		}
	}
	return multM, nil
}

// Transpose returns a new Matrix that is the result of transposing the passed Matrix.
// Transposing a Matrix turns the nth row into the nth column in the resulting Matrix.
func Transpose(m Matrix) *Matrix {
	transM := NewMatrix(m.cols, m.rows)

	for r := 0; r < int(m.rows); r++ {
		for c := 0; c < int(m.cols); c++ {
			transM.data[c][r] = m.data[r][c]
		}
	}

	return transM
}

// Determinant calculate and returns the determinant of the passed Matrix.
func Determinant(m Matrix) float64 {
	if m.rows == 1 && m.cols == 1 {
		return m.data[0][0]
	}

	if m.rows == 2 && m.cols == 2 {
		return (m.data[0][0] * m.data[1][1]) - (m.data[0][1] * m.data[1][0])
	}

	// for each column in the selected row
	var det float64 = 0
	row := 0
	for col := 0; col < int(m.cols); col++ {
		cofactor, _ := Cofactor(m, uint(row), uint(col))
		det = det + (m.data[row][col] * cofactor)
	}

	return det
}

// Submatrix returns a new Matrix that is the result of removing
// the passed row and column index from the passed Matrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Submatrix(m Matrix, row, col uint) (*Matrix, error) {
	err := CheckInBounds(m, row, col)
	if err != nil {
		return nil, err
	}

	subM := NewMatrix(m.rows-1, m.cols-1)
	for r := 0; r < int(m.rows); r++ {
		for c := 0; c < int(m.cols); c++ {
			if r == int(row) || c == int(col) {
				continue
			}

			// hold current value of r and c
			rowPlacement := r
			colPlacement := c

			// if current row or col are beyond the passed row and col to remove,
			// then their placement in the submatrix will be minus one of their
			// current location.
			if rowPlacement > int(row) {
				rowPlacement--
			}

			if colPlacement > int(col) {
				colPlacement--
			}

			subM.data[rowPlacement][colPlacement] = m.data[r][c]
		}
	}

	return subM, nil
}

// Minor returns the determinant of the submatrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Minor(m Matrix, row, col uint) (float64, error) {
	subM, err := Submatrix(m, row, col)
	if err != nil {
		return 0, err
	}

	return Determinant(*subM), nil
}

// Cofactor returns the cofactor of the submatrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Cofactor(m Matrix, row, col uint) (float64, error) {
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
func IsInvertible(m Matrix) bool {
	return Determinant(m) != 0
}

// Inverse returns the inverse of the passed Matrix.
func Inverse(m Matrix) (*Matrix, error) {
	if !IsInvertible(m) {
		return nil, errors.New("the passed matrix is not invertible")
	}

	mInverted := NewMatrix(m.rows, m.cols)

	// Calculate the determinant of m
	determinantM := Determinant(m)

	// Place the cofactor of each element divided by the determinant into a transposition of m.
	for row := 0; row < int(m.rows); row++ {
		for col := 0; col < int(m.cols); col++ {
			c, err := Cofactor(m, uint(row), uint(col))
			if err != nil {
				return nil, err
			}

			// note that col and row are reversed in the placement to accomplish transposing
			mInverted.data[col][row] = c / determinantM
		}
	}

	return mInverted, nil
}

// CheckInBounds returns an error if either the row or column values
// are out of bounds of the passed Matrix.
func CheckInBounds(m Matrix, row, col uint) error {
	if row < 0 || row >= m.rows {
		return errors.New("row is out of bounds of the passed matrix")
	}

	if col < 0 || col >= m.cols {
		return errors.New("col is out of bounds of the passed matrix")
	}

	return nil
}

// Translation returns a 4x4 translation Matrix.
//
// The translation Matrix returned has the form:
//   | 1 0 0 x |
//   | 0 1 0 y |
//   | 0 0 1 z |
//   | 0 0 0 1 |
func Translation(x, y, z float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.data[0][3] = x
	m.data[1][3] = y
	m.data[2][3] = z
	return m
}

// Scaling returns a 4x4 scaling Matrix.
//
// The scaling Matrix returned has the form:
//   | x 0 0 0 |
//   | 0 y 0 0 |
//   | 0 0 z 0 |
//   | 0 0 0 1 |
func Scaling(x, y, z float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.data[0][0] = x
	m.data[1][1] = y
	m.data[2][2] = z
	return m
}

// RotationX returns a 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the X axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | 1 0      0       0 |
//   | 0 cos(r) -sin(r) 0 |
//   | 0 sin(r) cos(r)  0 |
//   | 0 0      0       1 |
func RotationX(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.data[1][1] = math.Cos(radians)
	m.data[1][2] = -1 * math.Sin(radians)
	m.data[2][1] = math.Sin(radians)
	m.data[2][2] = math.Cos(radians)
	return m
}

// RotationY returns a 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the Y axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | cos(r)  0 sin(r) 0 |
//   | 0       1 0      0 |
//   | -sin(r) 0 cos(r) 0 |
//   | 0       0 0      1 |
func RotationY(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.data[0][0] = math.Cos(radians)
	m.data[0][2] = math.Sin(radians)
	m.data[2][0] = -1 * math.Sin(radians)
	m.data[2][2] = math.Cos(radians)
	return m
}

// RotationZ returns a 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the Z axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | cos(r) -sin(r) 0 0 |
//   | sin(r) cos(r)  0 0 |
//   | 0      0       1 0 |
//   | 0      0       0 1 |
func RotationZ(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.data[0][0] = math.Cos(radians)
	m.data[0][1] = -1 * math.Sin(radians)
	m.data[1][0] = math.Sin(radians)
	m.data[1][1] = math.Cos(radians)
	return m
}
