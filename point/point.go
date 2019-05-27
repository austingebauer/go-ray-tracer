package point

import "github.com/austingebauer/go-ray-tracer/tuple"

type Point struct {
	*tuple.Tuple
}

// NewPoint returns a new Point that has the passed x, y, and z values.
func NewPoint(x, y, z float64) *Point {
	tpl, _ := tuple.NewTuple(x, y, z, tuple.Point)
	return &Point{
		Tuple: tpl,
	}
}

func (pt *Point) Equals(ptQ *Point) bool {
	return tuple.Equals(pt.Tuple, ptQ.Tuple)
}
