
package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var password = "1234"
var input = ""

func getInput () bool {
	switch (key) {
		    case rl.KeyZero:
			    input += "0"
		    case rl.KeyOne:
			    input += "1"
		    case rl.KeyTwo:
			    input += "2"
		    case rl.KeyThree:
			    input += "3"
		    case rl.KeyFour:
			    input += "4"
		    case rl.KeyFive:
			    input += "5"
		    case rl.KeySix:
			    input += "6"
		    case rl.KeySeven:
			    input += "7"
		    case rl.KeyEight:
			    input += "8"
		    case rl.KeyNine:
			    input += "9"
	}
	if len(input) == 4:
		if password == input:
			return true
		else
			input = ""
	return false
}
