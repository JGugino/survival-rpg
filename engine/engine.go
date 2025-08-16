package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	entityHandler
}

type Window struct {
	WindowWidth  int32
	WindowHeight int32
	WindowTitle  string
	TargetFPS    int32
}

func (game *Game) OpenGameWindow(window Window) {
	PrintInfo(fmt.Sprintf("Initializing Game Window - %s", window.WindowTitle))
	rl.InitWindow(window.WindowWidth, window.WindowHeight, window.WindowTitle)
	defer rl.CloseWindow()

	PrintInfo(fmt.Sprintf("Setting Target FPS - %d", window.TargetFPS))
	rl.SetTargetFPS(window.TargetFPS)

	PrintInfo("Starting Game Loop")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.EndDrawing()
	}
}
