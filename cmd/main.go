package main

import (
	"github.com/PyMarcus/aerials/game"
	"github.com/PyMarcus/aerials/game/controllers"
	"github.com/hajimehoshi/ebiten/v2"
)

// IMPORTANT: sudo apt install -y xorg-dev libgl1-mesa-dev libglu1-mesa-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev build-essential
func main() {
	g := &game.Game{
		Player: controllers.NewPlayer(),
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
