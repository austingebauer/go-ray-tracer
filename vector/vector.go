package vector

import "github.com/austingebauer/go-ray-tracer/tuple"

type Vector struct {
	tuple.Tuple
}

// NewVector returns a new Vector that has the passed x, y, and z values.
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		tuple.Tuple{
			X: x,
			Y: y,
			Z: z,
			W: tuple.VectorW,
		},
	}
}
