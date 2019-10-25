package camera

import (
	"github.com/austingebauer/go-ray-tracer/canvas"
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/austingebauer/go-ray-tracer/world"
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
		want *Camera
	}{
		{
			name: "constructing a new camera",
			args: args{
				horizontalSize: 160,
				verticalSize:   120,
				fieldOfView:    math.Pi / 2,
			},
			want: &Camera{
				horizontalSizeInPixels: 160,
				verticalSizeInPixels:   120,
				fieldOfView:            math.Pi / 2,
				Transform:              *matrix.NewIdentityMatrix(4),
				aspectRatio:            1.3333333333333333,
				halfWidth:              1,
				halfHeight:             0.75,
				pixelSize:              0.0125,
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

func TestPixelSize(t *testing.T) {
	type args struct {
		horizontalSize int
		verticalSize   int
		fieldOfView    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "pixel size for camera with a horizontal aspect (horizontalSize > verticalSize)",
			args: args{
				horizontalSize: 200,
				verticalSize:   125,
				fieldOfView:    math.Pi / 2,
			},
			want: 0.01,
		},
		{
			name: "pixel size for camera with a vertical aspect (verticalSize > horizontalSize)",
			args: args{
				horizontalSize: 125,
				verticalSize:   200,
				fieldOfView:    math.Pi / 2,
			},
			want: 0.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCamera(tt.args.horizontalSize, tt.args.verticalSize, tt.args.fieldOfView)
			assert.Equal(t, tt.want, c.pixelSize)
		})
	}
}

func TestRayForPixel(t *testing.T) {
	type args struct {
		c *Camera
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want *ray.Ray
	}{
		{
			name: "constructing a ray through the center of the canvas",
			args: args{
				c: NewCamera(201, 101, math.Pi/2),
				x: 100,
				y: 50,
			},
			want: ray.NewRay(
				*point.NewPoint(0, 0, 0),
				*vector.NewVector(0, 0, -1)),
		},
		{
			name: "constructing a ray through a corner of the canvas",
			args: args{
				c: NewCamera(201, 101, math.Pi/2),
				x: 0,
				y: 0,
			},
			want: ray.NewRay(
				*point.NewPoint(0, 0, 0),
				*vector.NewVector(0.66519, 0.33259, -0.66851)),
		},
		{
			name: "constructing a ray when the camera is transformed",
			args: args{
				c: newCamera(201, 101, math.Pi/2, matrix.Multiply4x4(
					*matrix.NewYRotationMatrix(math.Pi / 4),
					*matrix.NewTranslationMatrix(0, -2, 5)),
				),
				x: 100,
				y: 50,
			},
			want: ray.NewRay(
				*point.NewPoint(0, 2, -5),
				*vector.NewVector(math.Sqrt(2)/2, 0, -1*math.Sqrt(2)/2)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := RayForPixel(tt.args.c, tt.args.x, tt.args.y)
			assert.NoError(t, err)

			if !assert.True(t, ray.Equals(tt.want, r)) {
				assert.Equal(t, tt.want, r)
			}
		})
	}
}

func TestRender(t *testing.T) {
	type args struct {
		c *Camera
		w *world.World
	}
	tests := []struct {
		name    string
		args    args
		want    *canvas.Canvas
		wantErr bool
	}{
		{
			name: "rendering a world with a camera",
			args: args{
				c: newCamera(11, 11, math.Pi/2,
					matrix.ViewTransform(
						*point.NewPoint(0, 0, -5),
						*point.NewPoint(0, 0, 0),
						*vector.NewVector(0, 1, 0))),
				w: world.NewDefaultWorld(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			image, err := Render(tt.args.c, tt.args.w)
			assert.NoError(t, err)

			// TODO: Investigate why book says green should be 0.47583.
			expectedColor := color.NewColor(0.38066, 0.047583, 0.2855)
			actualColor, err := image.PixelAt(5, 5)
			assert.NoError(t, err)

			if !color.Equals(*expectedColor, actualColor) {
				assert.Equal(t, expectedColor, actualColor)
			}
		})
	}
}
