package point

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/vector"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "NewPoint sets x, y, and z",
			args: args{
				x: float64(1),
				y: float64(-2),
				z: float64(0),
			},
			want: &Point{
				X: 1,
				Y: -2,
				Z: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPoint(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestPoint_Equals(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		ptQ *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "vectors are equal",
			fields: fields{
				1,
				2,
				3,
			},
			args: args{
				&Point{
					1,
					2,
					3,
				},
			},
			want: true,
		},
		{
			name: "vectors aren't equal",
			fields: fields{
				1,
				2,
				3,
			},
			args: args{
				&Point{
					1,
					-2,
					3,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, pt.Equals(tt.args.ptQ))
		})
	}
}

func TestPoint_Negate(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Point
	}{
		{
			name: "negate a point positive",
			fields: fields{
				1,
				2,
				3,
			},
			want: &Point{
				-1,
				-2,
				-3,
			},
		},
		{
			name: "negate a point negative",
			fields: fields{
				-1,
				-2,
				-3,
			},
			want: &Point{
				1,
				2,
				3,
			},
		},
		{
			name: "negate an origin zero point",
			fields: fields{
				0,
				0,
				0,
			},
			want: &Point{
				0,
				0,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := pt.Negate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Scale(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		scalar float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		{
			name: "scale the point",
			fields: fields{
				0,
				2,
				-3,
			},
			args: args{
				scalar: float64(2),
			},
			want: &Point{
				0,
				4,
				-6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, pt.Scale(tt.args.scalar))
		})
	}
}

func TestPoint_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec vector.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		{
			name: "point add vector method",
			fields: fields{
				0,
				-1,
				1,
			},
			args: args{
				vector.Vector{
					X: 2,
					Y: 3,
					Z: -1,
				},
			},
			want: &Point{
				2,
				2,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}

			assert.Equal(t, tt.want, pt.Add(tt.args.vec))
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		pt  Point
		vec vector.Vector
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		{
			name: "point add vector function",
			args: args{
				pt: Point{
					0,
					-1,
					1,
				},
				vec: vector.Vector{
					X: 2,
					Y: 3,
					Z: -1,
				},
			},
			want: Point{
				2,
				2,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Add(tt.args.pt, tt.args.vec))
		})
	}
}

func TestPoint_Subtract(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec vector.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		{
			name: "point subtract vector method",
			fields: fields{
				0,
				-1,
				1,
			},
			args: args{
				vector.Vector{
					X: 2,
					Y: 3,
					Z: -1,
				},
			},
			want: &Point{
				-2,
				-4,
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}

			assert.Equal(t, tt.want, pt.Subtract(tt.args.vec))
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		pt1 Point
		pt2 Point
	}
	tests := []struct {
		name string
		args args
		want vector.Vector
	}{
		{
			name: "subtract two points for vector function",
			args: args{
				pt1: Point{
					-1,
					0,
					1,
				},
				pt2: Point{
					-2,
					1,
					0,
				},
			},
			want: vector.Vector{
				X: 1,
				Y: -1,
				Z: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Subtract(tt.args.pt1, tt.args.pt2))
		})
	}
}

func TestTransformPoint(t *testing.T) {
	transform := matrix.Translation(5, -3, 2)
	p := NewPoint(-3, 4, 5)
	m, err := matrix.Multiply(*transform, *ToMatrix(*p))
	assert.NoError(t, err)
	assert.NotNil(t, m)

	ptMult, err := ToPoint(*m)
	assert.NoError(t, err)
	assert.Equal(t, NewPoint(2, 1, 7), ptMult)
}

func TestInverseTransformPoint(t *testing.T) {
	p := NewPoint(-3, 4, 5)
	transform := matrix.Translation(5, -3, 2)

	inverseT, err := matrix.Inverse(*transform)
	assert.NoError(t, err)
	assert.NotNil(t, inverseT)

	m, err := matrix.Multiply(*inverseT, *ToMatrix(*p))
	assert.NoError(t, err)
	assert.NotNil(t, m)

	ptMult, err := ToPoint(*m)
	assert.NoError(t, err)
	assert.Equal(t, NewPoint(-8, 7, 3), ptMult)
}

func TestToMatrix(t *testing.T) {
	type args struct {
		pt Point
	}
	tests := []struct {
		name string
		args args
		want *matrix.Matrix
	}{
		{
			name: "point to matrix conversion",
			args: args{
				pt: *NewPoint(1, -2, 3),
			},
			want: matrix.NewMatrix(4, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.want.SetValue(0, 0, 1)
			assert.NoError(t, err)
			err = tt.want.SetValue(1, 0, -2)
			assert.NoError(t, err)
			err = tt.want.SetValue(2, 0, 3)
			assert.NoError(t, err)
			err = tt.want.SetValue(3, 0, 1)
			assert.NoError(t, err)

			assert.Equal(t, tt.want, ToMatrix(tt.args.pt))
		})
	}
}

