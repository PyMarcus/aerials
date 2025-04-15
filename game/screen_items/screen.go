package screenitems

import (
	"github.com/PyMarcus/aerials/game/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawItem(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	screen.DrawImage(utils.Player, options)
}