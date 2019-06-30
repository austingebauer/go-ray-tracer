package matrix

import (
	"testing"

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
				m2: Identity,
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
				m1: Identity,
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
