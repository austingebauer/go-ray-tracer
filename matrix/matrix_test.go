package matrix

import (
	"math"
	"reflect"
	"testing"

	"github.com/austingebauer/go-ray-tracer/maths"
	"github.com/stretchr/testify/assert"
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
		m1 *Matrix
		m2 *Matrix
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
				m1: &Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 8, 7, 6},
						{5, 4, 3, 2},
					},
				},
				m2: &Matrix{
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
				m1: &Matrix{
					rows: 1,
					cols: 3,
					data: [][]float64{
						{1, 2, 3},
					},
				},
				m2: &Matrix{
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
				m1: &Matrix{
					rows: 2,
					cols: 3,
					data: [][]float64{
						{1, 2, 3},
						{3, 2, 1},
					},
				},
				m2: &Matrix{
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
				m1: &Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{1},
					},
				},
				m2: &Matrix{
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
				m1: &Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{1, 1},
						{1, 1},
					},
				},
				m2: &Matrix{
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
				m1: &Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 8, 7, 6},
						{5, 4, 3, 2},
					},
				},
				m2: NewIdentityMatrix(4),
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
				m1: NewIdentityMatrix(4),
				m2: &Matrix{
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
		m   *Matrix
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
				m: &Matrix{
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
				m: &Matrix{
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
				m: &Matrix{
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
				m: &Matrix{
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
		m   *Matrix
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
				m: &Matrix{
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
				m: &Matrix{
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
			name: "minor of a non-square matrix for an error",
			args: args{
				m: &Matrix{
					rows: 2,
					cols: 3,
					data: [][]float64{
						{3, 3, 4},
						{5, 5, 4},
					},
				},
				row: 0,
				col: 0,
			},
			want:    5,
			wantErr: true,
		},
		{
			name: "minor of a 3x3 matrix with submatrix row out of bounds for an error",
			args: args{
				m: &Matrix{
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
				m: &Matrix{
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
		m   *Matrix
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
				m: &Matrix{
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
				m: &Matrix{
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
				m: &Matrix{
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
				m: &Matrix{
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
				m: &Matrix{
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
		m *Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "determinant of 1x1 matrix",
			args: args{
				m: &Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{-11},
					},
				},
			},
			want:    -11,
			wantErr: false,
		},
		{
			name: "determinant of 2x2 matrix",
			args: args{
				m: &Matrix{
					rows: 2,
					cols: 2,
					data: [][]float64{
						{1, -2},
						{1, 3},
					},
				},
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "determinant of 3x3 matrix",
			args: args{
				m: &Matrix{
					rows: 3,
					cols: 3,
					data: [][]float64{
						{1, 2, 6},
						{-5, 8, -4},
						{2, 6, 4},
					},
				},
			},
			want:    -196,
			wantErr: false,
		},
		{
			name: "determinant of 4x4 matrix",
			args: args{
				m: &Matrix{
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
			want:    -4071,
			wantErr: false,
		},
		{
			name: "determinant of 3x4 non-square matrix returns error",
			args: args{
				m: &Matrix{
					rows: 3,
					cols: 4,
					data: [][]float64{
						{-2, -8, 3, 5},
						{-3, 1, 7, 3},
						{1, 2, -9, 6},
					},
				},
			},
			want:    123,
			wantErr: true,
		},
		{
			name: "determinant of 4x3 non-square matrix returns error",
			args: args{
				m: &Matrix{
					rows: 3,
					cols: 4,
					data: [][]float64{
						{-2, -8, 3},
						{-3, 1, 7},
						{1, 2, -9},
						{1, 2, -9},
					},
				},
			},
			want:    123,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			det, err := Determinant(tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, det)
			}
		})
	}
}

func TestIsInvertible(t *testing.T) {
	type args struct {
		m *Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "4x4 matrix is invertible",
			args: args{
				m: &Matrix{
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
				m: &Matrix{
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
		{
			name: "3x4 matrix is not invertible bc not a square matrix",
			args: args{
				m: &Matrix{
					rows: 3,
					cols: 4,
					data: [][]float64{
						{-4, 2, -2, -3},
						{9, 6, 2, 6},
						{0, -5, 1, -5},
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

func TestInverseDetailed(t *testing.T) {
	type args struct {
		a *Matrix
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
				a: &Matrix{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// inverse of a is b
			b, err := Inverse(tt.args.a)

			assert.NotNil(t, b)
			assert.NoError(t, err)
			assert.True(t, b.Equals(tt.want))

			// determinant of a is 532.0
			det, err := Determinant(tt.args.a)
			assert.NoError(t, err)
			assert.Equal(t, 532.0, det)

			// cofactor of a at 2,3 is -160.0
			cf, err := Cofactor(tt.args.a, 2, 3)
			assert.NoError(t, err)
			assert.Equal(t, -160.0, cf)

			// b[3][2] is -160.0/532.0
			assert.True(t, maths.Float64Equals(-160.0/532.0, b.data[3][2],
				maths.Epsilon))

			// cofactor of a at 3,2 is 105.0
			cf, err = Cofactor(tt.args.a, 3, 2)
			assert.NoError(t, err)
			assert.Equal(t, 105.0, cf)

			// b[2][3] is 105.0/532.0
			assert.True(t, maths.Float64Equals(105.0/532.0, b.data[2][3],
				maths.Epsilon))
		})
	}
}

func TestInverse(t *testing.T) {
	type args struct {
		a *Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name: "inverse of 4x4 matrix that's not invertible",
			args: args{
				a: &Matrix{
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
		{
			name: "inverse of 1x1 matrix",
			args: args{
				a: &Matrix{
					rows: 1,
					cols: 1,
					data: [][]float64{
						{2},
					},
				},
			},
			want: &Matrix{
				rows: 1,
				cols: 1,
				data: [][]float64{
					{0.5},
				},
			},
			wantErr: false,
		},
		{
			name: "inverse of 4x4 matrix 1",
			args: args{
				a: &Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{8, -5, 9, 2},
						{7, 5, 6, 1},
						{-6, 0, 9, 6},
						{-3, 0, -9, -4},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{-0.15385, -0.15385, -0.28205, -0.53846},
					{-0.07692, 0.12308, 0.02564, 0.03077},
					{0.35897, 0.35897, 0.43590, 0.92308},
					{-0.69231, -0.69231, -0.76923, -1.92308},
				},
			},
			wantErr: false,
		},
		{
			name: "inverse of 4x4 matrix 2",
			args: args{
				a: &Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{9, 3, 0, 9},
						{-5, -2, -6, -3},
						{-4, 9, 6, 4},
						{-7, 6, 6, 2},
					},
				},
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{-0.04074, -0.07778, 0.14444, -0.22222},
					{-0.07778, 0.03333, 0.36667, -0.33333},
					{-0.02901, -0.14630, -0.10926, 0.12963},
					{0.17778, 0.06667, -0.26667, 0.33333},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := Inverse(tt.args.a)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, b)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, b)
				assert.True(t, b.Equals(tt.want))
			}
		})
	}
}

func TestMultiplyProductByInverse(t *testing.T) {
	type args struct {
		a *Matrix
		b *Matrix
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "matrix product (c = a * b) multiplied by its inverse (c * inverse(b) = a)",
			args: args{
				a: &Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{3, -9, 7, 3},
						{3, -8, 2, -9},
						{-4, 4, 4, 1},
						{-6, 5, -1, 1},
					},
				},
				b: &Matrix{
					rows: 4,
					cols: 4,
					data: [][]float64{
						{8, 2, 2, 2},
						{3, -1, 7, 0},
						{7, 0, 5, 4},
						{6, -2, 0, 5},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// cA = a * b
			cA, err := Multiply(tt.args.a, tt.args.b)
			assert.NoError(t, err)
			assert.NotNil(t, cA)

			// cB = b * a
			cB, err := Multiply(tt.args.b, tt.args.a)
			assert.NoError(t, err)
			assert.NotNil(t, cB)

			// a = cA * inverse(b)
			inverseB, err := Inverse(tt.args.b)
			assert.NoError(t, err)
			assert.NotNil(t, inverseB)
			shouldBeA, err := Multiply(cA, inverseB)
			assert.NoError(t, err)
			assert.NotNil(t, shouldBeA)
			assert.True(t, tt.args.a.Equals(shouldBeA))

			// b = cB * inverse(a)
			inverseA, err := Inverse(tt.args.a)
			assert.NoError(t, err)
			assert.NotNil(t, inverseA)
			shouldBeB, err := Multiply(cB, inverseA)
			assert.NoError(t, err)
			assert.NotNil(t, shouldBeB)
			assert.True(t, tt.args.b.Equals(shouldBeB))
		})
	}
}

func TestMatrix_GetRows(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "get the number of rows in the matrix",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}

			assert.Equal(t, tt.want, m.GetRows())
		})
	}
}

func TestMatrix_GetCols(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "get the number of columns in the matrix",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}

			assert.Equal(t, tt.want, m.GetCols())
		})
	}
}

func TestMatrix_SetValue(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	type args struct {
		row uint
		col uint
		val float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set value in the row and column of the matrix",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			args: args{
				row: 1,
				col: 1,
				val: -99,
			},
			wantErr: false,
		},
		{
			name: "set value in the row and column of the matrix out of bounds",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			args: args{
				row: 1,
				col: 4,
				val: -99,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}

			err := m.SetValue(tt.args.row, tt.args.col, tt.args.val)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, m.data[tt.args.row][tt.args.col], tt.args.val)
			}
		})
	}
}

func TestMatrix_GetValue(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	type args struct {
		row uint
		col uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "get value in the row and column of the matrix",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{0, 0},
					{0, -99},
				},
			},
			args: args{
				row: 1,
				col: 1,
			},
			want:    -99,
			wantErr: false,
		},
		{
			name: "get value in the row and column of the matrix out of bounds",
			fields: fields{
				rows: 2,
				cols: 3,
				data: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			args: args{
				row: 1,
				col: 4,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}

			got, err := m.GetValue(tt.args.row, tt.args.col)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestNewIdentityMatrix(t *testing.T) {
	type args struct {
		size uint
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "1x1 identity matrix",
			args: args{
				size: 1,
			},
			want: &Matrix{
				rows: 1,
				cols: 1,
				data: [][]float64{
					{1},
				},
			},
		},
		{
			name: "2x2 identity matrix",
			args: args{
				size: 2,
			},
			want: &Matrix{
				rows: 2,
				cols: 2,
				data: [][]float64{
					{1, 0},
					{0, 1},
				},
			},
		},
		{
			name: "4x4 identity matrix",
			args: args{
				size: 4,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewIdentityMatrix(tt.args.size), tt.want)
		})
	}
}

