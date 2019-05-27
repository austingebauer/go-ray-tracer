package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	pt := NewPoint(1, -2, 3.14)
	assert.Equal(t, float64(1), pt.X)
	assert.Equal(t, float64(-2), pt.Y)
	assert.Equal(t, float64(3.14), pt.Z)
	assert.Equal(t, float64(1.0), pt.W)
}
