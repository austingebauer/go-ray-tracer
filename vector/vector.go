package vector

import "github.com/austingebauer/go-ray-tracer/tuple"

type Vector struct {
	*tuple.Tuple
}

// NewVector returns a new Vector that has the passed x, y, and z values.
func NewVector(x, y, z float64) *Vector {
	tpl, _ := tuple.NewTuple(x, y, z, tuple.Vector)
	return &Vector{
		Tuple: tpl,
	}
}

func (vec *Vector) Equals(vecQ *Vector) bool {
	return tuple.Equals(vec.Tuple, vecQ.Tuple)
}
