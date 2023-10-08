package views

import (
	"fmt"
	"TIROALBLANCO/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"golang.org/x/image/font/basicfont"
)

func DrawScore(score *models.Score, screen *ebiten.Image) {
	face := basicfont.Face7x13 
	text.Draw(screen, fmt.Sprintf("Puntuación: %d", score.Points), face, 10, 20, color.White)
}
