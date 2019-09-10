// Package material represents a material on the surface of an object.
package material

import (
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMaterial(t *testing.T) {
	tests := []struct {
		name string
		want *Material
	}{
		{
			name: "default material has values",
			want: &Material{
				Color:     *color.NewColor(1, 1, 1),
				Ambient:   0.1,
				Diffuse:   0.9,
				Specular:  0.9,
				Shininess: 200.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewDefaultMaterial())
		})
	}
}
