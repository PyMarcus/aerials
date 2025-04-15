package controllers

import (
	"math/rand"

	"github.com/PyMarcus/aerials/game/models"
	"github.com/PyMarcus/aerials/game/utils"
)

func NewMeteor() *models.Meteor {
	sprite := utils.MeteorSprites[rand.Intn(len(utils.MeteorSprites))]

	return &models.Meteor{
		Position: models.Vector{},
		Sprite:   sprite,
	}
}