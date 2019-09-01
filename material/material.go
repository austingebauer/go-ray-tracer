// Package material represents a material on the surface of an object.
package material

import "github.com/austingebauer/go-ray-tracer/color"

// Material represents a material on the surface of an object.
// It uses the Phong reflection model to simulate the reflection of light.
type Material struct {
	// The Color of the material. Default: white
	Color color.Color
	// The Ambient reflection of the material. Range is [0, 1]. Default is 0.1.
	Ambient float64
	// The Diffuse reflection of the material. Range is [0, 1]. Default is 0.9
	Diffuse float64
	// The Specular reflection of the material. Range is [0, 1]. Default is 0.9
	Specular float64
	// The Shininess of the material. Range is [10, 200]. Default is 200.0
	Shininess float64
}

// NewMaterial returns a new Material with default values.
func NewMaterial() *Material {
	return &Material{
		Color:     *color.NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
