package game

import (
	"github.com/PyMarcus/aerials/game/models"
	screenitems "github.com/PyMarcus/aerials/game/screen_items"
	"github.com/PyMarcus/aerials/game/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	PlayerPosition models.Vector
}

func (g *Game) Update() error {
	g.keyBoardListenerController()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(g.PlayerPosition.X, g.PlayerPosition.Y)
	screenitems.DrawItem(screen, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.WIDTH, settings.HEIGHT
}

func (g *Game) keyBoardListenerController() {

	speed := float64(settings.SPEED / float64(ebiten.TPS()))

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.PlayerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.PlayerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.PlayerPosition.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.PlayerPosition.X += speed
	}
}

