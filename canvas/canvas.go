// Package canvas represents a rectangular grid of pixels.
package canvas

import (
	"errors"
	"fmt"
)

// Canvas represents a rectangular grid of pixels.
type Canvas struct {
	width  uint64
	height uint64
	pixels [][]Color
}

// NewCanvas returns a new Canvas with the passed width and height.
func NewCanvas(width, height uint64) *Canvas {
	pixels := make([][]Color, height)
	for i := range pixels {
		pixels[i] = make([]Color, width)
	}

	return &Canvas{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

// WritePixel writes the passed Color to the Canvas at the pixel located at the passed x and y values.
func (c *Canvas) WritePixel(x uint64, y uint64, color Color) error {
	err := c.ValidateInCanvasBounds(x, y)
	if err != nil {
		return err
	}

	c.pixels[y][x] = color
	return nil
}

// PixelAt returns the Color at the pixel located at the passed x and y values.
func (c *Canvas) PixelAt(x, y uint64) (*Color, error) {
	err := c.ValidateInCanvasBounds(x, y)
	if err != nil {
		return nil, err
	}

	return &c.pixels[y][x], nil
}

// ValidateInCanvasBounds validates that the passed x and y values fit into the pixel bounds of the canvas.
func (c *Canvas) ValidateInCanvasBounds(x, y uint64) error {
	if y > c.height-1 {
		return errors.New(fmt.Sprintf("y value '%v' must be less than '%v'", y, c.height))
	}

	if x > c.width-1 {
		return errors.New(fmt.Sprintf("x value '%v' must be less than '%v'", x, c.width))
	}

	return nil
}