func TestToPoint(t *testing.T) {
	type args struct {
		m matrix.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    *Point
		wantErr bool
	}{
		{
			name: "matrix to point conversion",
			args: args{
				m: *matrix.NewMatrix(4, 1),
			},
			want:    NewPoint(1, -2, 3),
			wantErr: false,
		},
		{
			name: "matrix to point conversion error row length",
			args: args{
				m: *matrix.NewMatrix(2, 1),
			},
			want:    NewPoint(1, -2, 3),
			wantErr: true,
		},
		{
			name: "matrix to point conversion error col length",
			args: args{
				m: *matrix.NewMatrix(4, 2),
			},
			want:    NewPoint(1, -2, 3),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				got, err := ToPoint(tt.args.m)
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				err := tt.args.m.SetValue(0, 0, 1)
				assert.NoError(t, err)
				err = tt.args.m.SetValue(1, 0, -2)
				assert.NoError(t, err)
				err = tt.args.m.SetValue(2, 0, 3)
				assert.NoError(t, err)

				got, err := ToPoint(tt.args.m)
				assert.NotNil(t, got)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestScalingPoint(t *testing.T) {
	transform := matrix.Scaling(2, 3, 4)
	pt := NewPoint(-4, 6, 8)

	mult, err := matrix.Multiply(*transform, *ToMatrix(*pt))
	assert.NoError(t, err)

	ptMult, err := ToPoint(*mult)
	assert.NoError(t, err)

	assert.Equal(t, NewPoint(-8, 18, 32), ptMult)
}

func TestPointReflectionOverAxis(t *testing.T) {
	transform := matrix.Scaling(-1, 1, 1)
	pt := NewPoint(2, 3, 4)

	mult, err := matrix.Multiply(*transform, *ToMatrix(*pt))
	assert.NoError(t, err)

	ptReflectedOnX, err := ToPoint(*mult)
	assert.NoError(t, err)
	assert.Equal(t, NewPoint(-2, 3, 4), ptReflectedOnX)
}

func TestPointRotateXAxis(t *testing.T) {
	ptM := ToMatrix(*NewPoint(0, 1, 0))
	half_quarter := matrix.RotationX(math.Pi / 4)
	full_quarter := matrix.RotationX(math.Pi / 2)

	// rotate the point around the x axis Pi/4 radians
	rotM, err := matrix.Multiply(*half_quarter, *ptM)
	assert.NoError(t, err)
	rotMPoint, err := ToPoint(*rotM)
	assert.NoError(t, err)
	assert.True(t, NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2).Equals(rotMPoint))

	// rotate the point around the x axis Pi/2 radians
	rotM, err = matrix.Multiply(*full_quarter, *ptM)
	assert.NoError(t, err)
	rotMPoint, err = ToPoint(*rotM)
	assert.NoError(t, err)
	assert.True(t, NewPoint(0, 0, 1).Equals(rotMPoint))
}

func TestPointRotateYAxis(t *testing.T) {
	ptM := ToMatrix(*NewPoint(0, 0, 1))
	half_quarter := matrix.RotationY(math.Pi / 4)
	full_quarter := matrix.RotationY(math.Pi / 2)

	// rotate the point around the y axis Pi/4 radians
	rotM, err := matrix.Multiply(*half_quarter, *ptM)
	assert.NoError(t, err)
	rotMPoint, err := ToPoint(*rotM)
	assert.NoError(t, err)
	assert.True(t, NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equals(rotMPoint))

	// rotate the point around the y axis Pi/2 radians
	rotM, err = matrix.Multiply(*full_quarter, *ptM)
	assert.NoError(t, err)
	rotMPoint, err = ToPoint(*rotM)
	assert.NoError(t, err)
	assert.True(t, NewPoint(1, 0, 0).Equals(rotMPoint))
}

func TestPointRotateZAxis(t *testing.T) {
	ptM := ToMatrix(*NewPoint(0, 1, 0))
	half_quarter := matrix.RotationZ(math.Pi / 4)
	full_quarter := matrix.RotationZ(math.Pi / 2)

	// rotate the point around the z axis Pi/4 radians
	rotM, err := matrix.Multiply(*half_quarter, *ptM)
	assert.NoError(t, err)
	rotMPoint, err := ToPoint(*rotM)
	assert.NoError(t, err)
	assert.True(t, NewPoint(-1*math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equals(rotMPoint))

	// rotate the point around the z axis Pi/2 radians
	rotM, err = matrix.Multiply(*full_quarter, *ptM)
	assert.NoError(t, err)
	rotMPoint, err = ToPoint(*rotM)
	assert.NoError(t, err)
	assert.True(t, NewPoint(-1, 0, 0).Equals(rotMPoint))
}
