// Sphere represents a sphere object with an origin and radius.
package sphere

import (
	"math"
	"testing"

	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/stretchr/testify/assert"
)

func TestNewUnitSphere(t *testing.T) {
	tests := []struct {
		name string
		want *Sphere
	}{
		{
			name: "new unit sphere",
			want: &Sphere{
				Id:        "testID",
				Origin:    point.NewPoint(0, 0, 0),
				Radius:    1,
				Transform: matrix.NewIdentityMatrix(4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewUnitSphere("testID"), tt.want)
		})
	}
}

func TestNewSphere(t *testing.T) {
	type args struct {
		id     string
		origin *point.Point
		radius float64
	}
	tests := []struct {
		name string
		args args
		want *Sphere
	}{
		{
			name: "new sphere with origin and radius",
			args: args{
				id:     "testID",
				origin: point.NewPoint(1, 2, -3),
				radius: 11,
			},
			want: &Sphere{
				Id:        "testID",
				Origin:    point.NewPoint(1, 2, -3),
				Radius:    11,
				Transform: matrix.NewIdentityMatrix(4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewSphere(tt.args.id, tt.args.origin, tt.args.radius), tt.want)
		})
	}
}

func TestSphere_SetTransform(t *testing.T) {
	type args struct {
		m *matrix.Matrix
	}
	tests := []struct {
		name string
		s    *Sphere
		args args
		want *matrix.Matrix
	}{
		{
			name: "set sphere transform",
			s:    NewUnitSphere("testID"),
			args: args{
				m: matrix.NewMatrix(4, 4),
			},
			want: matrix.NewMatrix(4, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetTransform(tt.args.m)
			assert.Equal(t, tt.want, tt.s.Transform)
		})
	}
}

func TestNormalAtAreNormalized(t *testing.T) {
	type args struct {
		s *Sphere
		p *point.Point
	}
	tests := []struct {
		name string
		args args
		want *vector.Vector
	}{
		{
			name: "normal is normalized on a sphere at a point on the x axis",
			args: args{
				s: NewUnitSphere("testID"),
				p: point.NewPoint(1, 0, 0),
			},
			want: vector.NewVector(1, 0, 0),
		},
		{
			name: "normal is normalized on a sphere at a point on the y axis",
			args: args{
				s: NewUnitSphere("testID"),
				p: point.NewPoint(0, 1, 0),
			},
			want: vector.NewVector(0, 1, 0),
		},
		{
			name: "normal is normalized on a sphere at a point on the z axis",
			args: args{
				s: NewUnitSphere("testID"),
				p: point.NewPoint(0, 0, 1),
			},
			want: vector.NewVector(0, 0, 1),
		},
		{
			name: "normal is normalized on a sphere at a nonaxial point",
			args: args{
				s: NewUnitSphere("testID"),
				p: point.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			},
			want: vector.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Normal vector is what we expect
			normalVector, err := NormalAt(tt.args.s, tt.args.p)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, normalVector)

			// Normal vector is normalized
			normalizedNormal := vector.Normalize(*normalVector)
			assert.Equal(t, normalVector, normalizedNormal)
			assert.Equal(t, tt.want, normalizedNormal)
		})
	}
}

func TestNormalAtOnTransformedSphere(t *testing.T) {
	type args struct {
		s *Sphere
		p *point.Point
		transforms []*matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want *vector.Vector
	}{
		{
			name: "normal is normalized on a translated sphere",
			args: args{
				s: NewUnitSphere("testID"),
				p: point.NewPoint(1, 0, 0),
				transforms: []*matrix.Matrix{
					matrix.NewTranslationMatrix(0,1,0),
				},
			},
			want: vector.NewVector(0, 0.70711, -0.70711),
		},
		//{
		//	name: "normal is normalized on a transformed sphere",
		//	args: args{
		//		s: NewUnitSphere("testID"),
		//		p: point.NewPoint(1, 0, 0),
		//		transforms: []*matrix.Matrix{
		//			matrix.NewScalingMatrix(1,0.5,1),
		//			matrix.NewZRotationMatrix(math.Pi / 5),
		//		},
		//	},
		//	want: vector.NewVector(0, 0.97014, -0.24254),
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the transform to apply it to the sphere
			var err error
			transform := matrix.NewIdentityMatrix(4)
			for _, sphereTransform := range tt.args.transforms {
				transform, err = matrix.Multiply(transform, sphereTransform)
				assert.NoError(t, err)
			}
			tt.args.s.SetTransform(transform)

			// Normal vector is what we expect
			normalVector, err := NormalAt(tt.args.s, tt.args.p)
			assert.NoError(t, err)
			assert.True(t, tt.want.Equals(normalVector))

			// Normal vector is normalized
			normalizedNormal := vector.Normalize(*normalVector)
			assert.True(t, normalVector.Equals(normalizedNormal))
			assert.True(t, tt.want.Equals(normalizedNormal))
		})
	}
}
