package main

import (
	"github.com/PyMarcus/aerials/game"
	"github.com/PyMarcus/aerials/game/models"
	"github.com/PyMarcus/aerials/game/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

// IMPORTANT: sudo apt install -y xorg-dev libgl1-mesa-dev libglu1-mesa-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev build-essential
func main() {
	g := &game.Game{
		PlayerPosition: models.Vector{X: settings.INITIAL_X, Y: settings.INITIAL_Y},
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
