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
// eye and normal vectors, and shadow flag using the Phong reflection model.
func Lighting(mat *material.Material, light *PointLight, pt *point.Point, eyeVec,
	normalVec *vector.Vector, inShadow bool) *color.Color {
	// The three reflection contributions to get the final shading
	var ambient, diffuse, specular *color.Color

	// Combine the surface color with the light's color/intensity
	effectiveColor := color.Multiply(mat.Color, light.Intensity)

	// Get the direction vector to the light source
	lightVec := vector.Normalize(*point.Subtract(light.Position, *pt))

	// Compute the ambient contribution
	ambient = color.Scale(*effectiveColor, mat.Ambient)

	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector.
	lightDotNormal := vector.DotProduct(*lightVec, *normalVec)

	// A negative number means the light is on the other side of the surface.
	if lightDotNormal < 0 {
		// Only ambient light is present, so set diffuse and specular to black.
		diffuse = color.NewColor(0, 0, 0)
		specular = color.NewColor(0, 0, 0)
	} else {
		// Compute the diffuse contribution
		diffuse = color.Scale(*color.Scale(*effectiveColor, mat.Diffuse), lightDotNormal)

		// reflectDotEye represents the cosine of the angle between the
		// reflection vector and the eye vector.
		reflectVec := vector.Reflect(*vector.Scale(*lightVec, -1), *normalVec)
		reflectDotEye := vector.DotProduct(*reflectVec, *eyeVec)

		// A zero or negative number means the light reflects away from (not into) the eye.
		if reflectDotEye <= 0 {
			specular = color.NewColor(0, 0, 0)
		} else {
			// Compute the specular contribution
			factor := math.Pow(reflectDotEye, mat.Shininess)
			specular = color.Scale(*color.Scale(light.Intensity, mat.Specular), factor)
		}
	}

	finalColor := color.Add(color.Add(*ambient, *diffuse), *specular)
	return &finalColor
}
