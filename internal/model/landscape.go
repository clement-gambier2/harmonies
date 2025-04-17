package model

import (
	"harmonies/pkg"
)

// Landscape represents a player's personal board
type Landscape struct {
	Tokens [][]Token
}

func NewLandscape() *Landscape {
	personalBoard := &Landscape{
		Tokens: make([][]Token, pkg.BoardSize),
	}
	for i := range personalBoard.Tokens {
		personalBoard.Tokens[i] = make([]Token, pkg.BoardSize)
		for j := range personalBoard.Tokens[i] {
			personalBoard.Tokens[i][j] = Token{Color: Empty, Height: NoHeight, Cube: false}
		}
	}
	return personalBoard
}
