package point

import "github.com/austingebauer/go-ray-tracer/tuple"

type Point struct {
	tuple.Tuple
}

// NewPoint returns a new Point that has the passed x, y, and z values.
func NewPoint(x, y, z float64) *Point {
	return &Point{
		tuple.Tuple{
			X: x,
			Y: y,
			Z: z,
			W: tuple.PointW,
		},
	}
}
