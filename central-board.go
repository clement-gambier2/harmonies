package main

// CentralBoard represents the central board with token spaces
type CentralBoard struct {
	Spaces [][]TokenColor // 3 spaces with 3 tokens each
}

// Create a new central board
func NewCentralBoard() *CentralBoard {
	centralBoard := &CentralBoard{
		Spaces: make([][]TokenColor, 3),
	}
	for i := range centralBoard.Spaces {
		centralBoard.Spaces[i] = make([]TokenColor, 3)
	}
	return centralBoard
}
