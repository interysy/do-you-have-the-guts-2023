package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	orange1 = rl.NewColor(255, 109, 0, 1)
	orange2 = rl.NewColor(255, 121, 0, 1)
	orange3 = rl.NewColor(255, 133, 0, 1)
	orange4 = rl.NewColor(255, 145, 0, 1)
	orange5 = rl.NewColor(255, 158, 0, 1)
	purple1 = rl.NewColor(36, 0, 70, 1)
	purple2 = rl.NewColor(60, 9, 108, 1)
	purple3 = rl.NewColor(90, 24, 154, 1)
	purple4 = rl.NewColor(123, 44, 191, 1)
	purple5 = rl.NewColor(157, 78, 221, 1)
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
		rl.ClearBackground(purple1)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

func centraliseInX(size int) int32 {
	var centralXCoordinate = rl.GetScreenWidth()/2 - size/2

	return int32(centralXCoordinate)
}
