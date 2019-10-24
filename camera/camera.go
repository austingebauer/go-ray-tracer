// Package camera provides a virtual camera that allows for taking pictures of a scene.
package camera

import (
	"github.com/austingebauer/go-ray-tracer/matrix"
	"math"
)

// camera is a virtual camera that can be moved around,
// zoomed in and out, and transformed around a scene.
type camera struct {
	// The horizontal size in pixels
	horizontalSizeInPixels int
	// The vertical size in pixels
	verticalSizeInPixels int
	// An angle that describes how much the camera can see
	fieldOfView float64
	// A matrix describing how the world should be oriented relative to the camera
	transform matrix.Matrix
	// The ratio of the horizontal size of the canvas to its vertical size
	aspectRatio float64
	// Half of the width of the canvas
	halfWidth float64
	// Half of the height of the canvas
	halfHeight float64
	// The size, in world-space units, of the pixels on the canvas
	pixelSize float64
}

// NewCamera returns a new camera having the passed horizontal
// and vertical size in pixels, and field of view angle.
func NewCamera(horizontalSize int, verticalSize int, fieldOfView float64) *camera {
	c := &camera{
		horizontalSizeInPixels: horizontalSize,
		verticalSizeInPixels:   verticalSize,
		fieldOfView:            fieldOfView,
		transform:              *matrix.NewIdentityMatrix(4),
	}

	// Prepare the camera with important world space units
	c.prepareWorldSpaceUnits()

	return c
}

// prepareWorldSpaceUnits sets attributes on this camera related to world space units.
// It sets the cameras pixel size, half its width and height, and the aspect ratio.
func (c *camera) prepareWorldSpaceUnits() {
	// Compute the width of half of the canvas by taking the tangent of half of the field of view.
	// Cutting the field of view in half creates a right triangle on the canvas, which is 1 unit
	// away from the camera. The adjacent is 1 and the opposite is half of the canvas.
	halfView := math.Tan(c.fieldOfView / 2)

	// Compute the aspect ratio
	c.aspectRatio = float64(c.horizontalSizeInPixels) / float64(c.verticalSizeInPixels)

	// Compute half of the width and half of the height of the canvas.
	// This is different than the number of horizontal or vertical pixels.
	if c.aspectRatio >= 1 {
		// The horizontal size is greater than or equal to the vertical size
		c.halfWidth = halfView
		c.halfHeight = halfView / c.aspectRatio
	} else {
		// The vertical size is greater than the horizontal size
		c.halfWidth = halfView * c.aspectRatio
		c.halfHeight = halfView
	}

	// Divide half of the width * 2 by the number of horizontal pixels to get
	// the pixel size. Note that the assumption here is that the pixels are
	// square, so there is no need to compute the vertical size of the pixel.
	c.pixelSize = (c.halfWidth * 2) / float64(c.horizontalSizeInPixels)
}
