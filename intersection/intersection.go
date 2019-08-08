// Package intersection encapsulates an intersection of a ray with an object.
package intersection

import (
	"github.com/austingebauer/go-ray-tracer/sphere"
	"sort"
)

// Intersection encapsulates an intersection of a ray with an object.
type Intersection struct {
	// T represents the units +/- along a Ray where is intersected with Object.
	T float64
	// Object is the Sphere that was intersected by a Ray at T units.
	Object *sphere.Sphere
}

// NewIntersection returns a new Intersection with the passed t value and object.
func NewIntersection(t float64, object *sphere.Sphere) *Intersection {
	return &Intersection{
		T:      t,
		Object: object,
	}
}

// Intersections returns a slice of the passed Intersections.
func Intersections(intersections ...*Intersection) []*Intersection {
	return intersections
}

// Hit returns the Intersection with the lowest non-negative T value.
func Hit(intersections []*Intersection) *Intersection {
	if len(intersections) == 0 {
		return nil
	}

	// Sort the intersections ascending from index 0 on T.
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})

	// Search the intersections until a non-negative T value is found
	idx := 0
	for idx < len(intersections) && intersections[idx].T < 0 {
		idx++
	}

	// If the entire slice was searched, then we didn't find a positive T value
	// No hit, return nil.
	if idx == len(intersections) {
		return nil
	}

	// idx is sitting at the Intersection with the lowest, non-negative T value
	return intersections[idx]
}
