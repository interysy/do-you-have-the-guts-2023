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

	CENTRAL      = rl.GetScreenWidth() / 2
	SCREENHEIGHT = rl.GetScreenHeight()
	SCREENWIDTH  = rl.GetScreenWidth()
)

var state string = "startup"
var password bool = false
var authenticated bool = false
var lockFilePassword string

type File struct {
	texture rl.Texture2D
	open    bool
	file    rl.Texture2D
	name    string
	locked  bool
	text    string
}

func main() {

	rl.InitWindow(800, 450, "Gamer OS")
	rl.InitAudioDevice()
	fxCarve := rl.LoadSound("assets/audio/carve_pumpkin.wav")
	fxEmail := rl.LoadSound("assets/audio/email.wav")
	fxRunning := rl.LoadSound("assets/audio/running.ogg")
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
	var goatHead = rl.LoadTexture("assets/goathead_embed.png")
	var calendar = rl.LoadTexture("assets/calendar.png")
	var desktop_frame = rl.LoadTexture("assets/desktop_frame.png")
	var missing_poster = rl.LoadTexture("assets/missing_poster_v2.png")
	var cipher = rl.LoadTexture("assets/cipher.png")
	var bin = rl.LoadTexture("assets/recycle.png")

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
		"bin":           bin,
	}

	var fadeAlpha float32 = 1
	var cntr = 0
	var textureOrder = []string{"email", "file_explorer", "chrome", "bin"}

	var fileExplorerTextures = map[string]File{
		"textFile1":  File{texture: textFile, open: false, file: popout, name: "text", text: "We solemnly affirm our devotion, on the eve \nof our year's twilight. Samhain, the \ncommunion of the worshippers on the \nhallowed eve, where the veil between the \nliving and the damned is at its thinnest. \nIn the sanctuary of the moon’s wick, \nthe inky midnights’ calling brings us to unveil. \n"},
		"textFile2":  File{texture: textFile, open: false, file: popout, name: "text", text: "If you find this, stop looking around my stuff,\nbut listen.They're right, and you \nare all in denial, Samhain will change\n everything. Soon you'll see, I hope you can\njoin me in this journey.\nDo not worry for me, I am happier with them.\nWith The Brotherhood."},
		"imageFile1": File{texture: imageFile, open: false, file: cult, name: "cult"},
		"imageFile2": File{texture: imageFile, open: false, file: goatHead, name: "goat", locked: true},
		"imageFile3": File{texture: imageFile, open: false, file: calendar, name: "diary"},
		"imageFile4": File{texture: imageFile, open: false, file: cipher, name: "cipher"},
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if !rl.IsSoundPlaying(fxRunning) {
			rl.PlaySound(fxRunning)
		}
		if state == "startup" {
			rl.ClearBackground(PURPLE1)

			rl.DrawRectangle(0, 0, int32(SCREENWIDTH), int32(SCREENHEIGHT), rl.Fade(rl.Black, fadeAlpha))

			rl.DrawTexturePro(missing_poster, //texture
				rl.NewRectangle(0, 0, float32(missing_poster.Width), float32(missing_poster.Height)), //source
				rl.NewRectangle( //dest
					float32(centraliseInX(int(missing_poster.Width*2))),  //x
					float32(centraliseInY(int(missing_poster.Height*2))), //y
					float32(missing_poster.Width)*2,                      //width
					float32(missing_poster.Height)*2),                    //height
				rl.NewVector2(0, 0),
				0,
				rl.Fade(rl.White, fadeAlpha))

			if fadeAlpha < 0 {
				state = "login"
			} else if cntr < 200 {
				cntr++
			} else {
				fadeAlpha -= 0.02
			}
		}
		if state == "login" {
			var pumpkin rl.Texture2D = pumpkins[i-1]

			rl.ClearBackground(PURPLE1)

			xCentralRectangleCoordinate := float32(centraliseInX(90))
			rl.DrawCircle(int32(xCentralRectangleCoordinate+45), int32(rl.GetScreenHeight()/4+50), 50, rl.DarkPurple)

			rl.DrawTexture(pumpkin, centraliseInX(int(pumpkin.Width)), int32(rl.GetScreenHeight()/4)+5, rl.White)

			loginUserName := "Jonathan"
			rl.DrawText(loginUserName, centraliseInX(len(loginUserName)*7), 250, 15, rl.Orange)
			rl.DrawRectangleRounded(rl.NewRectangle(float32(centraliseInX(200)), 275, 200, 20), 0.1, 0, rl.Orange)
			rl.DrawText("password", 300, 275, 16, rl.White)

			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(xCentralRectangleCoordinate+45, float32(rl.GetScreenHeight()/4+50)), 50) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					i++
					if i < 10 {
						rl.PlaySound(fxCarve)
						particles = generateParticles(25, centraliseInX(25), 100)
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

			drawTaskbar(textures, textureOrder)

			var baseSecretFileSize = rl.NewRectangle(0, 0, float32(secretFile.Width), float32(secretFile.Height))
			var newWidth = float32(secretFile.Width) * 2
			var newHeight = float32(secretFile.Height) * 2
			var largeSecretFileSize = rl.NewRectangle(float32(desktopSingleMargin)+5, float32(desktopSingleMargin)+5, newWidth, newHeight)

			rl.DrawTexturePro(secretFile, baseSecretFileSize, largeSecretFileSize, rl.NewVector2(0, 0), 0, rl.White)

			var fileText = "click me"
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
					authenticated = !authenticated
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

					rl.DrawText("Pete (BLOCKED)", 155, 52, 16, rl.White)
					rl.DrawText("The Lottery Co.", 155, 102, 16, rl.White)
					rl.DrawText("Lukas", 155, 152, 16, rl.White)
					rl.DrawText("(Sender Unknown)", 155, 202, 16, rl.White)

					//subtitles
					rl.DrawText("Subject: Please listen to me...", 155, 65, 14, rl.White)
					rl.DrawText("Subject: You have won a free car", 155, 115, 14, rl.White)
					rl.DrawText("Subject: Soon...", 155, 165, 14, rl.White)
					rl.DrawText("Subject: You have won a free car", 155, 215, 14, rl.White)

					if real_email_popout1 == true {
						rl.DrawTexture(popout, 300, 25, rl.White)

						rl.DrawText("From: Pete (BLOCKED)\tDate: 15/08\nDude,\nI really don't think you should be getting involved \nwith these guys, I had a look at the stuff you \nsent me earlier today and it sounds really... weird.\n\nPlease respond, we gotta talk.", 315, 35, 14, rl.White)

						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(640, 35), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout1 = false
							}
						}
					}
					if real_email_popout2 == true {
						rl.DrawTexture(popout, 310, 35, rl.White)
						rl.DrawText("From: The Lottery Co.\t16/08\nYOU WON A FREE CAR!!!!\nCALL US NOW ON +44 (0)7314982430 TO\nCLAIM YOUR FREE 2001 HONDA CIVIC!!!!", 325, 45, 14, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(650, 45), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout2 = false
							}
						}
					}
					if real_email_popout3 == true {
						rl.DrawTexture(popout, 320, 45, rl.White)
						rl.DrawText("From: Lukas (SB)\t30/10\nGood evening,\nfirst:\nFestival will be fulfilled In the Very next Evening.\nEach Individual shall Get all He can Think of.\nthe brotherhood looks forward to seeing you,\n\nLukas", 335, 55, 14, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(660, 65), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout3 = false
							}
						}
					}

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 50, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout1 = !real_email_popout1
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 100, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout2 = !real_email_popout2
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 150, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout3 = !real_email_popout3
						}
					}

					//titles

					if real_email_popout4 == true {
						rl.DrawTexture(popout, 330, 55, rl.White)
						rl.DrawText("From: (Sender Unknown)\t07/11\nI see that you are finally looking in the right \nplaces.\nIf you ever want to see him again, \nmake sure to look closely.\nRemember when it all began.\nI have already given you the first.\nsecond:\nThe Rest, Each and Everything, is in your hands.", 345, 65, 14, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(670, 75), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout4 = false
							}
						}
					}

					//collision check on email 1,2,3,4
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 200, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout4 = !real_email_popout4
						}
					}

				} else {
					rl.DrawRectangle(150, 50, 300, 48, rl.DarkPurple)
					rl.DrawRectangle(150, 100, 300, 48, rl.DarkPurple)
					rl.DrawRectangle(150, 150, 300, 48, rl.DarkPurple)

					rl.DrawText("Pete (BLOCKED)", 155, 52, 16, rl.White)
					rl.DrawText("The Lottery Co.", 155, 102, 16, rl.White)
					rl.DrawText("Lukas", 155, 152, 16, rl.White)

					//subtitles
					rl.DrawText("Subject: Please listen to me...", 155, 65, 14, rl.White)
					rl.DrawText("Subject: You have won a free car", 155, 115, 14, rl.White)
					rl.DrawText("Subject: Soon...", 155, 165, 14, rl.White)

					if real_email_popout1 == true {
						rl.DrawTexture(popout, 300, 25, rl.White)
						//add text inside drawing

						rl.DrawText("From: Pete (BLOCKED)\tDate: 15/08\nDude,\nI really don't think you should be getting involved \nwith these guys, I had a look at the stuff you \nsent me earlier today and it sounds really... weird.\n\nPlease respond, we gotta talk.", 315, 35, 14, rl.White)

						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(640, 35), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout1 = false
							}
						}
					}
					if real_email_popout2 == true {
						rl.DrawTexture(popout, 310, 35, rl.White)
						rl.DrawText("From: The Lottery Co.\t16/08\nYOU WON A FREE CAR!!!!\nCALL US NOW ON +44 (0)7314982430 TO\nCLAIM YOUR FREE 2001 HONDA CIVIC!!!!", 325, 45, 14, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(650, 45), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout2 = false
							}
						}
					}
					if real_email_popout3 == true {
						rl.DrawTexture(popout, 320, 45, rl.White)
						rl.DrawText("From: Lukas (SB)\t30/10\nGood evening,\nfirst:\nFestival will be fulfilled In the Very next Evening.\nEach Individual shall Get all He can Think of.\nthe brotherhood looks forward to seeing you,\n\nLukas", 335, 55, 14, rl.White)
						if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(660, 65), 10) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								real_email_popout3 = false
							}
						}
					}

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 50, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout1 = !real_email_popout1
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 100, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout2 = !real_email_popout2
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 150, 300, 48)) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							real_email_popout3 = !real_email_popout3
						}
					}
				}

			}

			if file_explorer_popout == true {
				rl.DrawTexturePro(popout, rl.NewRectangle(0, 0, float32(popout.Width), float32(popout.Height)), rl.NewRectangle(380, 25, float32(popout.Width)*1.1, float32(popout.Height)), rl.NewVector2(0, 0), 0, rl.White)

				populateFileExplorer(fileExplorerTextures, popout, fxEmail, textFile)
				if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(395+float32(popout.Width), 35), 10) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						file_explorer_popout = !file_explorer_popout
					}
				}
			}

			//Collision check on email icon
			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(45, float32(rl.GetScreenHeight()-40)), 18) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					email_popout = !email_popout
				}
			}

			// Collision on file explorer icon
			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(90, float32(rl.GetScreenHeight()-40)), 18) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					file_explorer_popout = !file_explorer_popout
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
	var windowHeight = rl.GetScreenHeight()
	var totalIconWidth = 0

	for key := range textureOrder {
		var texture = textures[textureOrder[key]]
		rl.DrawTexturePro(texture,
			rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
			rl.NewRectangle( //dest
				float32(totalIconWidth)+float32(texture.Width)*1.5,
				float32(windowHeight)-60,
				float32(texture.Width)*1.5,
				float32(texture.Height)*1.5),
			rl.NewVector2(0, 0),
			0,
			rl.White)
		totalIconWidth += int(texture.Width) + 10
	}
}

