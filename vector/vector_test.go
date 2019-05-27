package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVector(t *testing.T) {
	vec := NewVector(1, -2, 3.14)
	assert.Equal(t, float64(1), vec.X)
	assert.Equal(t, float64(-2), vec.Y)
	assert.Equal(t, float64(3.14), vec.Z)
	assert.Equal(t, float64(0.0), vec.W)
}
