package ray

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

func TestNewRay(t *testing.T) {
	type args struct {
		origin    *point.Point
		direction *vector.Vector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ray has origin and direction",
			args: args{
				origin:    point.NewPoint(1, 2, 3),
				direction: vector.NewVector(4, 5, 6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRay(tt.args.origin, tt.args.direction)
			assert.Equal(t, tt.args.origin, r.Origin)
			assert.Equal(t, tt.args.direction, r.Direction)
		})
	}
}
