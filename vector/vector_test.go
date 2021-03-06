package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "NewPoint sets x, y, and z",
			args: args{
				x: float64(1),
				y: float64(-2),
				z: float64(0),
			},
			want: &Vector{
				1,
				-2,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewVector(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func TestVector_Equals(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vecQ *Vector
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
				&Vector{
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
				&Vector{
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
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Equals(tt.args.vecQ))
		})
	}
}

func TestVector_Magnitude(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "vector has magnitude",
			fields: fields{
				2,
				2,
				1,
			},
			want: 3,
		},
		{
			name: "negative component vector has same magnitude",
			fields: fields{
				-2,
				-2,
				-1,
			},
			want: 3,
		},
		{
			name: "unit vector x",
			fields: fields{
				1,
				0,
				0,
			},
			want: 1,
		},
		{
			name: "unit vector y",
			fields: fields{
				0,
				1,
				0,
			},
			want: 1,
		},
		{
			name: "unit vector z",
			fields: fields{
				0,
				0,
				1,
			},
			want: 1,
		},
		{
			name: "vector magnitude equals square root",
			fields: fields{
				1,
				2,
				3,
			},
			want: math.Sqrt(14),
		},
		{
			name: "negative vector magnitude equals square root",
			fields: fields{
				-1,
				-2,
				-3,
			},
			want: math.Sqrt(14),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Magnitude())
		})
	}
}

func TestVector_Negate(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector
	}{
		{
			name: "negate a vector positive",
			fields: fields{
				1,
				2,
				3,
			},
			want: &Vector{
				-1,
				-2,
				-3,
			},
		},
		{
			name: "negate a vector negative",
			fields: fields{
				-1,
				-2,
				-3,
			},
			want: &Vector{
				1,
				2,
				3,
			},
		},
		{
			name: "negate a zero vector",
			fields: fields{
				0,
				0,
				0,
			},
			want: &Vector{
				0,
				0,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Negate())
		})
	}
}

func TestVector_Scale(t *testing.T) {
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
		want   *Vector
	}{
		{
			name: "scale the vector",
			fields: fields{
				0,
				2,
				-3,
			},
			args: args{
				scalar: float64(2),
			},
			want: &Vector{
				0,
				4,
				-6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Scale(tt.args.scalar))
		})
	}
}

func TestVector_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec2 Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector
	}{
		{
			name: "vector add vector method",
			fields: fields{
				0,
				-1,
				1,
			},
			args: args{
				Vector{
					2,
					3,
					0,
				},
			},
			want: &Vector{
				2,
				2,
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Add(tt.args.vec2))
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		vec1 Vector
		vec2 Vector
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			name: "vector add vector function",
			args: args{
				vec1: Vector{
					0,
					-1,
					1,
				},
				vec2: Vector{
					2,
					3,
					0,
				},
			},
			want: Vector{
				2,
				2,
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Add(tt.args.vec1, tt.args.vec2))
		})
	}
}

func TestVector_Subtract(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	type args struct {
		vec2 Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector
	}{
		{
			name: "vector subtract vector method",
			fields: fields{
				0,
				-1,
				2,
			},
			args: args{
				Vector{
					0,
					2,
					3,
				},
			},
			want: &Vector{
				0,
				-3,
				-1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			assert.Equal(t, tt.want, vec.Subtract(tt.args.vec2))
		})
	}
}

func TestVector_Normalize(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector
	}{
		{
			name: "normalize vector method 1",
			fields: fields{
				X: 4,
				Y: 0,
				Z: 0,
			},
			want: &Vector{
				X: 1,
				Y: 0,
				Z: 0,
			},
		},
		{
			name: "normalize vector method 2",
			fields: fields{
				X: 1,
				Y: 2,
				Z: 3,
			},
			want: &Vector{
				X: 1 / math.Sqrt(14),
				Y: 2 / math.Sqrt(14),
				Z: 3 / math.Sqrt(14),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			vec.Normalize()
			assert.Equal(t, tt.want, vec)
			assert.Equal(t, float64(1), vec.Magnitude(),
				"magnitude of normalized vector is always 1")
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		vec Vector
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "normalize vector function 1",
			args: args{
				Vector{
					X: 4,
					Y: 0,
					Z: 0,
				},
			},
			want: &Vector{
				X: 1,
				Y: 0,
				Z: 0,
			},
		},
		{
			name: "normalize vector function 2",
			args: args{
				Vector{
					X: 1,
					Y: 2,
					Z: 3,
				},
			},
			want: &Vector{
				X: 1 / math.Sqrt(14),
				Y: 2 / math.Sqrt(14),
				Z: 3 / math.Sqrt(14),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalizedVec := Normalize(tt.args.vec)
			assert.Equal(t, tt.want, normalizedVec)
			assert.Equal(t, float64(1), normalizedVec.Magnitude())
		})
	}
}

