package utils

import (
	"testing"

	"github.com/PyMarcus/aerials/game/settings"
)

func TestMustLoadImage(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Panic unexpected! %v", r)
		}
	}()

	mustLoadImage(settings.PLAYER_IMG)
}
