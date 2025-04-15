package models

import "github.com/hajimehoshi/ebiten/v2"

type Meteor struct{
	Position Vector
	Sprite   *ebiten.Image
}