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

// Add modifies this Color by adding each of the passed Color's rgb values to this Color's rgb values.
func (c *Color) Add(c2 Color) *Color {
	c.Red += c2.Red
	c.Green += c2.Green
	c.Blue += c2.Blue
	return c
}
