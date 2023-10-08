package scenes

import (
	"image"
	_ "image/png"
	"os"
	"log"
	"TIROALBLANCO/models"
	"TIROALBLANCO/views"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var backgroundImage *ebiten.Image

func init() {
	file, err := os.Open("assets/paris.png")
	if err != nil {
		log.Fatalf("Error al abrir la imagen: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Error al decodificar la imagen: %v", err)
	}

	backgroundImage = ebiten.NewImageFromImage(img)
}

type TargetScene struct {
	Target *models.Target
	Bow    *models.Bow
	Arrow  *models.Arrow
	Score  *models.Score
}

func (g *TargetScene) Update() error {
	if g.Arrow.Fired {
		g.Arrow.X += g.Arrow.Velocity
		if g.Arrow.X > 800 {
			g.Arrow.Fired = false
			g.Arrow.X = g.Bow.X
		}
	}
	
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !g.Arrow.Fired {
		g.Arrow.Fired = true
	}

	if g.Arrow.Fired && g.Arrow.X+40 >= g.Target.X && g.Arrow.X <= g.Target.X+20 && 
		g.Arrow.Y+5 >= g.Target.Y && g.Arrow.Y <= g.Target.Y+20 {
		g.Score.Points++
		g.Arrow.Fired = false
		g.Arrow.X = g.Bow.X
	}

	return nil
}

func (g *TargetScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(backgroundImage, nil) // Dibuja el fondo de paris.png
	views.DrawBow(g.Bow, screen)
	views.DrawTarget(g.Target, screen)
	if g.Arrow.Fired {
		views.DrawArrow(g.Arrow, screen)
	}
	views.DrawScore(g.Score, screen)
}

func (g *TargetScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
