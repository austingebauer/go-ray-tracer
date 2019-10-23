// Package camera provides a virtual camera that allows for taking pictures of a scene.
package camera

import "github.com/austingebauer/go-ray-tracer/matrix"

// Camera is a virtual camera that can be moved around,
// zoomed in and out, and transformed around a scene.
type Camera struct {
	// The horizontal size in pixels
	horizontalSize int
	// The vertical size in pixels
	verticalSize int
	// An angle that describes how much the camera can see
	fieldOfView float64
	// A matrix describing how the world should be oriented relative to the camera
	transform matrix.Matrix
}

// NewCamera returns a new Camera having the passed horizontal
// and vertical size in pixels, and field of view angle.
func NewCamera(horizontalSize int, verticalSize int, fieldOfView float64) *Camera {
	return &Camera{
		horizontalSize: horizontalSize,
		verticalSize:   verticalSize,
		fieldOfView:    fieldOfView,
		transform:      *matrix.NewIdentityMatrix(4),
	}
}
