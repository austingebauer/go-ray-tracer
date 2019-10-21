// Package material represents a material on the surface of an object.
package material

import "github.com/austingebauer/go-ray-tracer/color"

const (
	// DefaultAmbient is the ambient reflection factor of a default material.
	DefaultAmbient = 0.1
	// DefaultDiffuse is the diffuse reflection factor of a default material.
	DefaultDiffuse = 0.9
	// DefaultSpecular is the specular reflection factor of a default material.
	DefaultSpecular = 0.9
	// DefaultShininess is the shininess factor of a default material.
	DefaultShininess = 200.0
)

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

// NewDefaultMaterial returns a new Material with default values.
func NewDefaultMaterial() *Material {
	return NewMaterial(*color.NewColor(1, 1, 1),
		DefaultAmbient, DefaultDiffuse, DefaultSpecular, DefaultShininess)
}

// NewDefaultMaterial returns a new Material with default values.
func NewMaterial(c color.Color, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{
		Color:     c,
		Ambient:   ambient,
		Diffuse:   diffuse,
		Specular:  specular,
		Shininess: shininess,
	}
}
