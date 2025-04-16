package main

// PersonalBoard represents a player's personal board
type PersonalBoard struct {
	Tokens [][]Token
}

// Create a new personal board
func NewPersonalBoard() *PersonalBoard {
	personalBoard := &PersonalBoard{
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
