package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuplePoint(t *testing.T) {
	tplPoint, err := NewTuple(1, -2, 3.14, 1.0)
	assert.NoError(t, err)
	assert.Equal(t, float64(1), tplPoint.X)
	assert.Equal(t, float64(-2), tplPoint.Y)
	assert.Equal(t, float64(3.14), tplPoint.Z)
	assert.Equal(t, float64(1.0), tplPoint.W)
}

func TestNewTupleVector(t *testing.T) {
	tplVector, err := NewTuple(-1, 2.22, -3.14, 0.0)
	assert.NoError(t, err)
	assert.Equal(t, float64(-1), tplVector.X)
	assert.Equal(t, float64(2.22), tplVector.Y)
	assert.Equal(t, float64(-3.14), tplVector.Z)
	assert.Equal(t, float64(0.0), tplVector.W)
}

func TestNewTupleWErrorPositive(t *testing.T) {
	tpl, err := NewTuple(-1, 2.22, -3.14, 1.0001)
	assert.Error(t, err)
	assert.Nil(t, tpl)
}

func TestNewTupleWErrorNegative(t *testing.T) {
	tpl, err := NewTuple(-1, 2.22, -3.14, -0.001)
	assert.Error(t, err)
	assert.Nil(t, tpl)
}

func TestTuple_Equals(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
		W float64
	}
	type args struct {
		tpl2 *Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Tuples are equal",
			fields: fields{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: 1.0,
			},
			args: args{
				tpl2: &Tuple{
					X: 1.0,
					Y: -2.0,
					Z: 3.0,
					W: 1.0,
				},
			},
			want: true,
		},
		{
			name: "Tuples are barely equal X",
			fields: fields{
				X: 0.99999,
				Y: -2.0,
				Z: 3.0,
				W: 1.0,
			},
			args: args{
				tpl2: &Tuple{
					X: 1.0,
					Y: -2.0,
					Z: 3.0,
					W: 1.0,
				},
			},
			want: true,
		},
		{
			name: "Tuples aren't equal X",
			fields: fields{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: 1.0,
			},
			args: args{
				tpl2: &Tuple{
					X: 1.1234,
					Y: -2.0,
					Z: 3.0,
					W: 1.0,
				},
			},
			want: false,
		},
		{
			name: "Tuples aren't equal Y",
			fields: fields{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: 1.0,
			},
			args: args{
				tpl2: &Tuple{
					X: 1.1234,
					Y: 2.0001,
					Z: 3.0,
					W: 1.0,
				},
			},
			want: false,
		},
		{
			name: "Tuples aren't equal Z",
			fields: fields{
				X: 1.0,
				Y: -2.0,
				Z: 3.14,
				W: 1.0,
			},
			args: args{
				tpl2: &Tuple{
					X: 1.1234,
					Y: -2.0,
					Z: -3.0,
					W: 1.0,
				},
			},
			want: false,
		},
		{
			name: "Tuples aren't equal W",
			fields: fields{
				X: 1.0,
				Y: -2.0,
				Z: 3.0,
				W: 1.0,
			},
			args: args{
				tpl2: &Tuple{
					X: 1.1234,
					Y: -2.0,
					Z: 3.0,
					W: 1.1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tpl1 := &Tuple{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := tpl1.Equals(tt.args.tpl2); got != tt.want {
				t.Errorf("Tuple.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
