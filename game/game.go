package game

import (
	"math"

	"github.com/PyMarcus/aerials/game/controllers"
	"github.com/PyMarcus/aerials/game/models"
	"github.com/PyMarcus/aerials/game/settings"
	"github.com/PyMarcus/aerials/game/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *models.Player
	Meteors          []*models.Meteor
	MeteorsSpawnTimer *models.Timer
}

func (g *Game) Update() error {
	g.keyBoardListenerController()
	g.Player.Update()
	g.MeteorsSpawnTimer.Update()

	if g.MeteorsSpawnTimer.Ready(){
		g.MeteorsSpawnTimer.Reset()
		meteor := controllers.NewMeteor()
		g.Meteors = append(g.Meteors, meteor)
	}

	for _, m := range g.Meteors {
		controllers.CreateMovement(m)
		m.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// background
	bg := &ebiten.DrawImageOptions{}
	X, Y := screen.Size()
	sw, sh := utils.Map.Size()

	bg.GeoM.Scale(
		float64(X)/float64(sw),
		float64(Y)/float64(sh),
	)
	screen.DrawImage(utils.Map, bg)

	// player
	g.Player.Draw(screen)
	for _, m := range g.Meteors {
		m.Draw(screen)
	}
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

