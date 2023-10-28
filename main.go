package main

import (
	"strconv"

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
	var state string = "login"
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if state == "login" {
			loginScreen()
		}
		rl.EndDrawing()
	}
}

func centraliseInX(size int) int32 {
	var centralXCoordinate = rl.GetScreenWidth()/2 - size/2
	return int32(centralXCoordinate)
}

func DrawBorderedRectangle(rect rl.Rectangle, borderWidth float32, fillColor rl.Color, borderColor rl.Color) {
	// Draw filled inner rectangle
	rl.DrawRectangleRec(rect, fillColor)

	// Draw top border
	rl.DrawRectangle(int32(rect.X), int32(rect.Y), int32(rect.Width), int32(borderWidth), borderColor)
	// Draw bottom border
	rl.DrawRectangle(int32(rect.X), int32(rect.Y+rect.Height-borderWidth), int32(rect.Width), int32(borderWidth), borderColor)
	// Draw left border
	rl.DrawRectangle(int32(rect.X), int32(rect.Y), int32(borderWidth), int32(rect.Height), borderColor)
	// Draw right border
	rl.DrawRectangle(int32(rect.X+rect.Width-borderWidth), int32(rect.Y), int32(borderWidth), int32(rect.Height), borderColor)
}

func getPumpkin(pumpkinOrder int) rl.Texture2D {

	var pumpkin = rl.LoadTexture("assets/pumpkins/pumpkin_stage_" + strconv.Itoa(pumpkinOrder) + ".png")

	return pumpkin
}
