package ray

import (
	"github.com/austingebauer/go-ray-tracer/maths"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"github.com/austingebauer/go-ray-tracer/vector"
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

	// The Point at which the ray intersected the object
	Point *point.Point

	// The OverPoint is the Point that has been slightly adjusted in
	// the direction of the NormalVec in order to avoid shadow acne from
	// self-intersection when determining if an intersection is in shadow.
	OverPoint *point.Point

	// The eye vector points in the opposite direction as the ray
	EyeVec *vector.Vector

	// The normal vector on the object surface at the Point of intersection
	NormalVec *vector.Vector

	// If Inside is true, the intersection occurred from Inside of the object.
	// Otherwise the intersection occurred from the outside of the object.
	Inside bool
}

// NewIntersection returns a new Intersection with the passed t value and object.
func NewIntersection(t float64, object *sphere.Sphere) *Intersection {
	return &Intersection{
		T:      t,
		Object: object,
	}
}

// PrepareComputations computes and returns additional information related to an intersection.
func PrepareComputations(i *Intersection, r *Ray) (*IntersectionComputations, error) {
	comps := &IntersectionComputations{
		Intersection: *i,
	}

	// Compute the Point at which the ray intersected the sphere
	rayIntersectionPt := Position(r, comps.Intersection.T)

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

	// Compute the over point in order to avoid rendering shadow acne
	// caused by the shadow ray intersecting with the object itself.
	comps.OverPoint = point.Add(comps.Point, vector.Scale(*comps.NormalVec, maths.Epsilon))

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
// If the ray intersects with the sphere at a single, tangent Point, then two equal t values are returned.
// If the ray does not intersect with the sphere, then an empty slice is returned.
func RaySphereIntersect(r *Ray, s *sphere.Sphere) []*Intersection {
	// Details on calculation: https://en.wikipedia.org/wiki/Line%E2%80%93sphere_intersection

	// transform the r by the inverse of the transformation associated with the s
	// in order to use unit s. Moving the r makes for more simple math and
	// same intersection results.
	sphereTransformInverse, _ := matrix.Inverse(s.Transform)
	transformedRay, _ := Transform(r, sphereTransformInverse)

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
