// Package camera provides a virtual camera that allows for taking pictures of a scene.
package camera

import (
	"github.com/austingebauer/go-ray-tracer/canvas"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/austingebauer/go-ray-tracer/world"
	"math"
)

// Camera is a virtual Camera that can be moved around,
// zoomed in and out, and transformed around a scene.
type Camera struct {
	// The horizontal size in pixels
	horizontalSizeInPixels int
	// The vertical size in pixels
	verticalSizeInPixels int
	// An angle that describes how much the camera can see
	fieldOfView float64
	// A matrix describing how the world should be moved/oriented relative to the camera
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
func NewCamera(horizontalSize int, verticalSize int, fieldOfView float64) *Camera {
	return newCamera(horizontalSize, verticalSize, fieldOfView, matrix.NewIdentityMatrix(4))
}

func newCamera(horizontalSize int, verticalSize int, fieldOfView float64,
	transform *matrix.Matrix) *Camera {

	c := &Camera{
		horizontalSizeInPixels: horizontalSize,
		verticalSizeInPixels:   verticalSize,
		fieldOfView:            fieldOfView,
		transform:              *transform,
	}
	c.prepareWorldSpaceUnits()

	return c
}

// prepareWorldSpaceUnits sets attributes on this camera related to world space units.
// It sets the cameras pixel size, half its width and height, and the aspect ratio.
func (c *Camera) prepareWorldSpaceUnits() {
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

// RayForPixel returns a new ray that starts at the passed camera
// and passes through the indicated (x, y) pixel on the canvas.
func RayForPixel(c *Camera, px int, py int) (*ray.Ray, error) {
	// Compute the offset from the left edge of the canvas to the pixel's center
	xOffset := c.pixelSize * (float64(px) + 0.5)
	YOffset := c.pixelSize * (float64(py) + 0.5)

	// The untransformed coordinates of the pixel in world space.
	// Note that the camera looks toward -z, so +x is to the left.
	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - YOffset

	// Using the camera matrix, transform the canvas point and
	// the origin, and then compute the ray's direction vector.
	// Note that the canvas is at z=-1
	inverseTransform, err := matrix.Inverse(c.transform)
	if err != nil {
		return nil, err
	}

	pixelM := matrix.Multiply4x4(*inverseTransform,
		*matrix.PointToMatrix(point.NewPoint(worldX, worldY, -1)))
	originM := matrix.Multiply4x4(*inverseTransform,
		*matrix.PointToMatrix(point.NewPoint(0, 0, 0)))

	pixelPt, err := matrix.MatrixToPoint(pixelM)
	if err != nil {
		return nil, err
	}
	originPt, err := matrix.MatrixToPoint(originM)
	if err != nil {
		return nil, err
	}

	directionVec := vector.Normalize(*point.Subtract(*pixelPt, *originPt))
	return ray.NewRay(*originPt, *directionVec), nil
}

// Render uses the passed camera to render the passed world into a canvas.
func Render(c *Camera, w *world.World) (*canvas.Canvas, error) {
	image := canvas.NewCanvas(c.horizontalSizeInPixels, c.verticalSizeInPixels)

	// For each pixel of the camera
	for y := 0; y < c.verticalSizeInPixels; y++ {
		for x := 0; x < c.horizontalSizeInPixels; x++ {
			// Compute the ray for the current pixel
			r, err := RayForPixel(c, x, y)
			if err != nil {
				return nil, err
			}

			// Intersect the ray with the world to get the color at the intersection
			c, err := world.ColorAt(w, r)
			if err != nil {
				return nil, err
			}

			// Write the color to the canvas at the current pixel
			err = image.WritePixel(x, y, *c)
			if err != nil {
				return nil, err
			}
		}
	}

	return image, nil
}
