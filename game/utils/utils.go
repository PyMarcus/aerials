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
var MeteorSprites = mustLoadImages("meteors/*.png")

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

func mustLoadImages(pattern string) []*ebiten.Image {
	fullPattern := filepath.Join(baseAssetsPath(), pattern)
	files, err := filepath.Glob(fullPattern)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		panic("no files matched pattern: " + pattern)
	}

	var images []*ebiten.Image
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			panic(err)
		}

		images = append(images, ebiten.NewImageFromImage(img))
	}

	return images
}


func baseAssetsPath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(currentFile), "..", "..", settings.ASSETS)
}
