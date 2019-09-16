package main

import (
	"fmt"
	"github.com/austingebauer/go-ray-tracer/intersection"
	"github.com/austingebauer/go-ray-tracer/light"
	"github.com/austingebauer/go-ray-tracer/material"
	"github.com/austingebauer/go-ray-tracer/ray"
	"github.com/austingebauer/go-ray-tracer/sphere"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/austingebauer/go-ray-tracer/canvas"
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

type Rendering struct {
	routine    func() *canvas.Canvas
	outputFile string
}

func main() {
	renderings := []Rendering{
		{
			routine:    RenderRayTracedSphere3D,
			outputFile: "docs/renderings/sphere_3d/sphere_3d.ppm",
		},
		{
			routine:    RenderRayTracedSphere2D,
			outputFile: "docs/renderings/sphere_2d/sphere_2d.ppm",
		},
		{
			routine:    RenderClock,
			outputFile: "docs/renderings/clock/clock.ppm",
		},
		{
			routine:    RenderProjectile,
			outputFile: "docs/renderings/projectile/projectile.ppm",
		},
	}

	// Add all renderings to the wait group
	var wg sync.WaitGroup
	wg.Add(len(renderings))

	// Start all renderings
	startTime := time.Now()
	for _, rendering := range renderings {
		go (func(r Rendering) {
			defer wg.Done()
			writeCanvasToFile(r.routine(), r.outputFile)
		})(rendering)
	}

	// Wait for all renderings to complete
	wg.Wait()

	// Log the elapsed time
	fmt.Printf("Render time: %v seconds\n\n", time.Now().Sub(startTime).Seconds())
}

// RenderRayTracedSphere3D renders a 3D ray traced sphere.
func RenderRayTracedSphere3D() *canvas.Canvas {
	// Create a light source
	lightPosition := point.NewPoint(-10, 10, -10)
	lightColor := color.NewColor(1, 1, 1)
	l := light.NewPointLight(*lightPosition, *lightColor)

	// Create a sphere and material for it
	s := sphere.NewUnitSphere("sphere_3d")
	mat := material.NewDefaultMaterial()
	mat.Color = *color.NewColor(0.7, 1, 1)
	s.Material = mat

	return renderSphere(s, l, true)
}

// RenderRayTracedSphere2D renders a 2D ray traced sphere.
func RenderRayTracedSphere2D() *canvas.Canvas {
	// Create a sphere and material for it
	s := sphere.NewUnitSphere("sphere_2d")
	mat := material.NewDefaultMaterial()
	mat.Color = *color.NewColor(0.7, 1, 1)
	s.Material = mat

	return renderSphere(s, nil, false)
}

// renderSphere renders the passed sphere onto the passed canvas using ray tracing.
func renderSphere(shape *sphere.Sphere, l *light.PointLight, render3D bool) *canvas.Canvas {
	c := canvas.NewCanvas(500, 500)

	// Pick an origin for the ray
	rayOrigin := point.NewPoint(0, 0, -5)

	// Pick a z value for the wall
	wallZ := 8.0

	// Pick the size of the wall based on extrapolating ray origin and sphere
	wallSize := 7.0

	// Half of the wall size when looking directly at the sphere
	halfWallSize := wallSize / 2.0

	// Divide the wall size by the number of canvas pixels to get
	// the size of a single pixel in world space units.
	pixelSize := wallSize / float64(c.Width)

	// For each row of pixels in the canvas
	for y := 0; y < c.Height; y++ {

		// Compute the world y coordinate (top = +half, bottom = -half)
		// 3.5 - 0.07 * (y = current row)
		worldY := halfWallSize - pixelSize*float64(y)

		// For each pixel in the row
		for x := 0; x < c.Width; x++ {

			// Compute the world x coordinate (left = -half, right = half)
			// -3.5 + 0.07 * (x = current pixel in row)
			worldX := (-1 * halfWallSize) + pixelSize*float64(x)

			// Describe the point on the wall that the Ray will target
			position := point.NewPoint(worldX, worldY, wallZ)

			// Create a ray from the ray origin to the position on the wall
			r := ray.NewRay(*rayOrigin, *vector.Normalize(*point.Subtract(*position, *rayOrigin)))

			// RaySphereIntersect the ray with the sphere
			xs := intersection.RaySphereIntersect(r, shape)

			// If there was a hit, write a pixel to the canvas
			hit := intersection.Hit(xs)
			if hit != nil {
				surfaceColor := hit.Object.Material.Color

				// Calculate the color at the surface using the shading function
				if render3D {
					pt := ray.Position(r, hit.T)
					normal, err := sphere.NormalAt(hit.Object, pt)
					if err != nil {
						log.Fatal(err)
					}
					eye := vector.Scale(r.Direction, -1)
					surfaceColor = *light.Lighting(hit.Object.Material, l, pt, eye, normal)
				}

				err := c.WritePixel(x, y, surfaceColor)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return c
}

// RenderClock renders a clock.
func RenderClock() *canvas.Canvas {
	// Orient the clock about the z-axis, such that the face of the clock
	// would be in the xy-plane while looking towards negative z-axis.

	// Set up the canvas
	c := canvas.NewCanvas(500, 500)
	var canvasRadius = float64(c.Width) / 4
	var canvasOriginWidth = float64(c.Width) / 2
	var canvasOriginHeight = float64(c.Height) / 2

	// Set some colors to render for different points
	white := color.NewColor(1, 1, 1)
	green := color.NewColor(0, 1, 0)

	// Set the origin
	origin := point.NewPoint(0, 0, 0)
	err := c.WritePixel(
		int(origin.X+canvasOriginWidth),
		int(origin.Y+canvasOriginHeight), *white)
	if err != nil {
		log.Fatal(err)
	}

	// Set twelve point on the clock
	twelve, err := point.ToPoint(point.ToMatrix(origin).Translate(0, 1, 0))
	if err != nil {
		log.Fatal(err)
	}
	err = c.WritePixel(
		int(canvasOriginWidth+(twelve.X*canvasRadius)),
		int(canvasOriginHeight-(twelve.Y*canvasRadius)),
		*green)
	if err != nil {
		log.Fatal(err)
	}

	// Set the next point to be rendered by a rotation about the z-axis
	next, err := point.ToPoint(point.ToMatrix(twelve).RotateZ(math.Pi / 6))
	if err != nil {
		log.Fatal(err)
	}

	// Rotate by pi/6 about the z-axis to render 1-12 o'clock
	for hour := 0; hour < 11; hour++ {

		// render next hour hand
		err = c.WritePixel(
			int(canvasOriginWidth+(next.X*canvasRadius)),
			int(canvasOriginHeight-(next.Y*canvasRadius)),
			*green)
		if err != nil {
			log.Fatal(err)
		}

		// rotate the next point pi/6
		next, err = point.ToPoint(point.ToMatrix(next).RotateZ(math.Pi / 6))
		if err != nil {
			log.Fatal(err)
		}
	}

	return c
}

// Projectile represents an object with a position and a velocity.
type Projectile struct {
	Position *point.Point
	Velocity *vector.Vector
}

// Environment represents a gravity and wind environment for Projectiles.
type Environment struct {
	Gravity *vector.Vector
	Wind    *vector.Vector
}

// RenderProjectile renders a projectile.
func RenderProjectile() *canvas.Canvas {
	c := canvas.NewCanvas(900, 600)

	// projectile starts one unit above the origin.
	start := point.NewPoint(0, 1, 0)

	// velocity is normalized to 1 unit/tick
	velocity := vector.NewVector(1, 1.8, 0).Normalize().Scale(11.25)

	// gravity -0.1 unit/tick
	gravity := vector.NewVector(0, -0.1, 0)

	// wind is 0.01 unit/tick
	wind := vector.NewVector(-0.01, 0, 0)

	proj := &Projectile{
		Position: start,
		Velocity: velocity,
	}

	env := &Environment{
		Gravity: gravity,
		Wind:    wind,
	}

	// run tick repeatedly until the projectile's y position is less than or equal to 0
	tickCount := 0
	for proj.Position.Y >= 0 {
		// Uncomment to view projectile x, y, and z values through ticks
		/*
			fmt.Printf("Tick %v: Projectile position <X: %v, Y: %v, Z: %v> \n", tickCount,
			proj.Position.X, proj.Position.Y, proj.Position.Z)
		*/

		// write the position of the projectile to the canvas
		white := color.NewColor(1, 1, 1)
		err := c.WritePixel(int(proj.Position.X), c.Height-int(proj.Position.Y), *white)
		if err != nil {
			log.Fatal(err)
		}

		tick(env, proj)
		tickCount++
	}

	return c
}

// tick moves the passed Projectile through the passed Environment.
func tick(env *Environment, proj *Projectile) Projectile {
	position := proj.Position.Add(proj.Velocity)
	velocity := proj.Velocity.Add(*env.Gravity).Add(*env.Wind)
	return Projectile{
		Position: position,
		Velocity: velocity,
	}
}

// writeCanvasToFile writes the passed canvas to a file at the passed path.
func writeCanvasToFile(c *canvas.Canvas, filePath string) {
	// Write the canvas to a PPM file
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = c.ToPPM(file, canvas.PixelMapTemplate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote ppm rendering to: %v\n\n", filePath)
}
