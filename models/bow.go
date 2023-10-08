package models

type Bow struct {
	X, Y float64
}

func NewBow(x, y float64) *Bow {
	return &Bow{X: x, Y: y}
}
