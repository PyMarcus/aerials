package utils

import (
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"runtime"

	"github.com/PyMarcus/aerials/game/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

var Player = mustLoadImage(settings.PLAYER_IMG)

func mustLoadImage(name string) *ebiten.Image {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(currentFile), "..", "..", settings.ASSETS)
	fullPath := filepath.Join(basePath, name)
	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
