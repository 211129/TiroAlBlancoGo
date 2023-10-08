package models

type Score struct {
	Points int
}

func NewScore() *Score {
	return &Score{Points: 0}
}
