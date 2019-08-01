package main

import (
	"fmt"
	"log"
	"os"

	"github.com/austingebauer/go-ray-tracer/canvas"
	"github.com/austingebauer/go-ray-tracer/color"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
)

const (
	projectilePPMFile      = "docs/renderings/clock/clock.ppm"
	projectileCanvasWidth  = 900
	projectileCanvasHeight = 600
	clockPPMFile           = "docs/renderings/clock/clock.ppm"
	clockCanvasWidth       = 400
	clockCanvasHeight      = 400
)

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

func main() {
	// RenderProjectile()
	RenderClock()
}

func RenderClock() {
	c := canvas.NewCanvas(clockCanvasWidth, clockCanvasHeight)
	white := color.NewColor(1, 1, 1)

	origin := point.NewPoint(clockCanvasWidth/2, clockCanvasHeight/2, 0)
	err := c.WritePixel(uint64(origin.X), uint64(origin.Y), *white)
	if err != nil {
		log.Fatal(err)
	}

	// Write the start to the canvas
	stopPt, err := point.ToPoint(point.ToMatrix(origin).Translate(100, 0, 0))
	if err != nil {
		log.Fatal(err)
	}
	err = c.WritePixel(uint64(stopPt.X), uint64(stopPt.Y), *white)
	if err != nil {
		log.Fatal(err)
	}

	// Rotate from the start by pi/6
	//currentPt, err := point.ToPoint(point.ToMatrix(stopPt).RotateZ(math.Pi / 6))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = c.WritePixel(uint64(currentPt.X), uint64(currentPt.Y), *white)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// While current point doesn't equal start start point, rotate pi/6 and write pixels
	//for !currentPt.Equals(stopPt) {
	//	var err error
	//	currentPt, err = point.ToPoint(point.ToMatrix(currentPt).RotateZ(math.Pi / 6))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	err = c.WritePixel(uint64(currentPt.X), uint64(currentPt.Y), *white)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	WriteCanvasToFile(c, clockPPMFile)
}

func RenderProjectile() {
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

	c := canvas.NewCanvas(projectileCanvasWidth, projectileCanvasHeight)

	// run tick repeatedly until the projectile's y position is less than or equal to 0
	tickCount := 0
	for proj.Position.Y >= 0 {
		fmt.Printf("Tick %v: Projectile position <X: %v, Y: %v, Z: %v> \n", tickCount,
			proj.Position.X, proj.Position.Y, proj.Position.Z)

		// write the position of the projectile to the canvas
		white := color.NewColor(1, 1, 1)
		err := c.WritePixel(uint64(proj.Position.X), uint64(projectileCanvasHeight-proj.Position.Y), *white)
		if err != nil {
			log.Fatal(err)
		}

		tick(env, proj)
		tickCount++
	}

	// Write the canvas to a PPM file
	WriteCanvasToFile(c, projectilePPMFile)
}

// tick moves the passed Projectile through the passed Environment.
func tick(env *Environment, proj *Projectile) Projectile {
	position := proj.Position.Add(*proj.Velocity)
	velocity := proj.Velocity.Add(*env.Gravity).Add(*env.Wind)
	return Projectile{
		Position: position,
		Velocity: velocity,
	}
}

func WriteCanvasToFile(c *canvas.Canvas, filePath string) {
	// Write the canvas to a PPM file
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = c.ToPPM(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote ppm rendering to: %v\n", filePath)
}
