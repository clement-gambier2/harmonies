package main

// Animal represents an animal card
type Animal struct {
	Name           string
	Points         []int
	HabitatPattern Pattern
	Color          TokenColor // Color on which the animal cube must be placed
	CubePosition   Coordinate // Position within the pattern where the cube goes
	CubesLeft      int        // Number of cubes left to place
	Completed      bool       // Whether all cubes have been placed
	Card           [][]rune   // Visual representation of the card
	Orientations   []Pattern  // Possible orientations of the pattern
}
