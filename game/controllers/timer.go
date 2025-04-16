package controllers

import (
	"time"

	"github.com/PyMarcus/aerials/game/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewTimer(duration time.Duration) *models.Timer{
	return &models.Timer{
		CurrentTicks: 0,
		TargetTicks: int(duration.Milliseconds()*int64(ebiten.TPS())/1000),
	}
}
