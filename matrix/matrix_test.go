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
		mQ *Matrix
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
				mQ: &Matrix{
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
				mQ: &Matrix{
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
				mQ: &Matrix{
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
			assert.Equal(t, tt.want, m.Equals(tt.args.mQ))
		})
	}
}
