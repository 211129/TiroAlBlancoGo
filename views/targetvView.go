package views

import (
	"image/color"
	"TIROALBLANCO/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawTarget(target *models.Target, screen *ebiten.Image) {
	const size = 20
	targetImg := ebiten.NewImage(size, size)
	targetImg.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(target.X, target.Y)
	screen.DrawImage(targetImg, opts)
}
