// Package light represents difference types of light sources.
package light

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/point"
)

func TestNewPointLight(t *testing.T) {
	type args struct {
		position  point.Point
		intensity color.Color
	}
	tests := []struct {
		name string
		args args
		want *PointLight
	}{
		{
			name: "point light has a position and intensity",
			args: args{
				position:  *point.NewPoint(0, 0, 0),
				intensity: *color.NewColor(1, 1, 1),
			},
			want: NewPointLight(*point.NewPoint(0, 0, 0), *color.NewColor(1, 1, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPointLight(tt.args.position, tt.args.intensity))
		})
	}
}
