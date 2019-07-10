package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/austingebauer/go-ray-tracer/math_utils"
)

func TestNewMatrix(t *testing.T) {
	type args struct {
		rows uint
		cols uint
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "create new 2x2 matrix",
			args: args{
				rows: 2,
				cols: 2,
			},
			want: &Matrix{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{0.0, 0.0},
					{0.0, 0.0},
				},
			},
		},
		{
			name: "create new 0x0 matrix",
			args: args{
				rows: 0,
				cols: 0,
			},
			want: &Matrix{
				rows: 0,
				cols: 0,
				data: [][]float64{},
			},
		},
		{
			name: "create new 1x0 matrix",
			args: args{
				rows: 1,
				cols: 0,
			},
			want: &Matrix{
				rows: 1,
				cols: 0,
				data: [][]float64{
					{},
				},
			},
		},
		{
			name: "create new 0x1 matrix",
			args: args{
				rows: 0,
				cols: 1,
			},
			want: &Matrix{
				rows: 0,
				cols: 1,
				data: [][]float64{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewMatrix(tt.args.rows, tt.args.cols))
		})
	}
}

func TestMatrix_Equals(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	type args struct {
		m1 *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "2x2 matrix equals 2x2 matrix",
			fields: fields{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{1.0, 2.0},
					{3.0, -4.0},
				},
			},
			args: args{
				m1: &Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{1.0, 2.0},
						{3.0, -4.0},
					},
				},
			},
			want: true,
		},
		{
			name: "2x2 matrix does not equal 2x2 matrix",
			fields: fields{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{1.0, 2.0},
					{3.0, -4.0},
				},
			},
			args: args{
				m1: &Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{-1.0, 2.0},
						{3.0, 4.0},
					},
				},
			},
			want: false,
		},
		{
			name: "2x3 matrix does not equal 1x1 matrix",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{1.0, 2.0, 3.0},
					{3.0, -4.0, -1.3},
				},
			},
			args: args{
				m1: &Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{-1.0},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			assert.Equal(t, tt.want, m.Equals(tt.args.m1))
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		m1 Matrix
		m2 Matrix
	}
	tests := []struct {
		name      string
		args      args
		want      *Matrix
		wantError bool
	}{
		{
			name: "multiply two 4x4 matrices",
			args: args{
				m1: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 8, 7, 6},
						{5, 4, 3, 2},
					},
				},
				m2: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-2, 1, 2, 3},
						{3, 2, 1, -1},
						{4, 3, 6, 5},
						{1, 2, 7, 8},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{20, 22, 50, 48},
					{44, 54, 114, 108},
					{40, 58, 110, 102},
					{16, 26, 46, 42},
				},
			},
			wantError: false,
		},
		{
			name: "multiply 1x3 and 3x2 matrices",
			args: args{
				m1: Matrix{
					rows: 1,
					cols: 3,
					data: [][]float64{
						{1, 2, 3},
					},
				},
				m2: Matrix{
					rows: 3,
					cols: 2,
					data: [][]float64{
						{4, 3},
						{2, 4},
						{1, 5},
					},
				},
			},
			want: &Matrix{
				rows: 1,
				cols: 2,
				data: [][]float64{
					{11, 26},
				},
			},
			wantError: false,
		},
		{
			name: "multiply 2x3 and 3x2 matrices",
			args: args{
				m1: Matrix{
					rows: 2,
					cols: 3,
					data: [][]float64{
						{1, 2, 3},
						{3, 2, 1},
					},
				},
				m2: Matrix{
					rows: 3,
					cols: 2,
					data: [][]float64{
						{1, 2},
						{2, 1},
						{1, 2},
					},
				},
			},
			want: &Matrix{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{8, 10},
					{8, 10},
				},
			},
		},
		{
			name: "multiply two 1x1 matrices",
			args: args{
				m1: Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{1},
					},
				},
				m2: Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{4},
					},
				},
			},
			want: &Matrix{
				rows: 1,
				cols: 1,
				data: [][]float64{
					{4},
				},
			},
			wantError: false,
		},
		{
			name: "multiply two 2x2 and 4x2 matrices for error",
			args: args{
				m1: Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{1, 1},
						{1, 1},
					},
				},
				m2: Matrix{
					rows: 4,
					cols: 2,
					data: [][]float64{
						{1, 1},
						{1, 1},
						{1, 1},
						{1, 1},
					},
				},
			},
			want:      nil,
			wantError: true,
		},
		{
			name: "multiply 4x4 matrix by 4x4 identity matrix",
			args: args{
				m1: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 8, 7, 6},
						{5, 4, 3, 2},
					},
				},
				m2: *NewIdentityMatrix(4),
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 8, 7, 6},
					{5, 4, 3, 2},
				},
			},
			wantError: false,
		},
		{
			name: "multiply 4x4 identity matrix by 4x1 tuple",
			args: args{
				m1: *NewIdentityMatrix(4),
				m2: Matrix{
					rows: 4,
					cols: 1,
					data: [][]float64{
						{1},
						{5},
						{9},
						{5},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 1,
				data: [][]float64{
					{1},
					{5},
					{9},
					{5},
				},
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Multiply(tt.args.m1, tt.args.m2)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, m)
				assert.Equal(t, true, m.Equals(tt.want))
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	type args struct {
		m1 Matrix
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "transposition of 4x4 matrix",
			args: args{
				m1: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{0, 9, 3, 0},
						{9, 8, 0, 8},
						{1, 8, 5, 3},
						{0, 0, 5, 8},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{0, 9, 1, 0},
					{9, 8, 8, 0},
					{3, 0, 5, 5},
					{0, 8, 3, 8},
				},
			},
		},
		{
			name: "transposition of 2x2 matrix",
			args: args{
				m1: Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{1, 2},
						{3, 4},
					},
				},
			},
			want: &Matrix{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{1, 3},
					{2, 4},
				},
			},
		},
		{
			name: "transposition of 2x3 matrix",
			args: args{
				m1: Matrix{
					rows: 2,
					cols: 3,
					data: [][]float64{
						{1, 2, 3},
						{4, 5, 6},
					},
				},
			},
			want: &Matrix{
				rows: 3,
				cols: 2,
				data: [][]float64{
					{1, 4},
					{2, 5},
					{3, 6},
				},
			},
		},
		{
			name: "transposition of the 4x4 identity matrix is the 4x4 identity matrix",
			args: args{
				m1: *NewIdentityMatrix(4),
			},
			want: NewIdentityMatrix(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Transpose(tt.args.m1))
		})
	}
}

