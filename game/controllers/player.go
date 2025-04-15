package controllers

import (
	"github.com/PyMarcus/aerials/game/models"
	"github.com/PyMarcus/aerials/game/settings"
	"github.com/PyMarcus/aerials/game/utils"
)

func NewPlayer() *models.Player{
	return &models.Player{
		Position: models.Vector{X: settings.INITIAL_X, Y: settings.INITIAL_Y},
		Sprite: utils.Player,
	}
}


