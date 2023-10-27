package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	loginScreen()
}

func loginScreen() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		var colour rl.Color = rl.NewColor(uint8(60), uint8(9), uint8(108), uint8(1))

		rl.ClearBackground(colour)
		rl.DrawRectangleRounded(rl.Rectangle{400, 225, 20, 20}, 0.1, 10, rl.Black)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
