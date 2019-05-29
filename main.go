package main

import (
	"fmt"

	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
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
	proj := &Projectile{
		// projectile starts one unit above the origin.
		Position: point.NewPoint(0, 1, 0),
		// velocity is normalized to 1 unit/tick.
		Velocity: vector.NewVector(1, 1, 0).Normalize(),
	}

	env := &Environment{
		// gravity -0.1 unit/tick
		Gravity: vector.NewVector(0, -0.1, 0),
		// wind is 0.01 unit/tick
		Wind: vector.NewVector(-0.01, 0, 0),
	}

	// run tick repeatedly until the projectile's y position is less than or equal to 0
	tickCount := 0
	for proj.Position.Y >= 0 {
		fmt.Printf("Tick %v: Projectile position <X: %v, Y: %v, Z: %v> \n", tickCount,
			proj.Position.X, proj.Position.Y, proj.Position.Z)
		tick(env, proj)
		tickCount++
	}
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