func TestSubmatrix(t *testing.T) {
	type args struct {
		m   Matrix
		row uint
		col uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name: "submatrix of a 3x3 matrix is a 2x2 matrix",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{1, 5, 0},
						{-3, 2, 7},
						{0, 6, -3},
					},
				},
				row: 0,
				col: 2,
			},
			want: &Matrix{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{-3, 2},
					{0, 6},
				},
			},
			wantErr: false,
		},
		{
			name: "submatrix of a 4x4 matrix is a 3x3 matrix",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-6, 1, 1, 6},
						{-8, 5, 8, 6},
						{-1, 0, 8, 2},
						{-7, 1, -1, 1},
					},
				},
				row: 2,
				col: 1,
			},
			want: &Matrix{
				rows: 3,
				cols: 3,
				data: [][]float64{
					{-6, 1, 6},
					{-8, 8, 6},
					{-7, -1, 1},
				},
			},
			wantErr: false,
		},
		{
			name: "submatrix of a 4x4 matrix with row out of bounds",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-6, 1, 1, 6},
						{-8, 5, 8, 6},
						{-1, 0, 8, 2},
						{-7, 1, -1, 1},
					},
				},
				row: 4,
				col: 1,
			},
			want: &Matrix{
				rows: 3,
				cols: 3,
				data: [][]float64{
					{-6, 1, 6},
					{-8, 8, 6},
					{-7, -1, 1},
				},
			},
			wantErr: true,
		},
		{
			name: "submatrix of a 4x4 matrix with col out of bounds",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-6, 1, 1, 6},
						{-8, 5, 8, 6},
						{-1, 0, 8, 2},
						{-7, 1, -1, 1},
					},
				},
				row: 0,
				col: 4,
			},
			want: &Matrix{
				rows: 3,
				cols: 3,
				data: [][]float64{
					{-6, 1, 6},
					{-8, 8, 6},
					{-7, -1, 1},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Submatrix(tt.args.m, tt.args.row, tt.args.col)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMinor(t *testing.T) {
	type args struct {
		m   Matrix
		row uint
		col uint
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "minor of a 3x3 matrix",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 1,
				col: 0,
			},
			want:    25,
			wantErr: false,
		},
		{
			name: "minor of a 2x2 matrix",
			args: args{
				m: Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{3, 3},
						{5, 5},
					},
				},
				row: 0,
				col: 0,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "minor of a 3x3 matrix with submatrix row out of bounds for an error",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 5,
				col: 0,
			},
			want:    25,
			wantErr: true,
		},
		{
			name: "minor of a 3x3 matrix with submatrix col out of bounds for an error",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 1,
				col: 4,
			},
			want:    25,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Minor(tt.args.m, tt.args.row, tt.args.col)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCofactor(t *testing.T) {
	type args struct {
		m   Matrix
		row uint
		col uint
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "cofactor of a 3x3 matrix 1",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 0,
				col: 0,
			},
			want:    -12,
			wantErr: false,
		},
		{
			name: "cofactor of a 3x3 matrix 2",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 1,
				col: 0,
			},
			want:    -25,
			wantErr: false,
		},
		{
			name: "cofactor of a 4x4 matrix",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-5, 2, 6, -8},
						{1, -5, 1, 8},
						{7, 7, -6, -7},
						{1, -3, 7, 4},
					},
				},
				row: 0,
				col: 0,
			},
			want:    116,
			wantErr: false,
		},
		{
			name: "cofactor of a 3x3 matrix with submatrix row out of bounds for an error",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 5,
				col: 0,
			},
			want:    25,
			wantErr: true,
		},
		{
			name: "cofactor of a 3x3 matrix with submatrix col out of bounds for an error",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{3, 5, 0},
						{2, -1, -7},
						{6, -1, 5},
					},
				},
				row: 1,
				col: 4,
			},
			want:    25,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cofactor(tt.args.m, tt.args.row, tt.args.col)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDeterminant(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "determinant of 1x1 matrix",
			args: args{
				m: Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{-11},
					},
				},
			},
			want: -11,
		},
		{
			name: "determinant of 2x2 matrix",
			args: args{
				m: Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{1, -2},
						{1, 3},
					},
				},
			},
			want: 5,
		},
		{
			name: "determinant of 3x3 matrix",
			args: args{
				m: Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{1, 2, 6},
						{-5, 8, -4},
						{2, 6, 4},
					},
				},
			},
			want: -196,
		},
		{
			name: "determinant of 4x4 matrix",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-2, -8, 3, 5},
						{-3, 1, 7, 3},
						{1, 2, -9, 6},
						{-6, 7, 7, -9},
					},
				},
			},
			want: -4071,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Determinant(tt.args.m))
		})
	}
}

