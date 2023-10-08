package main

import (
	"log"
	"TIROALBLANCO/scenes"
	"TIROALBLANCO/models" 
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &scenes.TargetScene{
		Bow:    models.NewBow(100, 250), 
		Target: models.NewTarget(600, 250),
		Arrow:  models.NewArrow(110, 262),
		Score:  models.NewScore(),
	}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Tiro al Blanco")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}