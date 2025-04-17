package main

func (g *Game) CalculateScore() int {
	score := 0
	// Score for buildings (simplified)
	buildings := g.CountBuildings()
	score += buildings * 5

	// Score for trees (simplified)
	trees := g.CountTrees()
	score += trees * 3

	// Score for mountains (simplified)
	mountains := g.CountMountains()
	score += mountains * 4

	// Score for fields (simplified)
	fields := g.CountFields()
	score += fields * 5

	// Score for rivers (simplified)
	rivers := g.CountRivers()
	score += rivers * 3

	return score
}

func (g *Game) CountBuildings() int {
	count := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == Red && g.Landscape.Tokens[i][j].Height == TwoHigh {
				// Check if surrounded by at least 3 different colors
				colors := make(map[TokenColor]bool)

				// Check adjacent spaces
				for _, dir := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
					ni, nj := i+dir.dx, j+dir.dy
					if ni >= 0 && ni < BoardSize && nj >= 0 && nj < BoardSize && g.Landscape.Tokens[ni][nj].Color != Empty {

						colors[g.Landscape.Tokens[ni][nj].Color] = true
					}
				}

				if len(colors) >= 3 {
					count++
				}
			}
		}
	}
	return count
}

func (g *Game) CountTrees() int {
	points := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == Green {
				// Trees score based on height
				points += int(g.Landscape.Tokens[i][j].Height)
			}
		}
	}
	return points
}

func (g *Game) CountMountains() int {
	points := 0
	mountains := make(map[Coordinate]bool)

	// First pass: identify all mountains
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == Gray && g.Landscape.Tokens[i][j].Height > NoHeight {
				mountains[Coordinate{i, j}] = false // false means not yet counted
			}
		}
	}

	// Second pass: check which mountains are adjacent to other mountains
	for coord := range mountains {
		for _, dir := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			ni, nj := coord.X+dir.dx, coord.Y+dir.dy
			if _, exists := mountains[Coordinate{ni, nj}]; exists {
				mountains[coord] = true // This mountain is adjacent to another mountain
				break
			}
		}
	}

	// Calculate points
	for coord, isAdjacent := range mountains {
		if isAdjacent {
			// Mountains score based on height
			points += int(g.Landscape.Tokens[coord.X][coord.Y].Height)
		}
	}

	return points
}

func (g *Game) CountFields() int {
	count := 0
	visited := make([][]bool, BoardSize)
	for i := range visited {
		visited[i] = make([]bool, BoardSize)
	}

	// Find contiguous groups of yellow tokens
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == Yellow && !visited[i][j] {
				size := dfs(g, visited, i, j, Yellow)
				if size >= 2 {
					count++
				}
			}
		}
	}

	return count
}

func (g *Game) CountRivers() int {
	// Find the longest river (consecutive blue tokens)
	maxLength := 0
	visited := make([][]bool, BoardSize)
	for i := range visited {
		visited[i] = make([]bool, BoardSize)
	}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == Blue && !visited[i][j] {
				length := dfs(g, visited, i, j, Blue)
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}

	// Score based on river length
	points := 0
	if maxLength <= 6 {
		points = maxLength
	} else {
		points = 6 + (maxLength-6)*4
	}

	return points
}

func (g *Game) CalculateSuns(score int, sideB bool, spiritBonus bool) int {
	suns := 0

	// Score thresholds for suns
	if score >= 40 {
		suns = 1
	}
	if score >= 70 {
		suns = 2
	}
	if score >= 90 {
		suns = 3
	}
	if score >= 110 {
		suns = 4
	}
	if score >= 130 {
		suns = 5
	}
	if score >= 150 {
		suns = 6
	}
	if score >= 160 {
		suns = 7
	}

	// Bonus for side B
	if sideB {
		suns += 1
	}

	// Bonus for certain Nature's Spirit cards
	if spiritBonus {
		suns += 1
	}

	return suns
}