func TestCheckInBounds(t *testing.T) {
	type args struct {
		m   *Matrix
		row uint
		col uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check that row and column are in bounds of the matrix",
			args: args{
				m:   NewMatrix(4, 4),
				row: 3,
				col: 0,
			},
			wantErr: false,
		},
		{
			name: "check that row and column are in bounds of the matrix",
			args: args{
				m:   NewMatrix(4, 4),
				row: 0,
				col: 3,
			},
			wantErr: false,
		},
		{
			name: "check that row and column are out of bounds of the matrix",
			args: args{
				m:   NewMatrix(4, 4),
				row: 3,
				col: 4,
			},
			wantErr: true,
		},
		{
			name: "check that row and column are out of bounds of the matrix",
			args: args{
				m:   NewMatrix(4, 4),
				row: 4,
				col: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckInBounds(tt.args.m, tt.args.row, tt.args.col)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewTranslationMatrix(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "translation matrix with x, y, and z values",
			args: args{
				x: 5,
				y: -3,
				z: 2,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 5},
					{0, 1, 0, -3},
					{0, 0, 1, 2},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewTranslationMatrix(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestMatrix_Translate(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		{
			name: "translate identity matrix with x, y, and z values",
			fields: fields{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			args: args{
				x: 3,
				y: 2,
				z: 1,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 3},
					{0, 1, 0, 2},
					{0, 0, 1, 1},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			assert.Equal(t, tt.want, m.Translate(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestNewScalingMatrix(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "scaling matrix with x, y, and z values",
			args: args{
				x: 5,
				y: -3,
				z: 2,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{5, 0, 0, 0},
					{0, -3, 0, 0},
					{0, 0, 2, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewScalingMatrix(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestMatrix_Scale(t *testing.T) {
	type fields struct {
		rows uint
		cols uint
		data [][]float64
	}
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		{
			name: "scale identity matrix with x, y, and z values",
			fields: fields{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			args: args{
				x: 3,
				y: 2,
				z: 1,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{3, 0, 0, 0},
					{0, 2, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Scale(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewXRotationMatrix(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "x-axis rotation matrix for radians",
			args: args{
				radians: math.Pi,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 0},
					{0, math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0},
					{0, math.Sin(math.Pi), math.Cos(math.Pi), 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewXRotationMatrix(tt.args.radians))
		})
	}
}

func TestMatrix_RotateX(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		m    *Matrix
		args args
		want *Matrix
	}{
		{
			name: "rotate identity matrix around x-axis",
			m:    NewIdentityMatrix(4),
			args: args{
				radians: math.Pi,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 0, 0, 0},
					{0, math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0},
					{0, math.Sin(math.Pi), math.Cos(math.Pi), 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.RotateX(tt.args.radians))
		})
	}
}

func TestNewYRotationMatrix(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "y-axis rotation matrix for radians",
			args: args{
				radians: math.Pi,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{math.Cos(math.Pi), 0, math.Sin(math.Pi), 0},
					{0, 1, 0, 0},
					{-1 * math.Sin(math.Pi), 0, math.Cos(math.Pi), 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewYRotationMatrix(tt.args.radians))
		})
	}
}

func TestMatrix_RotateY(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		m    *Matrix
		args args
		want *Matrix
	}{
		{
			name: "rotate identity matrix around y-axis",
			m:    NewIdentityMatrix(4),
			args: args{
				radians: math.Pi,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{math.Cos(math.Pi), 0, math.Sin(math.Pi), 0},
					{0, 1, 0, 0},
					{-1 * math.Sin(math.Pi), 0, math.Cos(math.Pi), 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.RotateY(tt.args.radians))
		})
	}
}

func TestNewZRotationMatrix(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "z-axis rotation matrix for radians",
			args: args{
				radians: math.Pi,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0, 0},
					{math.Sin(math.Pi), math.Cos(math.Pi), 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewZRotationMatrix(tt.args.radians))
		})
	}
}

func TestMatrix_RotateZ(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		m    *Matrix
		args args
		want *Matrix
	}{
		{
			name: "rotate identity matrix around z-axis",
			m:    NewIdentityMatrix(4),
			args: args{
				radians: math.Pi,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0, 0},
					{math.Sin(math.Pi), math.Cos(math.Pi), 0, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.RotateZ(tt.args.radians))
		})
	}
}

func TestNewShearingMatrix(t *testing.T) {
	type args struct {
		xy float64
		xz float64
		yx float64
		yz float64
		zx float64
		zy float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "shearing matrix with relevant coordinates",
			args: args{
				xy: 1,
				xz: 2,
				yx: 3,
				yz: 4,
				zx: 5,
				zy: 6,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 1, 2, 0},
					{3, 1, 4, 0},
					{5, 6, 1, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShearingMatrix(tt.args.xy, tt.args.xz, tt.args.yx, tt.args.yz, tt.args.zx, tt.args.zy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShearingMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Shear(t *testing.T) {
	type args struct {
		xy float64
		xz float64
		yx float64
		yz float64
		zx float64
		zy float64
	}
	tests := []struct {
		name string
		m    *Matrix
		args args
		want *Matrix
	}{
		{
			name: "shear an identity matrix",
			m:    NewIdentityMatrix(4),
			args: args{
				xy: 1,
				xz: 2,
				yx: 3,
				yz: 4,
				zx: 5,
				zy: 6,
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: [][]float64{
					{1, 1, 2, 0},
					{3, 1, 4, 0},
					{5, 6, 1, 0},
					{0, 0, 0, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.m.Shear(tt.args.xy, tt.args.xz, tt.args.yx, tt.args.yz, tt.args.zx, tt.args.zy))
		})
	}
}
