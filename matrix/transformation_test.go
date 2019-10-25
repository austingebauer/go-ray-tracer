package matrix

import (
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

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
				data: []float64{
					1, 0, 0, 5,
					0, 1, 0, -3,
					0, 0, 1, 2,
					0, 0, 0, 1,
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
		data []float64
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
				data: []float64{
					1, 0, 0, 0,
					0, 1, 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
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
				data: []float64{
					1, 0, 0, 3,
					0, 1, 0, 2,
					0, 0, 1, 1,
					0, 0, 0, 1,
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
				data: []float64{
					5, 0, 0, 0,
					0, -3, 0, 0,
					0, 0, 2, 0,
					0, 0, 0, 1,
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
		data []float64
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
				data: []float64{
					1, 0, 0, 0,
					0, 1, 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
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
				data: []float64{
					3, 0, 0, 0,
					0, 2, 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
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
				data: []float64{
					1, 0, 0, 0,
					0, math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0,
					0, math.Sin(math.Pi), math.Cos(math.Pi), 0,
					0, 0, 0, 1,
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
				data: []float64{
					1, 0, 0, 0,
					0, math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0,
					0, math.Sin(math.Pi), math.Cos(math.Pi), 0,
					0, 0, 0, 1,
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
				data: []float64{
					math.Cos(math.Pi), 0, math.Sin(math.Pi), 0,
					0, 1, 0, 0,
					-1 * math.Sin(math.Pi), 0, math.Cos(math.Pi), 0,
					0, 0, 0, 1,
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
				data: []float64{
					math.Cos(math.Pi), 0, math.Sin(math.Pi), 0,
					0, 1, 0, 0,
					-1 * math.Sin(math.Pi), 0, math.Cos(math.Pi), 0,
					0, 0, 0, 1,
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
				data: []float64{
					math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0, 0,
					math.Sin(math.Pi), math.Cos(math.Pi), 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
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
				data: []float64{
					math.Cos(math.Pi), -1 * math.Sin(math.Pi), 0, 0,
					math.Sin(math.Pi), math.Cos(math.Pi), 0, 0,
					0, 0, 1, 0,
					0, 0, 0, 1,
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
				data: []float64{
					1, 1, 2, 0,
					3, 1, 4, 0,
					5, 6, 1, 0,
					0, 0, 0, 1,
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
				data: []float64{
					1, 1, 2, 0,
					3, 1, 4, 0,
					5, 6, 1, 0,
					0, 0, 0, 1,
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

func TestTransformPoint(t *testing.T) {
	transform := NewTranslationMatrix(5, -3, 2)
	p := point.NewPoint(-3, 4, 5)
	m, err := Multiply(transform, PointToMatrix(p))
	assert.NoError(t, err)
	assert.NotNil(t, m)

	ptMult, err := MatrixToPoint(m)
	assert.NoError(t, err)
	assert.Equal(t, point.NewPoint(2, 1, 7), ptMult)
}

func TestInverseTransformPoint(t *testing.T) {
	p := point.NewPoint(-3, 4, 5)
	transform := NewTranslationMatrix(5, -3, 2)

	inverseT, err := Inverse(transform)
	assert.NoError(t, err)
	assert.NotNil(t, inverseT)

	m, err := Multiply(inverseT, PointToMatrix(p))
	assert.NoError(t, err)
	assert.NotNil(t, m)

	ptMult, err := MatrixToPoint(m)
	assert.NoError(t, err)
	assert.Equal(t, point.NewPoint(-8, 7, 3), ptMult)
}

func TestPointToMatrix(t *testing.T) {
	type args struct {
		pt *point.Point
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "point to matrix conversion",
			args: args{
				pt: point.NewPoint(1, -2, 3),
			},
			want: NewMatrix(4, 1),
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

			assert.Equal(t, tt.want, PointToMatrix(tt.args.pt))
		})
	}
}

func TestToPoint(t *testing.T) {
	type args struct {
		m *Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    *point.Point
		wantErr bool
	}{
		{
			name: "matrix to point conversion",
			args: args{
				m: NewMatrix(4, 1),
			},
			want:    point.NewPoint(1, -2, 3),
			wantErr: false,
		},
		{
			name: "matrix to point conversion error row length",
			args: args{
				m: NewMatrix(2, 1),
			},
			want:    point.NewPoint(1, -2, 3),
			wantErr: true,
		},
		{
			name: "matrix to point conversion error col length",
			args: args{
				m: NewMatrix(4, 2),
			},
			want:    point.NewPoint(1, -2, 3),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				got, err := MatrixToPoint(tt.args.m)
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				err := tt.args.m.SetValue(0, 0, 1)
				assert.NoError(t, err)
				err = tt.args.m.SetValue(1, 0, -2)
				assert.NoError(t, err)
				err = tt.args.m.SetValue(2, 0, 3)
				assert.NoError(t, err)

				got, err := MatrixToPoint(tt.args.m)
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestScalingPoint(t *testing.T) {
	transform := NewScalingMatrix(2, 3, 4)
	pt := point.NewPoint(-4, 6, 8)

	mult, err := Multiply(transform, PointToMatrix(pt))
	assert.NoError(t, err)

	ptMult, err := MatrixToPoint(mult)
	assert.NoError(t, err)

	assert.Equal(t, point.NewPoint(-8, 18, 32), ptMult)
}

func TestPointReflectionOverAxis(t *testing.T) {
	transform := NewScalingMatrix(-1, 1, 1)
	pt := point.NewPoint(2, 3, 4)

	mult, err := Multiply(transform, PointToMatrix(pt))
	assert.NoError(t, err)

	ptReflectedOnX, err := MatrixToPoint(mult)
	assert.NoError(t, err)
	assert.Equal(t, point.NewPoint(-2, 3, 4), ptReflectedOnX)
}

func TestPointRotateXAxis(t *testing.T) {
	ptM := PointToMatrix(point.NewPoint(0, 1, 0))
	halfQuarter := NewXRotationMatrix(math.Pi / 4)
	fullQuarter := NewXRotationMatrix(math.Pi / 2)

	// rotate the point around the x axis Pi/4 radians
	rotM, err := Multiply(halfQuarter, ptM)
	assert.NoError(t, err)
	rotMPoint, err := MatrixToPoint(rotM)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2).Equals(rotMPoint))

	// rotate the point around the x axis Pi/2 radians
	rotM, err = Multiply(fullQuarter, ptM)
	assert.NoError(t, err)
	rotMPoint, err = MatrixToPoint(rotM)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(0, 0, 1).Equals(rotMPoint))
}

func TestPointRotateYAxis(t *testing.T) {
	ptM := PointToMatrix(point.NewPoint(0, 0, 1))
	halfQuarter := NewYRotationMatrix(math.Pi / 4)
	fullQuarter := NewYRotationMatrix(math.Pi / 2)

	// rotate the point around the y axis Pi/4 radians
	rotM, err := Multiply(halfQuarter, ptM)
	assert.NoError(t, err)
	rotMPoint, err := MatrixToPoint(rotM)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equals(rotMPoint))

	// rotate the point around the y axis Pi/2 radians
	rotM, err = Multiply(fullQuarter, ptM)
	assert.NoError(t, err)
	rotMPoint, err = MatrixToPoint(rotM)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(1, 0, 0).Equals(rotMPoint))
}

func TestPointRotateZAxis(t *testing.T) {
	ptM := PointToMatrix(point.NewPoint(0, 1, 0))
	halfQuarter := NewZRotationMatrix(math.Pi / 4)
	fullQuarter := NewZRotationMatrix(math.Pi / 2)

	// rotate the point around the z axis Pi/4 radians
	rotM, err := Multiply(halfQuarter, ptM)
	assert.NoError(t, err)
	rotMPoint, err := MatrixToPoint(rotM)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(-1*math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equals(rotMPoint))

	// rotate the point around the z axis Pi/2 radians
	rotM, err = Multiply(fullQuarter, ptM)
	assert.NoError(t, err)
	rotMPoint, err = MatrixToPoint(rotM)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(-1, 0, 0).Equals(rotMPoint))
}

func TestPointShearing(t *testing.T) {
	type args struct {
		transform *Matrix
		pt        *point.Point
	}
	tests := []struct {
		name string
		args args
		want *point.Point
	}{
		{
			name: "shearing transformation of x in proportion to y",
			args: args{
				transform: NewShearingMatrix(1, 0, 0, 0, 0, 0),
				pt:        point.NewPoint(2, 3, 4),
			},
			want: point.NewPoint(5, 3, 4),
		},
		{
			name: "shearing transformation of x in proportion to z",
			args: args{
				transform: NewShearingMatrix(0, 1, 0, 0, 0, 0),
				pt:        point.NewPoint(2, 3, 4),
			},
			want: point.NewPoint(6, 3, 4),
		},
		{
			name: "shearing transformation of y in proportion to x",
			args: args{
				transform: NewShearingMatrix(0, 0, 1, 0, 0, 0),
				pt:        point.NewPoint(2, 3, 4),
			},
			want: point.NewPoint(2, 5, 4),
		},
		{
			name: "shearing transformation of y in proportion to z",
			args: args{
				transform: NewShearingMatrix(0, 0, 0, 1, 0, 0),
				pt:        point.NewPoint(2, 3, 4),
			},
			want: point.NewPoint(2, 7, 4),
		},
		{
			name: "shearing transformation of z in proportion to x",
			args: args{
				transform: NewShearingMatrix(0, 0, 0, 0, 1, 0),
				pt:        point.NewPoint(2, 3, 4),
			},
			want: point.NewPoint(2, 3, 6),
		},
		{
			name: "shearing transformation of z in proportion to y",
			args: args{
				transform: NewShearingMatrix(0, 0, 0, 0, 0, 1),
				pt:        point.NewPoint(2, 3, 4),
			},
			want: point.NewPoint(2, 3, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mPt := PointToMatrix(tt.args.pt)
			shearM, err := Multiply(tt.args.transform, mPt)
			assert.NoError(t, err)

			shearPt, err := MatrixToPoint(shearM)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, shearPt)
		})
	}
}

func TestPointTransformationSequence(t *testing.T) {
	// Individual transformations are applied in sequence
	pt := point.NewPoint(1, 0, 1)
	a := NewXRotationMatrix(math.Pi / 2)
	b := NewScalingMatrix(5, 5, 5)
	c := NewTranslationMatrix(10, 5, 7)

	// Apply rotation first
	p2M, err := Multiply(a, PointToMatrix(pt))
	assert.NoError(t, err)
	p2, err := MatrixToPoint(p2M)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(1, -1, 0).Equals(p2))

	// Then Apply scaling
	p3M, err := Multiply(b, p2M)
	assert.NoError(t, err)
	p3, err := MatrixToPoint(p3M)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(5, -5, 0).Equals(p3))

	// Then apply translation
	p4M, err := Multiply(c, p3M)
	assert.NoError(t, err)
	p4, err := MatrixToPoint(p4M)
	assert.NoError(t, err)
	assert.True(t, point.NewPoint(15, 0, 7).Equals(p4))
}

func TestPointTransformationChain(t *testing.T) {
	pt := point.NewPoint(1, 0, 1)
	a := NewXRotationMatrix(math.Pi / 2)
	b := NewScalingMatrix(5, 5, 5)
	c := NewTranslationMatrix(10, 5, 7)

	m, err := Multiply(c, b)
	assert.NoError(t, err)

	tSeq, err := Multiply(m, a)
	assert.NoError(t, err)

	resM, err := Multiply(tSeq, PointToMatrix(pt))
	assert.NoError(t, err)
	resP, err := MatrixToPoint(resM)
	assert.NoError(t, err)

	assert.True(t, point.NewPoint(15, 0, 7).Equals(resP))
}

func TestPointTransformationFluent(t *testing.T) {
	transform := NewIdentityMatrix(4).
		RotateX(math.Pi/2).
		Scale(5, 5, 5).
		Translate(10, 5, 7)

	resM, err := Multiply(transform, PointToMatrix(point.NewPoint(1, 0, 1)))
	assert.NoError(t, err)
	resP, err := MatrixToPoint(resM)
	assert.NoError(t, err)

	assert.True(t, point.NewPoint(15, 0, 7).Equals(resP))
}

func TestTranslateVector(t *testing.T) {
	transform := NewTranslationMatrix(5, -3, 2)
	vec := vector.NewVector(-3, 4, 5)
	prod, err := Multiply(transform, VectorToMatrix(vec))
	assert.NoError(t, err)
	assert.NotNil(t, prod)

	// The vector must be the same after a translation happens to it.
	// Moving a vector around in space does not change the direction it points.
	mVec, err := MatrixToVector(prod)
	assert.NoError(t, err)
	assert.Equal(t, vec, mVec)
}

func TestVectorToMatrix(t *testing.T) {
	type args struct {
		vec *vector.Vector
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "vector to matrix conversion",
			args: args{
				vec: vector.NewVector(1, -2, 3),
			},
			want: NewMatrix(4, 1),
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
			err = tt.want.SetValue(3, 0, 0)
			assert.NoError(t, err)

			assert.Equal(t, tt.want, VectorToMatrix(tt.args.vec))
		})
	}
}

func TestToVector(t *testing.T) {
	type args struct {
		m *Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    *vector.Vector
		wantErr bool
	}{
		{
			name: "matrix to vector conversion",
			args: args{
				m: NewMatrix(4, 1),
			},
			want:    vector.NewVector(1, -2, 3),
			wantErr: false,
		},
		{
			name: "matrix to vector conversion error row length",
			args: args{
				m: NewMatrix(2, 1),
			},
			want:    vector.NewVector(1, -2, 3),
			wantErr: true,
		},
		{
			name: "matrix to vector conversion error col length",
			args: args{
				m: NewMatrix(4, 2),
			},
			want:    vector.NewVector(1, -2, 3),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				got, err := MatrixToVector(tt.args.m)
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				err := tt.args.m.SetValue(0, 0, 1)
				assert.NoError(t, err)
				err = tt.args.m.SetValue(1, 0, -2)
				assert.NoError(t, err)
				err = tt.args.m.SetValue(2, 0, 3)
				assert.NoError(t, err)

				got, err := MatrixToVector(tt.args.m)
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestScalingVector(t *testing.T) {
	transform := NewScalingMatrix(2, 3, 4)
	vec := vector.NewVector(-4, 6, 8)

	mult, err := Multiply(transform, VectorToMatrix(vec))
	assert.NoError(t, err)

	vecMult, err := MatrixToVector(mult)
	assert.NoError(t, err)

	assert.Equal(t, vector.NewVector(-8, 18, 32), vecMult)
}

func TestScalingInverseVector(t *testing.T) {
	transform := NewScalingMatrix(2, 3, 4)
	invM, err := Inverse(transform)
	assert.NoError(t, err)

	vec := vector.NewVector(-4, 6, 8)
	mult, err := Multiply(invM, VectorToMatrix(vec))
	assert.NoError(t, err)

	vecMult, err := MatrixToVector(mult)
	assert.NoError(t, err)

	assert.Equal(t, vector.NewVector(-2, 2, 2), vecMult)
}

func TestViewTransform(t *testing.T) {
	type args struct {
		from point.Point
		to   point.Point
		up   vector.Vector
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "default orientation view transformation matrix looking from the " +
				"origin in negative z direction",
			args: args{
				from: *point.NewPoint(0, 0, 0),
				to:   *point.NewPoint(0, 0, -1),
				up:   *vector.NewVector(0, 1, 0),
			},
			want: NewIdentityMatrix(4),
		},
		{
			name: "view transformation matrix looking from the origin in positive z direction",
			args: args{
				from: *point.NewPoint(0, 0, 0),
				to:   *point.NewPoint(0, 0, 1),
				up:   *vector.NewVector(0, 1, 0),
			},
			want: NewScalingMatrix(-1, 1, -1),
		},
		{
			name: "view transformation actually moves the world backward 8 units along the z axis",
			args: args{
				from: *point.NewPoint(0, 0, 8),
				to:   *point.NewPoint(0, 0, 0),
				up:   *vector.NewVector(0, 1, 0),
			},
			want: NewTranslationMatrix(0, 0, -8),
		},
		{
			name: "an arbitrary view transformation",
			args: args{
				from: *point.NewPoint(1, 3, 2),
				to:   *point.NewPoint(4, -2, 8),
				up:   *vector.NewVector(1, 1, 0),
			},
			want: &Matrix{
				rows: 4,
				cols: 4,
				data: []float64{
					-0.50709, 0.50709, 0.67612, -2.36643,
					0.76772, 0.60609, 0.12122, -2.82843,
					-0.35857, 0.59761, -0.71714, 0.00000,
					0.00000, 0.00000, 0.00000, 1.00000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vtM := ViewTransform(tt.args.from, tt.args.to, tt.args.up)
			if !assert.True(t, tt.want.Equals(vtM)) {
				assert.Equal(t, tt.want, vtM)
			}
		})
	}
}
