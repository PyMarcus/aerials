package models

import "github.com/hajimehoshi/ebiten/v2"

type Meteor struct{
	Position Vector
	Sprite   *ebiten.Image
}

func (m *Meteor) Update(){

}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.Position.X, m.Position.Y)
	screen.DrawImage(m.Sprite, op)
}
