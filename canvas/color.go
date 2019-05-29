package canvas

type Color struct {
	Red, Green, Blue float64
}

// NewColor returns a new Color that has the passed red, green, and blue values.
func NewColor(red, green, blue float64) *Color {
	return &Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}
