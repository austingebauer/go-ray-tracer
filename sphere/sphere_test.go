// Sphere represents a sphere object with an origin and radius.
package sphere

import (
	"testing"

	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
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
				Id:     "testID",
				Origin: point.NewPoint(0, 0, 0),
				Radius: 1,
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
				Id:     "testID",
				Origin: point.NewPoint(1, 2, -3),
				Radius: 11,
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
			s: NewUnitSphere("testID"),
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
