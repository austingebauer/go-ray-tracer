// Sphere represents a sphere object with an origin and radius.
package sphere

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/austingebauer/go-ray-tracer/point"
)

func TestNewUnitSphere(t *testing.T) {
	tests := []struct {
		name string
		want *Sphere
	}{
		{
			name: "new unit sphere",
			want: &Sphere{
				id:     "testID",
				Origin: point.NewPoint(0, 0, 0),
				Radius: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnitSphere()
			assert.Equal(t, tt.want.Origin, got.Origin)
			assert.Equal(t, tt.want.Radius, got.Radius)
		})
	}
}

func TestNewSphere(t *testing.T) {
	type args struct {
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
				origin: point.NewPoint(1, 2, -3),
				radius: 11,
			},
			want: &Sphere{
				id:     "testID",
				Origin: point.NewPoint(1, 2, -3),
				Radius: 11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSphere(tt.args.origin, tt.args.radius)
			assert.Equal(t, tt.want.Origin, got.Origin)
			assert.Equal(t, tt.want.Radius, got.Radius)
		})
	}
}
