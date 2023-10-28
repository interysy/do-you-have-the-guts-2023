package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
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
)

func main() {
	loginScreen()
}

func loginScreen() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//var texture rl.Texture2D = rl.LoadTexture("assets/texture.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		// rl.DrawTexture(texture, rl.LoadImageFromScreen().Width/2, rl.LoadImageFromScreen().Height/2, rl.White)
		rl.ClearBackground(PURPLE1)
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


func generateParticles(num_particles int, x int32, y int32) []rl.Rectangle{
	var deviation int32 = 10
	particles := make([]rl.Rectangle, num_particles)

	for i := 0; i < num_particles; i++ {
		particle = rl.Rectangle{}
		particle.X = float32(x + rand.Int31n(deviation))
		particle.Y = float32(y + rand.Int31n(deviation))
		particle.Width = 25 
		particle.Height = 25
	}

	return particles
}

func updateParticles(particles []rl.Rectangle) []rl.Rectangle{
	var gravity int32 = 10

	for _, particle := range particles {
		particle.X += rand.Int31n(deviation)
		particle.Y += gravity // might need to be -
	}
  
	return particles
}

func renderParticles(particles []rl.Rectangle) {
	var colour = rl.Gray

	//TODO Add random to colour
	for _, particle := range particles {
		rl.DrawRectanglerec(particle, colour)
	}
}
