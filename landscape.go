package main

// Landscape represents a player's personal board
type Landscape struct {
	Tokens [][]Token
}

func NewLandscape() *Landscape {
	personalBoard := &Landscape{
		Tokens: make([][]Token, BoardSize),
	}
	for i := range personalBoard.Tokens {
		personalBoard.Tokens[i] = make([]Token, BoardSize)
		for j := range personalBoard.Tokens[i] {
			personalBoard.Tokens[i][j] = Token{Color: Empty, Height: NoHeight, Cube: false}
		}
	}
	return personalBoard
}
