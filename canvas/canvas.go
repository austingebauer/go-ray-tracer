// Package canvas represents a rectangular grid of Pixels.
package canvas

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
)

// ppmTemplate is a template used for rendering a Canvas to a portable pixmap (PPM) file.
const ppmTemplate = `{{ .PPMIdentifier }}
{{ .Width }} {{ .Height }}
{{ .MaxColorValue }}
{{ pixels .Pixels }}`

const (
	ppmID         = "P3"
	ppmFileName   = "pixels.ppm"
	maxColorValue = 255
	newLineChar   = "\n"
)

// Canvas represents a rectangular grid of Pixels.
type Canvas struct {
	Width         uint64
	Height        uint64
	Pixels        [][]Color
	PPMIdentifier string
	MaxColorValue uint8
}

// NewCanvas returns a new Canvas with the passed Width and Height.
func NewCanvas(width, height uint64) *Canvas {
	// TODO: validate width and height

	pixels := make([][]Color, height)
	for i := range pixels {
		pixels[i] = make([]Color, width)
	}

	return &Canvas{
		Width:         width,
		Height:        height,
		Pixels:        pixels,
		PPMIdentifier: ppmID,
		MaxColorValue: maxColorValue,
	}
}

// WritePixel writes the passed Color to the Canvas at the pixel
// located at the passed x and y values.
func (c *Canvas) WritePixel(x uint64, y uint64, color Color) error {
	err := c.ValidateInCanvasBounds(x, y)
	if err != nil {
		return err
	}

	c.Pixels[y][x] = color
	return nil
}

// PixelAt returns the Color at the pixel located at the passed x and y values.
func (c *Canvas) PixelAt(x, y uint64) (*Color, error) {
	err := c.ValidateInCanvasBounds(x, y)
	if err != nil {
		return nil, err
	}

	return &c.Pixels[y][x], nil
}

// ValidateInCanvasBounds validates that the passed x and y values
// fit into the pixel bounds of the canvas.
func (c *Canvas) ValidateInCanvasBounds(x, y uint64) error {
	if y > c.Height-1 {
		return errors.New(fmt.Sprintf("y value '%v' must be less than '%v'", y, c.Height))
	}

	if x > c.Width-1 {
		return errors.New(fmt.Sprintf("x value '%v' must be less than '%v'", x, c.Width))
	}

	return nil
}

// ToPPM writes the current canvas to a file in the portable pixmap (PPM) format.
func (c *Canvas) ToPPM(writer io.Writer) error {
	funcMap := template.FuncMap{
		"pixels": writePPMPixels,
	}

	tmpl, err := template.New(ppmID).Funcs(funcMap).Parse(ppmTemplate)
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
func writePPMPixels(pixels [][]Color) string {
	pixelBytes := bytes.Buffer{}
	for _, row := range pixels {
		for _, color := range row {
			pixelBytes.WriteString(fmt.Sprintf("%v %v %v%v",
				color.Red*maxColorValue,
				color.Green*maxColorValue,
				color.Blue*maxColorValue,
				newLineChar))
		}
	}

	return pixelBytes.String()
}
