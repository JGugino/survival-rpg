package game

import "github.com/JGugino/survival-rpg/engine"

func StartGame() engine.Game {
	println("Starting game...")

	game := engine.Game{}

	game.OpenGameWindow(engine.Window{
		WindowWidth:  1200,
		WindowHeight: 800,
		WindowTitle:  "Survival RPG - Gugino",
		TargetFPS:    60,
	})

	return game
}
