package main

import (
	"fmt"
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

var state string = "login"

func main() {

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	var pumpkins = loadPumpkin()
	var secretFile rl.Texture2D = rl.LoadTexture("assets/lock.png")

	var i = 1
	var particles []rl.Rectangle

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if state == "login" {
			//loginScreen()
			var pumpkin rl.Texture2D = pumpkins[i-1] // = rl.LoadTexture("assets/pumpkins/pumpkin_stage_1.png")

			rl.ClearBackground(PURPLE1)

			xCentralRectangleCoordinate := float32(centraliseInX(90))
			//var iconRectangle = rl.NewRectangle(xCentralRectangleCoordinate, float32(rl.GetScreenHeight())/4+7, 90, 90)

			// DrawBorderedRectangle(iconRectangle, 5, PURPLE1, rl.White)
			rl.DrawCircle(int32(xCentralRectangleCoordinate+45), int32(rl.GetScreenHeight()/4+50), 50, rl.DarkPurple)

			//rl.DrawTexturePro(texture, rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)), rl.NewRectangle(posX, posY, width*scaleX, height*scaleY), rl.NewVector2(0, 0), 0, tint)
			rl.DrawTexture(pumpkin, centraliseInX(int(pumpkin.Width)), int32(rl.GetScreenHeight()/4)+5, rl.White)

			loginUserName := "common jp morgan enjoyer"
			rl.DrawText(loginUserName, centraliseInX(len(loginUserName)*7), 250, 15, rl.Orange)
			rl.DrawRectangleRounded(rl.NewRectangle(float32(centraliseInX(200)), 275, 200, 20), 0.1, 0, rl.Orange)
			rl.DrawText("password", 300, 275, 16, rl.White)

			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(xCentralRectangleCoordinate+45, float32(rl.GetScreenHeight()/4+50)), 50) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					i++
					if i < 10 {
						fmt.Println(i)
						particles = generateParticles(10, centraliseInX(25), 100)
					}
					if i == 9 {
						state = "desktop"
					}

				}
				updateParticles(particles, int32(CENTRAL))
				renderParticles(particles)
			}
		}
		if state == "desktop" {

			rl.ClearBackground(PURPLE3)

			rl.DrawRectangle(25, 25, int32(rl.GetScreenWidth())-50, int32(rl.GetScreenHeight())-50, rl.Purple)
			drawTaskbar()

			rl.DrawTexturePro(secretFile, rl.NewRectangle(0, 0, float32(secretFile.Width), float32(secretFile.Height)), rl.NewRectangle(float32(centraliseInX(int(secretFile.Width*3))+25), float32(centraliseInY(int(secretFile.Height*3))), float32(secretFile.Width)*3, float32(secretFile.Height)*3), rl.NewVector2(0, 0), 0, rl.White)
			rl.DrawText("click me", centraliseInX(len("click me")*3), centraliseInY(1)+50, 12, rl.Orange)

		}
		rl.EndDrawing()
	}
}

func centraliseInX(size int) int32 {
	var centralXCoordinate = rl.GetScreenWidth()/2 - size/2
	return int32(centralXCoordinate)
}

func centraliseInY(size int) int32 {
	var centralYCoordinate = rl.GetScreenHeight()/2 - size/2
	return int32(centralYCoordinate)
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

func loadPumpkin() []rl.Texture2D {
	var pumpkin []rl.Texture2D
	for i := 1; i < 10; i++ {
		pumpkin = append(pumpkin, rl.LoadTexture("assets/pumpkins/pumpkin_stage_"+strconv.Itoa(i)+".png"))
	}
	return pumpkin
}

func drawTaskbar() {

	var chrome rl.Texture2D = rl.LoadTexture("assets/chrome.png")
	var hambuga rl.Texture2D = rl.LoadTexture("assets/hambuga.png")

	var windowHeight = rl.GetScreenHeight()
	var windowWidth = rl.GetScreenWidth()

	rl.DrawRectangle(25, int32(windowHeight)-60, int32(windowWidth)-50, 40, rl.DarkPurple)
	rl.DrawTexturePro(hambuga, rl.NewRectangle(0, 0, float32(hambuga.Width), float32(hambuga.Height)), rl.NewRectangle(25, float32(windowHeight)-60, float32(hambuga.Width)*1.5, float32(hambuga.Height)*1.5), rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawTexturePro(chrome, rl.NewRectangle(0, 0, float32(chrome.Width), float32(chrome.Height)), rl.NewRectangle(25+float32(chrome.Width)*1.5, float32(windowHeight)-60, float32(hambuga.Width)*1.5, float32(hambuga.Height)*1.5), rl.NewVector2(0, 0), 0, rl.White)

}
