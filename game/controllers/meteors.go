package controllers

import (
	"math/rand"
	"time"

	"github.com/PyMarcus/aerials/game/models"
	"github.com/PyMarcus/aerials/game/settings"
	"github.com/PyMarcus/aerials/game/utils"
)

func NewMeteor() *models.Meteor {
	sprite := utils.MeteorSprites[rand.Intn(len(utils.MeteorSprites))]

	return &models.Meteor{
		Position: models.Vector{},
		Sprite:   sprite,
	}
}

func CreateMovement(m *models.Meteor){
	rand.Seed(time.Now().UnixNano())

	velocityX := rand.Intn(401) - rand.Intn(401)
	velocityY := rand.Intn(401) - rand.Intn(401)

	m.Position.X += float64(velocityX) * settings.METEOR_SPEED
	m.Position.Y += float64(velocityY) * settings.METEOR_SPEED
}