package main

// Animal habitat pattern orientation
type Orientation int

const (
	North Orientation = iota
	East
	South
	West
)

// Coordinate represents a position on the board
type Coordinate struct {
	X, Y int
}

// Pattern represents a required pattern for animal habitats
type Pattern []Coordinate

// Generate all possible orientations of a pattern
func GenerateOrientations(pattern Pattern) []Pattern {
	orientations := make([]Pattern, 4)

	// North orientation (original)
	orientations[North] = make(Pattern, len(pattern))
	copy(orientations[North], pattern)

	// East orientation (rotate 90 degrees clockwise)
	orientations[East] = make(Pattern, len(pattern))
	for i, coord := range pattern {
		orientations[East][i] = Coordinate{coord.Y, -coord.X}
	}

	// South orientation (rotate 180 degrees)
	orientations[South] = make(Pattern, len(pattern))
	for i, coord := range pattern {
		orientations[South][i] = Coordinate{-coord.X, -coord.Y}
	}

	// West orientation (rotate 270 degrees clockwise)
	orientations[West] = make(Pattern, len(pattern))
	for i, coord := range pattern {
		orientations[West][i] = Coordinate{-coord.Y, coord.X}
	}

	// Normalize each orientation (shift so the top-left coordinate is at 0,0)
	for i := range orientations {
		normalizePattern(orientations[i])
	}

	return orientations
}

// Normalize a pattern so that the minimum x and y values are 0
func normalizePattern(pattern Pattern) {
	// Find the minimum x and y values
	minX, minY := pattern[0].X, pattern[0].Y
	for _, coord := range pattern {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
	}

	// Shift all coordinates
	for i := range pattern {
		pattern[i].X -= minX
		pattern[i].Y -= minY
	}
}
