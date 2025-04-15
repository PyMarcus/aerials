package models

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	Position Vector
	Sprite   *ebiten.Image
	Rotation float64
}


func (p *Player) Update(){

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	op.GeoM.Rotate(p.Rotation)
	screen.DrawImage(p.Sprite, op)
}
