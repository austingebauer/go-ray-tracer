// Package light provides difference types of light sources and behavior using light.
package light

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"

	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/material"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

func TestLighting(t *testing.T) {
	type args struct {
		eyeVec    *vector.Vector
		normalVec *vector.Vector
		l         *PointLight
		m         *material.Material
		pt        *point.Point
	}
	tests := []struct {
		name string
		args args
		want *color.Color
	}{
		{
			name: "lighting with the eye between the light and the surface",
			args: args{
				eyeVec:    vector.NewVector(0, 0, -1),
				normalVec: vector.NewVector(0, 0, -1),
				l: NewPointLight(
					*point.NewPoint(0, 0, -10),
					*color.NewColor(1, 1, 1),
				),

				// material and point illuminated constant for this test table
				m:  material.NewMaterial(),
				pt: point.NewPoint(0, 0, 0),
			},
			want: color.NewColor(1.9, 1.9, 1.9),
		},
		{
			name: "lighting with the eye between the light and the surface, eye offset 45 degrees",
			args: args{
				eyeVec: vector.NewVector(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2),
				l: NewPointLight(
					*point.NewPoint(0, 0, -10),
					*color.NewColor(1, 1, 1),
				),

				// material, normal, and point illuminated constant for this test table
				normalVec: vector.NewVector(0, 0, -1),
				m:         material.NewMaterial(),
				pt:        point.NewPoint(0, 0, 0),
			},
			want: color.NewColor(1.0, 1.0, 1.0),
		},
		{
			name: "lighting with eye opposite surface, light offset 45 degrees",
			args: args{
				eyeVec: vector.NewVector(0, 0, -1),
				l: NewPointLight(
					*point.NewPoint(0, 10, -10),
					*color.NewColor(1, 1, 1),
				),

				// material, normal, and point illuminated constant for this test table
				normalVec: vector.NewVector(0, 0, -1),
				m:         material.NewMaterial(),
				pt:        point.NewPoint(0, 0, 0),
			},
			want: color.NewColor(0.7364, 0.7364, 0.7364),
		},
		{
			name: "lighting with eye and light offset 45 degrees for zero specular reflection",
			args: args{
				eyeVec: vector.NewVector(0, 1, 1),
				l: NewPointLight(
					*point.NewPoint(0, 10, -10),
					*color.NewColor(1, 1, 1),
				),

				// material, normal, and point illuminated constant for this test table
				normalVec: vector.NewVector(0, 0, -1),
				m:         material.NewMaterial(),
				pt:        point.NewPoint(0, 0, 0),
			},
			want: color.NewColor(0.73639, 0.73639, 0.73639),
		},
		{
			name: "lighting with eye in the path of the reflection vector",
			args: args{
				eyeVec: vector.NewVector(0, -1*math.Sqrt(2)/2, -1*math.Sqrt(2)/2),
				l: NewPointLight(
					*point.NewPoint(0, 10, -10),
					*color.NewColor(1, 1, 1),
				),

				// material, normal, and point illuminated constant for this test table
				normalVec: vector.NewVector(0, 0, -1),
				m:         material.NewMaterial(),
				pt:        point.NewPoint(0, 0, 0),
			},
			want: color.NewColor(1.6364, 1.6364, 1.6364),
		},
		{
			name: "lighting with the light behind the surface",
			args: args{
				eyeVec: vector.NewVector(0, 0, -1),
				l: NewPointLight(
					*point.NewPoint(0, 0, 10),
					*color.NewColor(1, 1, 1),
				),

				// material, normal, and point illuminated constant for this test table
				normalVec: vector.NewVector(0, 0, -1),
				m:         material.NewMaterial(),
				pt:        point.NewPoint(0, 0, 0),
			},
			want: color.NewColor(0.1, 0.1, 0.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := Lighting(tt.args.m, tt.args.l, tt.args.pt, tt.args.eyeVec, tt.args.normalVec)

			if !color.Equals(*lc, *tt.want) {
				assert.Equal(t, tt.want, lc)
			}
		})
	}
}
