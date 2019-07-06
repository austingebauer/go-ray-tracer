package matrix

import (
	"errors"

	"github.com/austingebauer/go-ray-tracer/math_utils"
)

// Identity is a 4x4 identity matrix.
var Identity = Matrix{
	rows: 4,
	cols: 4,
	data: [][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	},
}

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

// NewIdentityMatrix returns a new identity Matrix having row and column length equal to the passed size.
func NewIdentityMatrix(size uint) *Matrix {
	m := NewMatrix(size, size)
	for i := 0; i < int(m.rows); i++ {
		m.data[i][i] = 1
	}

	return m
}

// Equals returns true if this Matrix has identical rows, columns, and elements as the passed Matrix.
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

// Determinant calculate and returns the determinant of the passed 2x2 Matrix.
// If the passed matrix is not 2x2, then an error is returned.
func Determinant2x2(m Matrix) (float64, error) {
	/*
			m = | a b |
		        | c d |

			det(m) = ad - bc
	*/

	if m.rows != 2 || m.cols != 2 {
		return 0, errors.New("matrix must have row and column length of 2")
	}

	return (m.data[0][0] * m.data[1][1]) - (m.data[0][1] * m.data[1][0]), nil
}

// Submatrix returns a new Matrix that is the result of removing
// the passed row and column index from the passed Matrix.
// If the passed row or col are not in bounds of the passed Matrix,
// then an error is returned.
func Submatrix(m Matrix, row, col uint) (*Matrix, error) {
	if row < 0 || row >= m.rows {
		return nil, errors.New("row is out of bounds of the passed matrix")
	}

	if col < 0 || col >= m.cols {
		return nil, errors.New("col is out of bounds of the passed matrix")
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

// Minor3x3 returns the determinant of the submatrix.
// If the passed matrix is not 3x3, then an error is returned.
func Minor3x3(m Matrix, row, col uint) (float64, error) {
	if m.rows != 3 || m.cols != 3 {
		return 0, errors.New("matrix must have row and column length of 3")
	}

	subM, err := Submatrix(m, row, col)
	if err != nil {
		return 0, err
	}

	return Determinant2x2(*subM)
}