func TestDotProduct(t *testing.T) {
	type args struct {
		vec1 Vector
		vec2 Vector
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "dot product of two vectors",
			args: args{
				vec1: Vector{
					X: 1,
					Y: 2,
					Z: 3,
				},
				vec2: Vector{
					X: 2,
					Y: 3,
					Z: 4,
				},
			},
			want: 20,
		},
		{
			name: "dot product of two identical unit vectors",
			args: args{
				vec1: Vector{
					X: 0,
					Y: 0,
					Z: 1,
				},
				vec2: Vector{
					X: 0,
					Y: 0,
					Z: 1,
				},
			},
			want: 1,
		},
		{
			name: "dot product of two identical vectors",
			args: args{
				vec1: Vector{
					X: 1,
					Y: 1,
					Z: 1,
				},
				vec2: Vector{
					X: 1,
					Y: 1,
					Z: 1,
				},
			},
			want: 3,
		},
		{
			name: "dot product of two vectors with 90 degree angle",
			args: args{
				vec1: Vector{
					X: 0,
					Y: 2,
					Z: 0,
				},
				vec2: Vector{
					X: 2,
					Y: 0,
					Z: 0,
				},
			},
			want: 0,
		},
		{
			name: "dot product of two unit vectors with 180 degree (opposite direction) angle",
			args: args{
				vec1: Vector{
					X: 1,
					Y: 0,
					Z: 0,
				},
				vec2: Vector{
					X: -1,
					Y: 0,
					Z: 0,
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DotProduct(tt.args.vec1, tt.args.vec2))
		})
	}
}

func TestCrossProduct(t *testing.T) {
	type args struct {
		vec1 Vector
		vec2 Vector
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			name: "cross product of two unit vectors",
			args: args{
				vec1: Vector{
					X: 1,
					Y: 0,
					Z: 0,
				},
				vec2: Vector{
					X: 0,
					Y: 1,
					Z: 0,
				},
			},
			want: Vector{
				X: 0,
				Y: 0,
				Z: 1,
			},
		},
		{
			name: "cross product of two unit vectors",
			args: args{
				vec1: Vector{
					X: 0,
					Y: 1,
					Z: 0,
				},
				vec2: Vector{
					X: 1,
					Y: 0,
					Z: 0,
				},
			},
			want: Vector{
				X: 0,
				Y: 0,
				Z: -1,
			},
		},
		{
			name: "cross product of two vectors",
			args: args{
				vec1: Vector{
					X: 1,
					Y: 2,
					Z: 3,
				},
				vec2: Vector{
					X: 2,
					Y: 3,
					Z: 4,
				},
			},
			want: Vector{
				X: -1,
				Y: 2,
				Z: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			perpendicularVec := CrossProduct(tt.args.vec1, tt.args.vec2)
			assert.Equal(t, tt.want, perpendicularVec)
			assert.Equal(t, float64(0), DotProduct(tt.args.vec1, perpendicularVec),
				"dot product of vec1 and cross product vector must be 0 (90 degree angle)")
			assert.Equal(t, float64(0), DotProduct(tt.args.vec2, perpendicularVec),
				"dot product of vec2 and cross product vector must be 0 (90 degree angle)")
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		vec1 Vector
		vec2 Vector
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "vector subtract vector function",
			args: args{
				vec1: Vector{
					0,
					-1,
					1,
				},
				vec2: Vector{
					2,
					3,
					0,
				},
			},
			want: &Vector{
				-2,
				-4,
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Subtract(tt.args.vec1, tt.args.vec2))
		})
	}
}

func TestReflect(t *testing.T) {
	type args struct {
		in     Vector
		normal Vector
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "reflecting a vector approaching at 45 degrees from +x, -y",
			args: args{
				in:     *NewVector(1, -1, 0),
				normal: *NewVector(0, 1, 0),
			},
			want: NewVector(1, 1, 0),
		},
		{
			name: "reflecting a vector approaching at 45 degrees from +x, +y with +y normal",
			args: args{
				in:     *NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
				normal: *NewVector(0, 1, 0),
			},
			want: NewVector(math.Sqrt(2)/2, -1*math.Sqrt(2)/2, 0),
		},
		{
			name: "reflecting a vector approaching at 45 degrees from +x, +y with -y normal",
			args: args{
				in:     *NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
				normal: *NewVector(0, -1, 0),
			},
			want: NewVector(math.Sqrt(2)/2, -1*math.Sqrt(2)/2, 0),
		},
		{
			name: "reflecting a vector off of a slanted surface",
			args: args{
				in:     *NewVector(0, -1, 0),
				normal: *NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
			},
			want: NewVector(1, 0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflectionVec := Reflect(tt.args.in, tt.args.normal)
			vectorsEqual := reflectionVec.Equals(tt.want)

			// If the vectors are not equal, log the difference
			if !vectorsEqual {
				assert.Equal(t, tt.want, reflectionVec)
			}
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		vec    *Vector
		scalar float64
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{
			name: "scaling a vector returns new scaled vector",
			args: args{
				vec:    NewVector(1, 0, -1),
				scalar: 3,
			},
			want: NewVector(3, 0, -3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Scale(*tt.args.vec, tt.args.scalar))
		})
	}
}
