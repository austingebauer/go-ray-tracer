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
		want Matrix
	}{
		{
			name: "create 4x4 matrix",
			args: args{
				rows: 4,
				cols: 4,
			},
			want: [][]float64{
				{0.0, 0.0, 0.0, 0.0},
				{0.0, 0.0, 0.0, 0.0},
				{0.0, 0.0, 0.0, 0.0},
				{0.0, 0.0, 0.0, 0.0},
			},
		},
		{
			name: "create 2x2 matrix",
			args: args{
				rows: 2,
				cols: 2,
			},
			want: [][]float64{
				{0.0, 0.0},
				{0.0, 0.0},
			},
		},
		{
			name: "create 3x1 matrix",
			args: args{
				rows: 3,
				cols: 1,
			},
			want: [][]float64{
				{0.0},
				{0.0},
				{0.0},
			},
		},
		{
			name: "create 0x0 matrix",
			args: args{
				rows: 0,
				cols: 0,
			},
			want: [][]float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewMatrix(tt.args.rows, tt.args.cols))
		})
	}
}
