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
var lockFilePassword string

type File struct {
	texture rl.Texture2D
	open    bool
	file    rl.Texture2D
	name    string
	locked  bool
}

func main() {

	rl.InitWindow(800, 450, "Gamer OS")
	rl.InitAudioDevice()
	fxCarve := rl.LoadSound("assets/audio/carve_pumpkin.wav")
	fxEmail := rl.LoadSound("assets/audio/email.wav")
	fxRunning := rl.LoadMusicStream("assets/audio/running.ogg")
	fxStartup := rl.LoadSound("assets/audio/startup.ogg")
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
	var textFile = rl.LoadTexture("assets/txt_file.png")
	var imageFile = rl.LoadTexture("assets/image.png")
	var cult = rl.LoadTexture("assets/cult.png")
	var goatHead = rl.LoadTexture("assets/goathead.png")
	var calendar = rl.LoadTexture("assets/calendar.png")
	var desktop_frame = rl.LoadTexture("assets/desktop_frame.png")

	var i = 1
	var particles []rl.Rectangle

	var windowHeight = rl.GetScreenHeight()
	var windowWidth = rl.GetScreenWidth()

	var email_popout = false
	var real_email_popout1 = false
	var real_email_popout2 = false
	var real_email_popout3 = false
	var real_email_popout4 = false

	var file_explorer_popout = false

	//object containing name of texture and texture location

	var textures = map[string]rl.Texture2D{
		"chrome":        chrome,
		"file_explorer": file_explorer,
		"email":         mail,
	}

	var textureOrder = []string{"email", "file_explorer", "chrome"}
	rl.PlayMusicStream(fxRunning)

	var fileExplorerTextures = map[string]File{
		"textFile1":  File{texture: textFile, open: false, file: popout, name: "text"},
		"textFile2":  File{texture: textFile, open: false, file: popout, name: "text"},
		"textFile3":  File{texture: textFile, open: false, file: popout, name: "text"},
		"imageFile1": File{texture: imageFile, open: false, file: cult, name: "cult"},
		"imageFile2": File{texture: imageFile, open: false, file: goatHead, name: "goat", locked: true},
		"imageFile3": File{texture: imageFile, open: false, file: calendar, name: "diary"},
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if state == "login" {
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
						rl.PlaySound(fxStartup)
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
			rl.DrawTexture(desktop_frame, 0, 0, rl.DarkPurple)
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

			rl.DrawText(fileText, desktopSingleMargin+7, desktopSingleMargin+int32(newHeight)+5, 12, rl.Orange)

			if password == true {
				rectX := centraliseInX(300)
				rectY := centraliseInY(100)

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
				rl.DrawTexture(popout, 125, 25, rl.White)
				if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(465, 35), 10) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						email_popout = false
					}
				}

				//add 4 boxed emails here
				if textures["email"] == mail_notif {
					//add 4 boxes as wide as the popout with "email 1" text inside the box
					rl.DrawRectangle(150, 50, 300, 48, rl.DarkPurple)
					rl.DrawRectangle(150, 100, 300, 48, rl.DarkPurple)
					rl.DrawRectangle(150, 150, 300, 48, rl.DarkPurple)
					rl.DrawRectangle(150, 200, 300, 48, rl.DarkPurple)

					//titles
					rl.DrawText("Pete (BLOCKED)", 155, 52, 16, rl.White)
					rl.DrawText("The Lottery Co.", 155, 102, 16, rl.White)
					rl.DrawText("Lukas", 155, 152, 16, rl.White)
					rl.DrawText("Email 4", 155, 202, 16, rl.White)

					//subtitles
					rl.DrawText("Subject: Please listen to me...", 155, 65, 12, rl.White)
					rl.DrawText("Subject: You have won a free car", 155, 115, 12, rl.White)
					rl.DrawText("Subject: Soon...", 155, 165, 12, rl.White)
					rl.DrawText("Subject: You have won a free car", 155, 215, 12, rl.White)

					if real_email_popout1 == true {
						rl.DrawTexture(popout, 300, 25, rl.White)
						//add text inside drawing

						rl.DrawText("From: Pete (BLOCKED)\tDate: 15/08\nDude,\nI really don't think you should be getting involved \nwith these guys, I had a look at the stuff you sent me \nearlier today and it sounds really... weird.\nPlease respond, we gotta talk.", 315, 35, 12, rl.White)

						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(640, 35), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout1 = false
							}
						}
					}
					if real_email_popout2 == true {
						rl.DrawTexture(popout, 310, 35, rl.White)
						rl.DrawText("From: The Lottery Co.\t16/08\nYOU WON A FREE CAR!!!!\nCALL US NOW ON +44 (0)7314982430 TO\nCLAIM YOUR FREE 2001 HONDA CIVIC!!!!", 325, 45, 12, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(650, 45), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout2 = false
							}
						}
					}
					if real_email_popout3 == true {
						rl.DrawTexture(popout, 320, 45, rl.White)
						rl.DrawText("From: Lukas (SB)\t30/10\nGood evening,\nfirst:\nthe Festival will be fulfilled In the Very next Evening.\nEach Individual shall Get all He can Think of.\nthe brotherhood looks forward to seeing you,\nLukas", 335, 55, 12, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(660, 65), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout3 = false
							}
						}
					}
					if real_email_popout4 == true {
						rl.DrawTexture(popout, 330, 55, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(670, 75), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout4 = false
							}
						}
					}

					//collision check on email 1,2,3,4
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 50, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout1 = true
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 100, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout2 = true
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 150, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout3 = true
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 200, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout4 = true
						}
					}

				} else {
					rl.DrawText("You have no new emails", 150, 50, 12, rl.White)
				}

			}

			if file_explorer_popout == true {
				// rl.DrawTexture(popout, 400, 25, rl.White)
				rl.DrawTexturePro(popout, rl.NewRectangle(0, 0, float32(popout.Width), float32(popout.Height)), rl.NewRectangle(380, 25, float32(popout.Width)*1.1, float32(popout.Height)), rl.NewVector2(0, 0), 0, rl.White)
				populateFileExplorer(fileExplorerTextures, popout, fxEmail)
				rl.DrawRectangle(int32(float32(popout.Width))+395, 35, 10, 10, rl.White)
				if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(395+float32(popout.Width), 35), 10) {
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
	rl.DrawRectangleRec(rect, fillColor)

	rl.DrawRectangle(int32(rect.X), int32(rect.Y), int32(rect.Width), int32(borderWidth), borderColor)
	rl.DrawRectangle(int32(rect.X), int32(rect.Y+rect.Height-borderWidth), int32(rect.Width), int32(borderWidth), borderColor)
	rl.DrawRectangle(int32(rect.X), int32(rect.Y), int32(borderWidth), int32(rect.Height), borderColor)
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

func populateFileExplorer(fileExplorerTextures map[string]File, popout rl.Texture2D, unlockSound rl.Sound) {

	var order = []string{"textFile1", "textFile2", "textFile3", "imageFile1", "imageFile2", "imageFile3"}
	var i float32 = 0
	// var g bool = false
	var nextLineY = 0
	for key := range order {
		//fmt.Println(string(key) + "\n")
		var texture = fileExplorerTextures[order[key]].texture
		// if g {
		// 	nextLineY = int(texture.Height*2 + 10)
		// } else {
		// 	nextLineY = 0
		// }

		var x = float32(400) + float32(texture.Width) + ((float32(texture.Width)*2 + 10) * i)
		var y = float32(25) + float32(texture.Height) + float32(nextLineY)

		rl.DrawTexturePro(texture, //texture
			rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
			rl.NewRectangle(
				x,                          //x
				y,                          //y
				float32(texture.Width)*2,   //width
				float32(texture.Height)*2), //height
			rl.NewVector2(0, 0),
			0,
			rl.White)
		var fileText = fileExplorerTextures[order[key]].name
		var textSizing = rl.MeasureTextEx(rl.GetFontDefault(), fileText, 12, 1)

		rl.DrawText(fileText, int32(x)+texture.Width/2, int32(y)+int32(textSizing.Y)*5, 12, rl.Orange)

		if fileExplorerTextures[order[key]].open == true && fileExplorerTextures[order[key]].locked == false {
			fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: openPopUpFileExpolorer(popout, fileExplorerTextures[order[key]].file, int(centraliseInX(int(fileExplorerTextures[order[key]].file.Width))), int(centraliseInY(int(fileExplorerTextures[order[key]].file.Height))), fileExplorerTextures, order[key]), file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: fileExplorerTextures[order[key]].locked}
		} else if fileExplorerTextures[order[key]].open == true && fileExplorerTextures[order[key]].locked == true {
			fmt.Print("WTH")
			fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: fileExplorerTextures[order[key]].open, file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: unlockFile(unlockSound)}
		}

		// rl.DrawCircle(int32(x), int32(y), float32(texture.Width), rl.White)
		// rl.DrawRectangle(int32(x), int32(y), int32(texture.Width)*2, int32(texture.Height)*2, rl.White)
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{x, y, float32(texture.Width * 2), float32(texture.Height * 2)}) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: openPopUpFileExpolorer(popout, fileExplorerTextures[order[key]].file, int(centraliseInX(int(fileExplorerTextures[order[key]].file.Width))), int(centraliseInY(int(fileExplorerTextures[order[key]].file.Height))), fileExplorerTextures, order[key]), file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: fileExplorerTextures[order[key]].locked}
				// fileExplorerTextures[order[key]] = File{texture: texture, open: true}
			}
		}
		i++
		// if g {
		// 	i++
		// }
		// g = !g
	}

}

func openPopUpFileExpolorer(popout rl.Texture2D, image rl.Texture2D, x int, y int, textures map[string]File, key string) bool {
	var locked = textures[key].locked
	if locked == true {
		return true
	}
	rl.DrawTexture(image, int32(x)+10, int32(y)+10, rl.White)
	if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(float32(x)+float32(image.Width)-5, float32(y)+10), 20) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			return false
		}
	}
	return true
}

func unlockFile(sound rl.Sound) bool {
	var rectX int32 = 400
	var rectY int32 = 200

	rl.DrawRectangle(centraliseInX(int(300)), centraliseInY(int(100)), 300, 100, rl.Orange)
	rl.DrawText("Enter Password", centraliseInX(int(rl.MeasureText("Enter Password", 12)))-10, centraliseInY(100)+105, 16, rl.White)
	for i := 0; i < len(fileInput); i++ {
		rl.DrawCircle(rectX+int32(i*(300/4))+25, rectY+50, 25, rl.White)
	}
	if fileGetInput() {
		for i := 0; i < len(fileInput); i++ {
			rl.DrawCircle(rectX+int32(i*(300/4))+25, rectY+50, 25, rl.Black)
		}
		rl.PlaySound(sound)
		authenticated = true
		return false
	}
	return true
}
