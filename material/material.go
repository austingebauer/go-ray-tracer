// Package material represents a material on the surface of an object.
package material

import "github.com/austingebauer/go-ray-tracer/color"

// Material represents a material on the surface of an object.
type Material struct {
	// The color of the material. Default: white
	color *color.Color
	// The ambient reflection of the material. Default: 0.1
	ambient float64
	// The diffuse reflection of the material. Default: 0.9
	diffuse float64
	// The specular reflection of the material. Default: 0.9
	specular float64
	// The shininess of the material. Default: 200.0
	shininess float64
}

// NewMaterial returns a new Material with default values.
func NewMaterial() *Material {
	return &Material{
		color:     color.NewColor(1, 1, 1),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}
