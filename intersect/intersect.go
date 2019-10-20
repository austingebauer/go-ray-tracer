// Package intersect encapsulates an intersection of a ray with an object.
package intersect

import (
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
	"sort"
)

// Intersection encapsulates an intersection of a ray with an object.
type Intersection struct {
	// T represents the units +/- along a Ray where is intersected with Object.
	T float64
	// TODO: use interface for "Object" here instead of sphere
	// Object is the Sphere that was intersected by a Ray at T units.
	Object *sphere.Sphere
}

// IntersectionComputations encapsulates some precomputed information
// related to an intersection.
//
// Note:
//   This could potentially be in the intersection structure itself.
//   It's directly related to an intersection and carries its T and Object values.
type IntersectionComputations struct {
	Intersection

	// The Point at which the ray intersected the object
	Point *point.Point

	// The eye vector points in the opposite direction as the ray
	EyeVec *vector.Vector

	// The normal vector on the object surface at the Point of intersection
	NormalVec *vector.Vector

	// If Inside is true, the intersection occurred from Inside of the object.
	// Otherwise the intersection occurred from the outside of the object.
	Inside bool
}

// NewIntersection returns a new Intersection with the passed t value and object.
func NewIntersection(t float64, object sphere.Sphere) *Intersection {
	return &Intersection{
		T:      t,
		Object: &object,
	}
}

// PrepareComputations computes and returns additional information
// related to an intersection.
func PrepareComputations(i *Intersection, r *ray.Ray) (*IntersectionComputations, error) {
	comps := &IntersectionComputations{
		Intersection: *i,
	}

	// Compute the Point at which the ray intersected the sphere
	rayIntersectionPt := ray.Position(r, comps.Intersection.T)

	// Compute the eye vector
	eyeVec := vector.Scale(*r.Direction, -1)

	// Compute the normal vector on the surface of the sphere at the intersection Point
	normalVec, err := sphere.NormalAt(comps.Intersection.Object, rayIntersectionPt)
	if err != nil {
		return nil, err
	}

	// If the dot product of the normal vector and ray direction vector is negative,
	// then the intersection occurred from the Inside of the object. Otherwise,
	// the intersection occurred from the outside of the object.
	dotProduct := vector.DotProduct(*normalVec, *eyeVec)
	if dotProduct < 0 {
		comps.Inside = true
		normalVec.Negate()
	}

	comps.Point = rayIntersectionPt
	comps.EyeVec = eyeVec
	comps.NormalVec = normalVec
	return comps, nil
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
	SortIntersectionsAsc(intersections)

	// Search the intersections until a non-negative T value is found
	idx := 0
	for idx < len(intersections) && intersections[idx].T < 0 {
		idx++
	}

	// If the entire slice was searched, then we didn't find a positive T value.
	// No hit, return nil.
	if idx == len(intersections) {
		return nil
	}

	// idx is sitting at the Intersection with the lowest, non-negative T value
	return intersections[idx]
}

// SortIntersectionsAsc sorts the passed intersections into ascending order.
func SortIntersectionsAsc(intersections []*Intersection) {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})
}

// SortIntersectionsDesc sorts the passed intersections into descending order.
func SortIntersectionsDesc(intersections []*Intersection) {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T > intersections[j].T
	})
}
