package models

import (
	"math/rand"
	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct{
	Position Vector
	Sprite   *ebiten.Image
}

func (m *Meteor) Update(){

}

func (m *Meteor) Draw(screen *ebiten.Image) {
	rotationSpeed := -0.02 + rand.Float64()*0.04
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.Position.X, m.Position.Y)
	op.GeoM.Rotate(rotationSpeed)
	screen.DrawImage(m.Sprite, op)
}
