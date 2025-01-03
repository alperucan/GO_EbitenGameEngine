package assets

import (
	"embed"
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("player.png")
var MeteorSprites = mustLoadImages("meteors/*.png")
var LaserSprite = mustLoadImage("laser.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}
