// Package canvas represents a rectangular grid of Pixels.
package canvas

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/austingebauer/go-ray-tracer/color"
	"html/template"
	"io"
	"math"
)

// PixelMapTemplate is a template used for rendering a Canvas to a portable pixmap (PPM) file.
const PixelMapTemplate = `{{ .PPMIdentifier }}
{{ .Width }} {{ .Height }}
{{ .MaxColorValue }}
{{ pixels .Pixels }}`

const (
	ppmID         = "P3"
	maxColorValue = 255
	minColorValue = 0
	newLineChar   = "\n"
)

// Canvas represents a rectangular grid of Pixels.
type Canvas struct {
	Width         int
	Height        int
	Pixels        [][]*color.Color
	PPMIdentifier string
	MaxColorValue uint8
	MinColorValue uint8
}

// NewCanvas returns a new Canvas with the passed Width and Height.
func NewCanvas(width, height int) *Canvas {
	pixels := make([][]*color.Color, height)
	for i := range pixels {
		pixels[i] = make([]*color.Color, width)
	}

	// Set the canvas default color for each pixel to black
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixels[y][x] = color.NewColor(0, 0, 0)
		}
	}

	return &Canvas{
		Width:         width,
		Height:        height,
		Pixels:        pixels,
		PPMIdentifier: ppmID,
		MaxColorValue: maxColorValue,
		MinColorValue: minColorValue,
	}
}

// WritePixel writes the passed Color to the Canvas at the pixel
// located at the passed x and y values.
func (c *Canvas) WritePixel(x, y int, color *color.Color) error {
	err := c.ValidateInCanvasBounds(x, y)
	if err != nil {
		return err
	}

	c.Pixels[y][x] = color
	return nil
}

// PixelAt returns the Color at the pixel located at the passed x and y values.
func (c *Canvas) PixelAt(x, y int) (*color.Color, error) {
	err := c.ValidateInCanvasBounds(x, y)
	if err != nil {
		return nil, err
	}

	return c.Pixels[y][x], nil
}

// ValidateInCanvasBounds validates that the passed x and y values
// fit into the pixel bounds of the canvas.
func (c *Canvas) ValidateInCanvasBounds(x, y int) error {
	if x < 0 {
		return fmt.Errorf("x value '%v' must be greater than zero", x)
	}

	if y < 0 {
		return fmt.Errorf("y value '%v' must be greater than zero", x)
	}

	if y > c.Height-1 {
		return fmt.Errorf("y value '%v' must be less than '%v'", y, c.Height)
	}

	if x > c.Width-1 {
		return fmt.Errorf("x value '%v' must be less than '%v'", x, c.Width)
	}

	return nil
}

// ToPPM writes the current canvas to a file in the portable pixmap (PPM) format.
func (c *Canvas) ToPPM(writer io.Writer, goTemplate string) error {
	if writer == nil {
		return errors.New("writer must not be nil")
	}

	funcMap := template.FuncMap{
		"pixels": writePPMPixels,
	}

	tmpl, err := template.New(ppmID).Funcs(funcMap).Parse(goTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(writer, c)
	if err != nil {
		return err
	}

	return nil
}

// writePPMPixels returns a string containing rows of pixels with rgb values.
func writePPMPixels(pixels [][]*color.Color) string {
	pixelBytes := bytes.Buffer{}
	for _, row := range pixels {
		for _, colorVal := range row {

			// Convert RGB color values to 8 bit integer values [0-255]
			redEightBit := int(math.Min(math.Max(minColorValue, colorVal.Red*maxColorValue),
				maxColorValue))
			greenEightBit := int(math.Min(math.Max(minColorValue, colorVal.Green*maxColorValue),
				maxColorValue))
			blueEightBit := int(math.Min(math.Max(minColorValue, colorVal.Blue*maxColorValue),
				maxColorValue))

			pixelBytes.WriteString(fmt.Sprintf("%d %d %d%s",
				redEightBit,
				greenEightBit,
				blueEightBit,
				newLineChar))
		}
	}

	return pixelBytes.String()
}
