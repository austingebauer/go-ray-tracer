// Package intersect encapsulates an intersection of a ray with an object.
package intersect

import (
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/austingebauer/go-ray-tracer/world"
	"math"
	"sort"
)

// Intersection encapsulates an intersection of a ray with an object.
type Intersection struct {
	// T represents the units +/- along a Ray where is intersected with Object.
	T float64
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

	//
	pt *point.Point

	//
	eyeVec *vector.Vector

	//
	normalVec *vector.Vector
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

	// Compute the point at which the ray intersected the sphere
	rayIntersectionPt := ray.Position(r, comps.Intersection.T)

	// Compute the eye vector
	eyeVec := vector.Scale(r.Direction, -1)

	// Compute the normal vector on the surface of the sphere at the intersection point
	normalVec, err := sphere.NormalAt(comps.Intersection.Object, rayIntersectionPt)
	if err != nil {
		return nil, err
	}

	comps.pt = rayIntersectionPt
	comps.eyeVec = eyeVec
	comps.normalVec = normalVec
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

// RaySphereIntersect intersects the passed ray with the passed sphere.
//
// It returns the t values (i.e., intersection units +/- away from the origin of the Ray)
// where the Ray intersects with the sphere.
//
// If the ray intersects with the sphere at two points, then two different intersection t values are returned.
// If the ray intersects with the sphere at a single, tangent point, then two equal t values are returned.
// If the ray does not intersect with the sphere, then an empty slice is returned.
func RaySphereIntersect(r *ray.Ray, s *sphere.Sphere) []*Intersection {
	// Details on calculation: https://en.wikipedia.org/wiki/Line%E2%80%93sphere_intersection

	// Transform the r by the inverse of the transformation associated with the s
	// in order to use unit s. Moving the r makes for more simple math and
	// same intersection results.
	sphereTransformInverse, _ := matrix.Inverse(s.Transform)
	transformedRay, _ := ray.Transform(r, sphereTransformInverse)

	// The vector from the s origin to the r origin.
	sphereToRayVec := point.Subtract(*transformedRay.Origin, *s.Origin)

	// Compute the discriminant to tell whether the r intersects with the s at all.
	a := vector.DotProduct(*transformedRay.Direction, *transformedRay.Direction)
	b := 2 * vector.DotProduct(*transformedRay.Direction, *sphereToRayVec)
	c := vector.DotProduct(*sphereToRayVec, *sphereToRayVec) - 1
	discriminant := math.Pow(b, 2) - 4*a*c

	// If the discriminant is negative, then the r misses the s and no intersections occur.
	if discriminant < 0 {
		return []*Intersection{}
	}

	// Compute the t values.
	t1 := ((-1 * b) - math.Sqrt(discriminant)) / (2 * a)
	t2 := ((-1 * b) + math.Sqrt(discriminant)) / (2 * a)

	// Return the intersection t values and object in increasing order
	return []*Intersection{
		{
			T:      t1,
			Object: s,
		},
		{
			T:      t2,
			Object: s,
		},
	}
}

// RayWorldIntersect intersects the passed ray with the passed world.
func RayWorldIntersect(r *ray.Ray, w *world.World) []*Intersection {
	allObjectIntersections := make([]*Intersection, 0)
	for _, sphereObj := range w.GetObjects() {
		intersections := RaySphereIntersect(r, sphereObj)
		allObjectIntersections = append(allObjectIntersections, intersections...)
	}

	// Sort the entire collection of intersections
	SortIntersectionsAsc(allObjectIntersections)

	return allObjectIntersections
}
