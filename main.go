package main

import (
	"github.com/JGugino/survival-rpg/engine"
	"github.com/JGugino/survival-rpg/game"
)

var gameInstance engine.Game

func main() {
	gameInstance = game.StartGame()
}
