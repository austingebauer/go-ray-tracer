package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuplePoint(t *testing.T) {
	tplPoint := NewPoint(1, -2, 3.14)
	assert.NotNil(t, tplPoint)
	assert.Equal(t, float64(1), tplPoint.x)
	assert.Equal(t, float64(-2), tplPoint.y)
	assert.Equal(t, float64(3.14), tplPoint.z)
	assert.Equal(t, point, tplPoint.w)
}

func TestNewTupleVector(t *testing.T) {
	tplVector := NewVector(-1, 2.22, -3.14)
	assert.NotNil(t, tplVector)
	assert.Equal(t, float64(-1), tplVector.x)
	assert.Equal(t, float64(2.22), tplVector.y)
	assert.Equal(t, float64(-3.14), tplVector.z)
	assert.Equal(t, vector, tplVector.w)
}

func TestNewTupleErrorPositive(t *testing.T) {
	tpl, err := newTuple(-1, 2.22, -3.14, 1.0001)
	assert.Error(t, err)
	assert.Nil(t, tpl)
}

func TestNewTupleErrorNegative(t *testing.T) {
	tpl, err := newTuple(-1, 2.22, -3.14, -0.001)
	assert.Error(t, err)
	assert.Nil(t, tpl)
}

func TestEquals(t *testing.T) {
	type args struct {
		tpl1 *tuple
		tpl2 *tuple
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Tuples are equal",
			args: args{
				tpl1: &tuple{
					x: 1.0,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
				tpl2: &tuple{
					x: 1.0,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
			},
			want: true,
		},
		{
			name: "Tuples are barely equal x",
			args: args{
				tpl1: &tuple{
					x: 1.0,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
				tpl2: &tuple{
					x: 0.99999,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
			},
			want: true,
		},
		{
			name: "Tuples aren't equal x",
			args: args{
				tpl1: &tuple{
					x: 1.0,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
				tpl2: &tuple{
					x: -1.34,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
			},
			want: false,
		},
		{
			name: "Tuples aren't equal y",
			args: args{
				tpl1: &tuple{
					x: 1.0,
					y: 22.0,
					z: 3.0,
					w: 1.0,
				},
				tpl2: &tuple{
					x: -1.34,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
			},
			want: false,
		},
		{
			name: "Tuples aren't equal z",
			args: args{
				tpl1: &tuple{
					x: 1.0,
					y: -2.0,
					z: 3.0,
					w: 1.0,
				},
				tpl2: &tuple{
					x: -1.34,
					y: -2.0,
					z: -3.10,
					w: 1.0,
				},
			},
			want: false,
		},
		{
			name: "Tuples aren't equal w",
			args: args{
				tpl1: &tuple{
					x: 1.0,
					y: -2.0,
					z: 3.0,
					w: 3.14,
				},
				tpl2: &tuple{
					x: -1.34,
					y: -2.0,
					z: 3.0,
					w: 10.0000,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Equals(tt.args.tpl1, tt.args.tpl2))
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		tpl1 *tuple
		tpl2 *tuple
	}
	type want struct {
		err bool
		tpl *tuple
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Adds point and vector tuples",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 1,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 1,
					w: 0,
				},
			},
			want: want{
				tpl: &tuple{
					x: 1,
					y: 1,
					z: 6,
					w: 1,
				},
				err: false,
			},
		},
		{
			name: "Adds vector and vector tuples",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 0,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 1,
					w: 0,
				},
			},
			want: want{
				tpl: &tuple{
					x: 1,
					y: 1,
					z: 6,
					w: 0,
				},
				err: false,
			},
		},
		{
			name: "Adds point and point tuples and expects an error",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 1,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 1,
					w: 1,
				},
			},
			want: want{
				tpl: nil,
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualTpl, err := Add(tt.args.tpl1, tt.args.tpl2)
			if tt.want.err {
				assert.Error(t, err)
				assert.Nil(t, actualTpl)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.tpl, actualTpl)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		tpl1 *tuple
		tpl2 *tuple
	}
	type want struct {
		err bool
		tpl *tuple
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Subtracts vector from point tuples",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 1,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 1,
					w: 0,
				},
			},
			want: want{
				tpl: &tuple{
					x: 5,
					y: -5,
					z: 4,
					w: 1,
				},
				err: false,
			},
		},
		{
			name: "Subtracts vector and vector tuples",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 0,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 7,
					w: 0,
				},
			},
			want: want{
				tpl: &tuple{
					x: 5,
					y: -5,
					z: -2,
					w: 0,
				},
				err: false,
			},
		},
		{
			name: "Negates a vector by subtracting from the zero vector",
			args: args{
				tpl1: &tuple{
					x: 0,
					y: 0,
					z: 0,
					w: 0,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 7,
					w: 0,
				},
			},
			want: want{
				tpl: &tuple{
					x: 2,
					y: -3,
					z: -7,
					w: 0,
				},
				err: false,
			},
		},
		{
			name: "Subtracts point from vector and expects an error",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 0,
				},
				tpl2: &tuple{
					x: -2,
					y: 3,
					z: 1,
					w: 1,
				},
			},
			want: want{
				tpl: nil,
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualTpl, err := Subtract(tt.args.tpl1, tt.args.tpl2)
			if tt.want.err {
				assert.Error(t, err)
				assert.Nil(t, actualTpl)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.tpl, actualTpl)
			}
		})
	}
}

func TestNegate(t *testing.T) {
	type args struct {
		tpl1 *tuple
	}
	type want struct {
		err bool
		tpl *tuple
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Negates a vector",
			args: args{
				tpl1: &tuple{
					x: 3,
					y: -2,
					z: 5,
					w: 0,
				},
			},
			want: want{
				tpl: &tuple{
					x: -3,
					y: 2,
					z: -5,
					w: 0,
				},
				err: false,
			},
		},
		{
			name: "Negates a point",
			args: args{
				tpl1: &tuple{
					x: 1,
					y: 1,
					z: 1,
					w: 1,
				},
			},
			want: want{
				tpl: &tuple{
					x: -1,
					y: -1,
					z: -1,
					w: 1,
				},
				err: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualTpl, err := Negate(tt.args.tpl1)
			if tt.want.err {
				assert.Error(t, err)
				assert.Nil(t, actualTpl)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.tpl, actualTpl)
			}
		})
	}
}
