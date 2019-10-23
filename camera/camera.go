// Package camera provides a virtual camera that allows for taking pictures of a scene.
package camera

import "github.com/austingebauer/go-ray-tracer/matrix"

type Camera struct {
	horizontalSize int
	verticalSize   int
	fieldOfView    float64
	transform      matrix.Matrix
}

func NewCamera() *Camera {
	return &Camera{}
}
