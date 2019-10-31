// Package light provides light sources and shading.
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

	// Combine the surface color with the light's color/intensity
	effectiveColor := color.Multiply(mat.Color, light.Intensity)

	// The three reflection contributions to get the final shading
	ambient := color.Scale(*effectiveColor, mat.Ambient)
	diffuse := color.NewColor(0, 0, 0)
	specular := color.NewColor(0, 0, 0)

	// If the point is in a shadow, use only the ambient contribution.
	// The diffuse and specular contribution
	if inShadow {
		return ambient
	}

	// Get the vector from the point to the light source
	lightVec := vector.Normalize(*point.Subtract(light.Position, *pt))

	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector.
	lightDotNormal := vector.DotProduct(*lightVec, *normalVec)

	// A negative number means the light is on the other side of the surface
	// and only ambient light is present
	if lightDotNormal < 0 {
		return ambient
	}

	// Compute the diffuse contribution
	diffuse = color.Scale(*color.Scale(*effectiveColor, mat.Diffuse), lightDotNormal)

	// reflectDotEye represents the cosine of the angle between the
	// reflection vector and the eye vector.
	reflectVec := vector.Reflect(*vector.Scale(*lightVec, -1), *normalVec)
	reflectDotEye := vector.DotProduct(*reflectVec, *eyeVec)

	// A positive number means that the light reflects into the eye.
	// A zero or negative number means the light reflects away from (not into) the eye.
	if reflectDotEye > 0 {
		// Compute the specular contribution
		factor := math.Pow(reflectDotEye, mat.Shininess)
		specular = color.Scale(*color.Scale(light.Intensity, mat.Specular), factor)
	}

	return color.Add(*color.Add(*ambient, *diffuse), *specular)
}
