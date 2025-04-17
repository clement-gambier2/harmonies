package main

// Depth-first search to find contiguous groups
func dfs(g *Game, visited [][]bool, i, j int, color TokenColor) int {
	if i < 0 || i >= BoardSize || j < 0 || j >= BoardSize ||
		visited[i][j] || g.Landscape.Tokens[i][j].Color != color {
		return 0
	}

	visited[i][j] = true
	size := 1

	// Check all four directions
	for _, dir := range []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		size += dfs(g, visited, i+dir.dx, j+dir.dy, color)
	}

	return size
}