func TestIsInvertible(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "4x4 matrix is invertible",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{6, 4, 4, 4},
						{5, 5, 7, 6},
						{4, -9, 3, -7},
						{9, 1, 7, -6},
					},
				},
			},
			want: true,
		},
		{
			name: "4x4 matrix is not invertible",
			args: args{
				m: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-4, 2, -2, -3},
						{9, 6, 2, 6},
						{0, -5, 1, -5},
						{0, 0, 0, 0},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsInvertible(tt.args.m))
		})
	}
}

func TestInverse(t *testing.T) {
	type args struct {
		a Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name: "inverse of 4x4 matrix",
			args: args{
				a: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-5, 2, 6, -8},
						{1, -5, 1, 8},
						{7, 7, -6, -7},
						{1, -3, 7, 4},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{0.21805, 0.45113, 0.24060, -0.04511},
					{-0.80827, -1.45677, -0.44361, 0.52068},
					{-0.07895, -0.22368, -0.05263, 0.19737},
					{-0.52256, -0.81391, -0.30075, 0.30639},
				},
			},
			wantErr: false,
		},
		{
			name: "inverse of 4x4 matrix that's not invertible",
			args: args{
				a: Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{-4, 2, -2, -3},
						{9, 6, 2, 6},
						{0, -5, 1, -5},
						{0, 0, 0, 0},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{0.21805, 0.45113, 0.24060, -0.04511},
					{-0.80827, -1.45677, -0.44361, 0.52068},
					{-0.07895, -0.22368, -0.05263, 0.19737},
					{-0.52256, -0.81391, -0.30075, 0.30639},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// inverse of a is b
			b, err := Inverse(tt.args.a)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, b)
			} else {
				assert.NotNil(t, b)
				assert.True(t, true, b.Equals(tt.want))

				// determinant of a is 532.0
				assert.Equal(t, 532.0, Determinant(tt.args.a))

				// cofactor of a at 2,3 is -160.0
				cf, err := Cofactor(tt.args.a, 2, 3)
				assert.NoError(t, err)
				assert.Equal(t, -160.0, cf)

				// b[3][2] is -160.0/532.0
				assert.True(t, math_utils.Float64Equals(-160.0/532.0, b.data[3][2],
					math_utils.Epsilon))

				// cofactor of a at 3,2 is 105.0
				cf, err = Cofactor(tt.args.a, 3, 2)
				assert.NoError(t, err)
				assert.Equal(t, 105.0, cf)

				// b[2][3] is 105.0/532.0
				assert.True(t, math_utils.Float64Equals(105.0/532.0, b.data[2][3],
					math_utils.Epsilon))
			}
		})
	}
}
