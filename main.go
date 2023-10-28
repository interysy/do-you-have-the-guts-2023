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

	CENTRAL      = rl.GetScreenWidth() / 2
	SCREENHEIGHT = rl.GetScreenHeight()
	SCREENWIDTH  = rl.GetScreenWidth()
)

var state string = "login"
var password bool = false
var authenticated bool = false

func main() {

	rl.InitWindow(800, 450, "Gamer OS")
	rl.InitAudioDevice()
	fxCarve := rl.LoadSound("assets/audio/carve_pumpkin.wav")
	fxEmail := rl.LoadSound("assets/audio/email.wav")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	SCREENHEIGHT = rl.GetScreenHeight()
	SCREENWIDTH = rl.GetScreenWidth()

	var pumpkins = loadPumpkin()
	var secretFile rl.Texture2D = rl.LoadTexture("assets/lock.png")

	var chrome rl.Texture2D = rl.LoadTexture("assets/chrome.png")
	//var hambuga rl.Texture2D = rl.LoadTexture("assets/hambuga.png")
	var mail rl.Texture2D = rl.LoadTexture("assets/mail.png")
	var mail_notif rl.Texture2D = rl.LoadTexture("assets/mail_notif.png")
	var file_explorer rl.Texture2D = rl.LoadTexture("assets/file_explorer.png")
	var popout rl.Texture2D = rl.LoadTexture("assets/popout.png")

	var i = 1
	var particles []rl.Rectangle

	var windowHeight = rl.GetScreenHeight()
	var windowWidth = rl.GetScreenWidth()

	var email_popout = false
	var file_explorer_popout = false

	//object containing name of texture and texture location

	var textures = map[string]rl.Texture2D{
		"chrome":        chrome,
		"file_explorer": file_explorer,
		"email":         mail,
	}

	var textureOrder = []string{"email", "file_explorer", "chrome"}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if state == "login" {
			//loginScreen()
			getInput()
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
						rl.PlaySound(fxCarve)
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
			var desktopSingleMargin int32 = 25
			var desktopDoubleMargin int32 = 50

			rl.ClearBackground(PURPLE3)

			rl.DrawRectangle(desktopSingleMargin, desktopSingleMargin, int32(SCREENWIDTH)-desktopDoubleMargin, int32(SCREENHEIGHT)-desktopDoubleMargin, rl.Purple)
			rl.DrawRectangle(25, int32(windowHeight)-60, int32(windowWidth)-50, 40, rl.DarkPurple)
			//call draw function that passes in the map

			drawTaskbar(textures, textureOrder)

			var baseSecretFileSize = rl.NewRectangle(0, 0, float32(secretFile.Width), float32(secretFile.Height))
			var newWidth = float32(secretFile.Width) * 2
			var newHeight = float32(secretFile.Height) * 2
			var largeSecretFileSize = rl.NewRectangle(float32(desktopSingleMargin)+5, float32(desktopSingleMargin)+5, newWidth, newHeight)

			rl.DrawTexturePro(secretFile, baseSecretFileSize, largeSecretFileSize, rl.NewVector2(0, 0), 0, rl.White)

			var fileText = "click me"
			// collision check on secret file
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), largeSecretFileSize) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					rl.PlaySound(fxEmail)
					textures["email"] = mail_notif
					password = !password
				}
			}
			// var fileTextWidth = rl.MeasureText(fileText, 12)
			//var arr = rl.MeasureTextEx(rl.GetFontDefault(), fileText, 12, 1)
			//var fileTextWidth = arr.X
			//var fileTextHeight = arr.Y

			rl.DrawText(fileText, desktopSingleMargin+7, desktopSingleMargin+int32(newHeight)+5, 12, rl.Orange)

			if password == true {
				rectX := centraliseInX(300)
				rectY := centraliseInY(100)

				//TODO: use fergus' art to draw a rectangle with a border
				//TODO: Fix this so that the text is centralised properly
				rl.DrawRectangle(rectX, rectY, 300, 100, rl.Orange)

				rl.DrawText("Enter Password", centraliseInX(int(rl.MeasureText("Enter Password", 12)))-10, centraliseInY(100)+105, 16, rl.White)
				for i := 0; i < len(input); i++ {
					rl.DrawCircle(rectX+int32(i*(300/4))+25, rectY+50, 25, rl.White)
				}
				if getInput() {
					for i := 0; i < len(input); i++ {
						rl.DrawCircle(rectX+int32(i*(300/4))+25, rectY+50, 25, rl.Black)
					}
					rl.PlaySound(fxEmail)
					authenticated = true
				}
			}
			if email_popout == true {
				rl.DrawTexture(popout, 25, 25, rl.White)
				if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(365, 35), 10) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						email_popout = false
					}
				}
			}

			if file_explorer_popout == true {
				rl.DrawTexture(popout, 400, 25, rl.White)
				if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(745, 35), 10) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						file_explorer_popout = false
					}
				}
			}

			//Collision check on email icon
			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(45, float32(rl.GetScreenHeight()-40)), 18) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					email_popout = true
				}
			}

			// Collision on file explorer icon
			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(90, float32(rl.GetScreenHeight()-40)), 18) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					file_explorer_popout = true
				}
			}
		}
		rl.EndDrawing()
	}
	rl.UnloadSound(fxCarve)
	rl.UnloadSound(fxEmail)
	rl.CloseAudioDevice()
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

func drawTaskbar(textures map[string]rl.Texture2D, textureOrder []string) {
	//rl.DrawTexturePro(hambuga, rl.NewRectangle(0, 0, float32(hambuga.Width), float32(hambuga.Height)), rl.NewRectangle(25, float32(windowHeight)-60, float32(hambuga.Width)*1.5, float32(hambuga.Height)*1.5), rl.NewVector2(0, 0), 0, rl.White)
	//rl.DrawTexturePro(chrome, rl.NewRectangle(0, 0, float32(chrome.Width), float32(chrome.Height)), rl.NewRectangle(25+float32(chrome.Width)*1.5, float32(windowHeight)-60, float32(hambuga.Width)*1.5, float32(hambuga.Height)*1.5), rl.NewVector2(0, 0), 0, rl.White)
	// for each key in the map, draw the texture
	var windowHeight = rl.GetScreenHeight()
	//var windowWidth = rl.GetScreenWidth()

	var totalIconWidth = 0
	//sort.Strings(textures)
	for key := range textureOrder {
		var texture = textures[textureOrder[key]]
		rl.DrawTexturePro(texture, //texture
			rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)), //source
			rl.NewRectangle( //dest
				float32(totalIconWidth)+float32(texture.Width)*1.5, //x
				float32(windowHeight)-60,                           //y
				float32(texture.Width)*1.5,                         //width
				float32(texture.Height)*1.5),                       //height
			rl.NewVector2(0, 0),
			0,
			rl.White)
		totalIconWidth += int(texture.Width) + 10
	}
}
