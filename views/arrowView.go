package views

import (
	"image/color"
	"TIROALBLANCO/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawArrow(arrow *models.Arrow, screen *ebiten.Image) {
	const width, height = 40, 5
	arrowImg := ebiten.NewImage(width, height)
	arrowImg.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(arrow.X, arrow.Y)
	screen.DrawImage(arrowImg, opts)
}
