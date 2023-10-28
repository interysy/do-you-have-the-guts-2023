package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var passwordString = "1234"
var input = ""
var prev = rl.KeyNull

func getInput() bool {
	key := rl.GetKeyPressed()
	if prev == int(key) {
		return false
	}

	if len(input) == 4 && key != rl.KeyBackspace || len(input) == 0 && key == rl.KeyBackspace {
		return false
	}
	switch key {
	case rl.KeyZero:
		input += "0"
		break
	case rl.KeyOne:
		input += "1"
		break
	case rl.KeyTwo:
		input += "2"
		break
	case rl.KeyThree:
		input += "3"
		break
	case rl.KeyFour:
		input += "4"
		break
	case rl.KeyFive:
		input += "5"
		break
	case rl.KeySix:
		input += "6"
		break
	case rl.KeySeven:
		input += "7"
		break
	case rl.KeyEight:
		input += "8"
		break
	case rl.KeyNine:
		input += "9"
		break
	case rl.KeyBackspace:
		input = input[:len(input)-1]
		break
	}

	prev = int(key)
	if len(input) >= 4 {
		if passwordString == input {
			fxScream := rl.LoadSound("assets/audio/scream.ogg")
			rl.PlaySound(fxScream)
			return true
		}
		return false
	}
	return false
}

var filePasswordString = "1234"
var fileInput = ""
var filePrev = rl.KeyNull

func fileGetInput() bool {
	key := rl.GetKeyPressed()
	if prev == int(key) {
		return false
	}

	if len(input) == 4 && key != rl.KeyBackspace || len(input) == 0 && key == rl.KeyBackspace {
		return false
	}
	switch key {
	case rl.KeyZero:
		fileInput += "0"
		break
	case rl.KeyOne:
		fileInput += "1"
		break
	case rl.KeyTwo:
		fileInput += "2"
		break
	case rl.KeyThree:
		fileInput += "3"
		break
	case rl.KeyFour:
		fileInput += "4"
		break
	case rl.KeyFive:
		fileInput += "5"
		break
	case rl.KeySix:
		fileInput += "6"
		break
	case rl.KeySeven:
		fileInput += "7"
		break
	case rl.KeyEight:
		fileInput += "8"
		break
	case rl.KeyNine:
		fileInput += "9"
		break
	case rl.KeyBackspace:
		fileInput = fileInput[:len(fileInput)-1]
		break
	}

	filePrev = int(key)
	if len(fileInput) >= 4 {
		if filePasswordString == fileInput {
			return true
		}
		return false
	}
	return false
}
