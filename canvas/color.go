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
	c.Red = c.Red + c2.Red
	c.Green = c.Green + c2.Green
	c.Blue = c.Blue + c2.Blue
	return c
}

// Add returns a new Color with rgb values set to the sum of the passed Color rgb values.
func Add(c1, c2 Color) Color {
	return Color{
		Red:   c1.Red + c2.Red,
		Green: c1.Green + c2.Green,
		Blue:  c1.Blue + c2.Blue,
	}
}

// Add modifies this Color by adding each of the passed Color's rgb values to this Color's rgb values.
func (c *Color) Subtract(c2 Color) *Color {
	c.Red = c.Red - c2.Red
	c.Green = c.Green - c2.Green
	c.Blue = c.Blue - c2.Blue
	return c
}

// Add returns a new Color with rgb values set to the sum of the passed Color rgb values.
func Subtract(c1, c2 Color) Color {
	return Color{
		Red:   c1.Red - c2.Red,
		Green: c1.Green - c2.Green,
		Blue:  c1.Blue - c2.Blue,
	}
}

// Scale multiplies each of this Color's rgb values by the passed scalar value.
func (c *Color) Scale(scalar float64) *Color {
	c.Red = c.Red * scalar
	c.Green = c.Green * scalar
	c.Blue = c.Blue * scalar
	return c
}
