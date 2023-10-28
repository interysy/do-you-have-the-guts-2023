package main

import rl "github.com/gen2brain/raylib-go/raylib"

func loginScreen() {
	var particles []rl.Rectangle
	var i = 1
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var pumpkin rl.Texture2D = rl.LoadTexture("assets/pumpkins/pumpkin_stage_1.png")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

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
				if i < 7 {
					i++
					particles = generateParticles(10, centraliseInX(25), 100)
					pumpkin = getPumpkin(i)
				}

			}
		}
		updateParticles(particles, int32(CENTRAL))
		renderParticles(particles)

		rl.EndDrawing()
	}
}
