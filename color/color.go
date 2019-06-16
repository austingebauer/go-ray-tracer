package color

// A Color is represented by percentages of red, green, and blue in the range of 0-1.
// The red, green, or blue values of a color may be lower or higher than 0-1.
type Color struct {
	Red   float32
	Green float32
	Blue  float32
}

// NewColor returns a new Color that has the passed red, green, and blue values.
func NewColor(red, green, blue float32) *Color {
	return &Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

// Add modifies this Color by adding each of the passed Color's rgb values
// to this Color's rgb values.
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

// Add modifies this Color by adding each of the passed Color's rgb values
// to this Color's rgb values.
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
func (c *Color) Scale(scalar float32) *Color {
	c.Red = c.Red * scalar
	c.Green = c.Green * scalar
	c.Blue = c.Blue * scalar
	return c
}

// Multiply modifies this Color by multiplying each of the passed Color's rgb values
// against this Color's rgb values.
func (c *Color) Multiply(c1 Color) *Color {
	c.Red = c.Red * c1.Red
	c.Green = c.Green * c1.Green
	c.Blue = c.Blue * c1.Blue
	return c
}

// Multiply returns a new Color with rgb values set to the
// product of the passed Color rgb values.
func Multiply(c1, c2 Color) *Color {
	return &Color{
		Red:   c1.Red * c2.Red,
		Green: c1.Green * c2.Green,
		Blue:  c1.Blue * c2.Blue,
	}
}
