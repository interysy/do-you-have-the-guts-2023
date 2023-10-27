package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func load_texture(path string) {
  image := rl.LoadImage(path)  // Loaded in CPU memory (RAM)
  texture := rl.LoadTextureFromImage(image) // Image converted to texture, GPU memory (VRAM)

  rl.UnloadImage(image) // Once image has been converted to texture and uploaded to VRAM, it can be unloaded from RAM
}

