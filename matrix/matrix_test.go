package matrix

import (
	"reflect"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
