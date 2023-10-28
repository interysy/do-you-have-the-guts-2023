package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)
type Particle struct {
	x int32
	y int32
}

func generateParticles(num_particles int, x int32, y int32) []rl.Rectangle {
	var deviation int32 = 10
	particles := make([]rl.Rectangle, num_particles)

	for i := 0; i < num_particles; i++ {
		particles[i].X = float32(x + rand.Int31n(deviation))
		particles[i].Y = float32(y + rand.Int31n(deviation))
		particles[i].Width = 25
		particles[i].Height = 25
	}

	return particles
}

func updateParticles(particles []rl.Rectangle, centralXCoordinate int32) {
	var deviation int32 = 10
	var gravity int32 = 1 // Adjust the gravity value for a noticeable effect

	for i := range particles {
		var min float32 = -1.0
		var max float32 = 1.0
		r := min + rand.Float32()*(max-min)
		particles[i].X += float32(centralXCoordinate) + (r * float32(rand.Int31n(deviation)))

		particles[i].Y -= float32(gravity)
	}
}

func renderParticles(particles []rl.Rectangle) {
	// Loop through particles and draw each one
	for _, particle := range particles {
		// Generate a random color for each particle
		color := rl.LightGray
		rl.DrawRectangleRec(particle, color)
	}
}
