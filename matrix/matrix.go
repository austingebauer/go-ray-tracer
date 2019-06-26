package matrix

import "github.com/austingebauer/go-ray-tracer/utils"

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

// Equals returns true if this Matrix has identical rows, columns, and elements as the passed Matrix.
func (m *Matrix) Equals(mQ *Matrix) bool {
	if m.rows != mQ.rows || m.cols != mQ.cols {
		return false
	}

	for r := 0; r < int(m.rows); r++ {
		for c := 0; c < int(m.cols); c++ {
			if !utils.Float64Equals(m.data[r][c], mQ.data[r][c], utils.Epsilon) {
				return false
			}
		}
	}

	return true
}