func populateFileExplorer(fileExplorerTextures map[string]File, popout rl.Texture2D, unlockSound rl.Sound, textFile rl.Texture2D) {
	var order = []string{"textFile1", "textFile2", "imageFile4", "imageFile1", "imageFile2", "imageFile3"}
	var i float32 = 0
	var nextLineY = 0
	for key := range order {
		var texture = fileExplorerTextures[order[key]].texture

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
			fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: openPopUpFileExpolorer(popout, fileExplorerTextures[order[key]].file, int(centraliseInX(int(fileExplorerTextures[order[key]].file.Width))), int(centraliseInY(int(fileExplorerTextures[order[key]].file.Height))), fileExplorerTextures, order[key], textFile), file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: fileExplorerTextures[order[key]].locked, text: fileExplorerTextures[order[key]].text}

		} else if fileExplorerTextures[order[key]].open == true && fileExplorerTextures[order[key]].locked == true {
			fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: fileExplorerTextures[order[key]].open, file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: unlockFile(unlockSound)}
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{x, y, float32(texture.Width * 2), float32(texture.Height * 2)}) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if fileExplorerTextures[order[key]].open == true {
					fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: false, file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: fileExplorerTextures[order[key]].locked, text: fileExplorerTextures[order[key]].text}
				} else {
					fileExplorerTextures[order[key]] = File{texture: fileExplorerTextures[order[key]].texture, open: openPopUpFileExpolorer(popout, fileExplorerTextures[order[key]].file, int(centraliseInX(int(fileExplorerTextures[order[key]].file.Width))), int(centraliseInY(int(fileExplorerTextures[order[key]].file.Height))), fileExplorerTextures, order[key], textFile), file: fileExplorerTextures[order[key]].file, name: fileExplorerTextures[order[key]].name, locked: fileExplorerTextures[order[key]].locked, text: fileExplorerTextures[order[key]].text}
				}
			}
		}
		i++
	}

}

