package camera

import (
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewCamera1(t *testing.T) {
	type args struct {
		horizontalSize int
		verticalSize   int
		fieldOfView    float64
	}
	tests := []struct {
		name string
		args args
		want *camera
	}{
		{
			name: "constructing a new camera",
			args: args{
				horizontalSize: 160,
				verticalSize:   120,
				fieldOfView:    math.Pi / 2,
			},
			want: &camera{
				horizontalSize: 160,
				verticalSize:   120,
				fieldOfView:    math.Pi / 2,
				transform:      *matrix.NewIdentityMatrix(4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want,
				NewCamera(tt.args.horizontalSize, tt.args.verticalSize, tt.args.fieldOfView))
		})
	}
}
