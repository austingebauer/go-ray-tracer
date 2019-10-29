// Package world represents a collection of all Objects that make up a scene.
package world

import (
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/light"
	"github.com/austingebauer/go-ray-tracer/material"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/sphere"
)

// World represents a collection of all Objects that make up a scene.
type World struct {
	Objects []*sphere.Sphere
	Light   *light.PointLight
}

// NewWorld returns a new World.
func NewWorld() *World {
	return &World{
		Objects: make([]*sphere.Sphere, 0),
		Light:   nil,
	}
}

// NewDefaultWorld returns a new default world, which contains
// two spheres and a point light source.
func NewDefaultWorld() *World {
	// Create a default light source
	defaultLight := light.NewPointLight(
		*point.NewPoint(-10, 10, -10),
		*color.NewColor(1, 1, 1))

	// Create a default sphere number 1
	s1 := sphere.NewUnitSphere("s1")
	s1.Material = material.NewMaterial(*color.NewColor(0.8, 0.1, 0.6),
		material.DefaultAmbient, 0.7, 0.2, material.DefaultShininess)

	// Create a default sphere number 2
	s2 := sphere.NewUnitSphere("s2")
	s2.Transform = matrix.NewScalingMatrix(0.5, 0.5, 0.5)

	return &World{
		Objects: []*sphere.Sphere{s1, s2},
		Light:   defaultLight,
	}
}

// RayWorldIntersect intersects the passed ray with the passed world.
func RayWorldIntersect(r *ray.Ray, w *World) []*ray.Intersection {
	allObjectIntersections := make([]*ray.Intersection, 0)
	for _, sphereObj := range w.Objects {
		intersections := ray.RaySphereIntersect(r, sphereObj)
		allObjectIntersections = append(allObjectIntersections, intersections...)
	}

	// Sort the entire collection of intersections
	ray.SortIntersectionsAsc(allObjectIntersections)

	return allObjectIntersections
}

// ColorAt intersects the given ray with the given world and
// returns the color at the resulting intersection.
func ColorAt(w *World, r *ray.Ray) (*color.Color, error) {
	intersections := RayWorldIntersect(r, w)
	hit := ray.Hit(intersections)
	if hit == nil {
		return color.NewColor(0, 0, 0), nil
	}

	comps, err := ray.PrepareComputations(hit, r)
	if err != nil {
		return nil, err
	}

	return ShadeHit(w, comps), nil
}

// ShadeHit returns the color at the intersection encapsulated by
// an intersections computations.
func ShadeHit(w *World, comps *ray.IntersectionComputations) *color.Color {
	return light.Lighting(
		comps.Object.Material,
		w.Light,
		comps.Point,
		comps.EyeVec,
		comps.NormalVec,
		false)
}

// IsShadowed returns true if the passed point lies in
// the shadow of an object in the passed world.
func IsShadowed(world *World, pt *point.Point) bool {
	return false
}
