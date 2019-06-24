package matrix

// Matrix is a two-dimensional slice of floating point numbers.
type Matrix [][]float64

// NewMatrix returns a new two-dimensional matrix having the passed rows and columns.
func NewMatrix(rows, cols uint) Matrix {
	m := make([][]float64, rows)
	for row := range m {
		m[row] = make([]float64, cols)
	}
	return m
}