func openPopUpFileExpolorer(popout rl.Texture2D, image rl.Texture2D, x int, y int, textures map[string]File, key string, textFile rl.Texture2D) bool {
	var locked = textures[key].locked
	if locked == true {
		return true
	}
	rl.DrawTexture(image, int32(x)+10, int32(y)+10, rl.White)
	if textures[key].texture == textFile {
		rl.DrawText(textures[key].text, int32(x)+20, int32(y)+40, 16, rl.White)
	}

	if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.NewVector2(float32(x)+float32(image.Width)-5, float32(y)+10), 20) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			return false
		}
	}
	return true
}

func unlockFile(sound rl.Sound) bool {
	var rectX int32 = 300
	var rectY int32 = 200

	rl.DrawRectangle(centraliseInX(int(300)), centraliseInY(int(100)), 300, 100, rl.Orange)
	rl.DrawText("Enter Password", centraliseInX(int(rl.MeasureText("Enter Password", 12)))-10, centraliseInY(100)+105, 16, rl.White)
	for i := 0; i < len(fileInput); i++ {
		rl.DrawCircle(rectX+int32(i*(300/4))-10, rectY+25, 25, rl.White)
	}
	if fileGetInput() {
		for i := 0; i < len(fileInput); i++ {
			rl.DrawCircle(rectX+int32(i*(300/4))-10, rectY+25, 25, rl.Black)
		}
		rl.PlaySound(sound)
		authenticated = true
		return false
	}
	return true
}
