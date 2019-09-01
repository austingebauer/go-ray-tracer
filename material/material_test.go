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
				color:     color.NewColor(1, 1, 1),
				ambient:   0.1,
				diffuse:   0.9,
				specular:  0.9,
				shininess: 200.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewMaterial())
		})
	}
}
