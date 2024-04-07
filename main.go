package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 1000
)

var (
	running      = true
	bkgColor     = rl.NewColor(147, 211, 196, 255)
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerMoving                                  bool
	playerDir                                     int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame                                   int

	frameCount int

	playerSpeed float32 = 3
	musicPaused bool
	music       rl.Music
	cam         rl.Camera2D
)

func drawscene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyE) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerDir = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDir = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyF) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir = 3
		playerRight = true
	}
	if rl.IsKeyPressed(rl.KeyP) {
		musicPaused = !musicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()

	playerSrc.X = 0

	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if frameCount%8 == 1 {
			playerFrame++
		}

		playerSrc.X = playerSrc.Width * float32(playerFrame)
	}
	rl.UpdateMusicStream(music)
	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}

	playerSrc.Y = playerSrc.Height * float32(playerDir)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height)))
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
}
func render() {
	rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height)))
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)
	drawscene()
	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "go-farm")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("Sprites/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("Sprites/Characters/BasicCharacterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("audio/tap-room-rag-small.mp3")
	musicPaused = false
	//	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(
			float32(playerDest.X-(playerDest.Width/2)),
			float32(playerDest.Y-(playerDest.Height))),
		0.0,
		2.0)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}
	rl.CloseWindow()
}
