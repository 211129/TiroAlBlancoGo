package views

import (
	"image/color"
	"TIROALBLANCO/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawBow(bow *models.Bow, screen *ebiten.Image) {
	const width, height = 10, 40
	bowImg := ebiten.NewImage(width, height)
	bowImg.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(bow.X, bow.Y)
	screen.DrawImage(bowImg, opts)
}
