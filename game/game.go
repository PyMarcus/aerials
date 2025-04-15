package game

import (
	"math"

	"github.com/PyMarcus/aerials/game/models"
	"github.com/PyMarcus/aerials/game/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *models.Player
	Meteors          []*models.Meteor
}

func (g *Game) Update() error {
	g.keyBoardListenerController()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.WIDTH, settings.HEIGHT
}

func (g *Game) keyBoardListenerController() {

	speed := float64(settings.SPEED / float64(ebiten.TPS()))
	rotation := (math.Pi / float64(ebiten.TPS()) ) / 2

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Player.Position.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Player.Position.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Player.Rotation -= rotation
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Player.Rotation += rotation
	}
}

