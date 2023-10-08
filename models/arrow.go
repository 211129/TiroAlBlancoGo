package models

type Arrow struct {
	X, Y     float64
	Velocity float64
	Fired    bool
}

func NewArrow(x, y float64) *Arrow {
	return &Arrow{X: x, Y: y, Velocity: 10, Fired: false}
}
