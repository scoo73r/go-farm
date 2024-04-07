package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth := float32(800)
	screenHeight := float32(450)
	
	rl.InitWindow(screenWidth, screenHeight, "go-farm")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
