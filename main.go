package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	ORANGE1 = rl.NewColor(255, 109, 0, 1)
	ORANGE2 = rl.NewColor(255, 121, 0, 1)
	ORANGE3 = rl.NewColor(255, 133, 0, 1)
	ORANGE4 = rl.NewColor(255, 145, 0, 1)
	ORANGE5 = rl.NewColor(255, 158, 0, 1)
	PURPLE1 = rl.NewColor(36, 0, 70, 1)
	PURPLE2 = rl.NewColor(60, 9, 108, 1)
	PURPLE3 = rl.NewColor(90, 24, 154, 1)
	PURPLE4 = rl.NewColor(123, 44, 191, 1)
	PURPLE5 = rl.NewColor(157, 78, 221, 1)

	CENTRAL = rl.GetScreenWidth() / 2
)

func main() {
	loginScreen()
}

func loginScreen() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//var texture rl.Texture2D = rl.LoadTexture("assets/texture.png")

	var particles = generateParticles(10, centraliseInX(25), 100)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// rl.DrawTexture(texture, rl.LoadImageFromScreen().Width/2, rl.LoadImageFromScreen().Height/2, rl.White)
		rl.ClearBackground(PURPLE1)

		//rl.DrawRectangleRounded(rl.NewRectangle(centraliseInX(800), 50, 400, 350), 0.5, 0, ORANGE1)
		updateParticles(particles, int32(CENTRAL))
		renderParticles(particles)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

func centraliseInX(size int) int32 {
	var centralXCoordinate = rl.GetScreenWidth()/2 - size/2
	return int32(centralXCoordinate)
}

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
