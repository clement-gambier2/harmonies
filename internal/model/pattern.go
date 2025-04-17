package model

type Orientation int

// Coordinate represents a position on the board
type Coordinate struct {
	X, Y int
}

// Pattern represents a required pattern for animal habitats
type Pattern []Coordinate
