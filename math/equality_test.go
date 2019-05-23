package math

import (
	"testing"
)

func TestFloat64Equals(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "float64s are equal",
			args: args{
				a: 1.1234001,
				b: 1.1234001,
				c: Epsilon,
			},
			want: true,
		},
		{
			name: "float64s are barely equal",
			args: args{
				a: 1.0,
				b: 0.99999,
				c: Epsilon,
			},
			want: true,
		},
		{
			name: "float64s are not equal",
			args: args{
				a: 2.1234000,
				b: -1.12,
				c: Epsilon,
			},
			want: false,
		},
		{
			name: "float64s are barely not equal",
			args: args{
				a: 1.0,
				b: 0.99998,
				c: Epsilon,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Equals(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Float64Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
