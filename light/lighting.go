// Package light provides difference types of light sources and behavior using light.
package light

import (
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/material"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
	"math"
)

// Lighting computes the shading for a material given the light source, point being illuminated,
// and the eye and normal vectors using the Phong reflection model.
func Lighting(m *material.Material, light *PointLight, pt *point.Point, eyeVec,
	normalVec *vector.Vector) *color.Color {
	// The three reflection contributions to get the final shading
	var ambient, diffuse, specular *color.Color

	// Combine the surface color with the light's color/intensity
	effectiveColor := color.Multiply(m.Color, light.Intensity)

	// Get the direction vector to the light source
	lightVec := point.Subtract(light.Position, *pt).Normalize()

	// Compute the ambient contribution
	ambient = color.Scale(*effectiveColor, m.Ambient)

	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector.
	lightDotNormal := vector.DotProduct(*lightVec, *normalVec)

	// A negative number means the light is on the other side of the surface.
	if lightDotNormal < 0 {
		// Only ambient light is present, so set diffuse and specular to black.
		diffuse = color.NewColor(0, 0, 0)
		specular = color.NewColor(0, 0, 0)

		return ambient.Add(*diffuse).Add(*specular)
	}

	// Compute the diffuse contribution
	diffuse = color.Scale(*color.Scale(*effectiveColor, m.Diffuse), lightDotNormal)

	// reflectDotEye represents the cosine of the angle between the
	// reflection vector and the eye vector.
	reflectVec := vector.Reflect(*lightVec.Negate(), *normalVec)
	reflectDotEye := vector.DotProduct(*reflectVec, *eyeVec)

	// A zero or negative number means the light reflects away from (not into) the eye.
	if reflectDotEye <= 0 {
		specular = color.NewColor(0, 0, 0)
		return ambient.Add(*diffuse).Add(*specular)
	}

	// Compute the specular contribution
	factor := math.Pow(reflectDotEye, m.Shininess)
	specular = color.Scale(*color.Scale(light.Intensity, m.Specular), factor)
	return ambient.Add(*diffuse).Add(*specular)
}
