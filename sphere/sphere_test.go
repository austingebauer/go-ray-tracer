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
				Id:     "testID",
				Origin: point.NewPoint(0, 0, 0),
				Radius: 1,
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewSphere(tt.args.id, tt.args.origin, tt.args.radius), tt.want)
		})
	}
}
