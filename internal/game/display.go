package game

import (
	"fmt"
	"harmonies/internal/model"
	"harmonies/pkg"
)

// Display the game state
func (g *Game) Display() {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	// Display central board
	fmt.Println("=== CENTRAL BOARD ===")
	for i, space := range g.CentralBoard.Spaces {
		fmt.Printf("Space %d: ", i+1)
		for _, token := range space {
			fmt.Printf("%s %s \033[0m", model.ColorCode(token), model.ColorName(token))
		}
		fmt.Println()
	}
	fmt.Println()
	// Display personal board
	fmt.Println("=== YOUR PERSONAL BOARD ===")
	fmt.Println("    1   2   3   4   5   6   7  ")
	fmt.Println("  +---+---+---+---+---+---+---+")
	for i := 0; i < pkg.BoardSize; i++ {
		fmt.Printf("%d | ", i+1)
		for j := 0; j < pkg.BoardSize; j++ {
			token := g.Landscape.Tokens[i][j]
			if token.Color == model.Empty {
				fmt.Print("   | ")
			} else {
				symbol := "   "
				if token.Cube {
					symbol = " * "
				} else if token.Height > model.OneHigh {
					symbol = fmt.Sprintf(" %d ", int(token.Height))
				}
				fmt.Printf("%s%s\033[0m| ", model.ColorCode(token.Color), symbol)
			}
		}
		fmt.Println()
		fmt.Println("  +---+---+---+---+---+---+---+")
	}
	fmt.Println()

	// Display pouch information
	fmt.Printf("Tokens left in pouch: %d\n", len(g.Pouch.Tokens))
	fmt.Printf("Turn: %d\n", g.TurnCount)
	fmt.Printf("Current Score: %d\n", g.Score)
	fmt.Println("--- Count ---")
	fmt.Printf("Buildings: %d\n", g.CountBuildings())
	fmt.Printf("Trees: %d\n", g.CountTrees())
	fmt.Printf("Mountains: %d\n", g.CountMountains())
	fmt.Printf("Fields: %d\n", g.CountFields())
	fmt.Printf("Rivers: %d\n\n", g.CountRivers())

}

func (g *Game) DisplayFinalResults() {
	score := g.CalculateScore()
	g.Score = score

	fmt.Println("\n===== GAME OVER =====")
	fmt.Println("Final Score:", score)

	// Calculate suns (assuming side A and no spirit bonus)
	suns := g.CalculateSuns(score, false, false)

	fmt.Print("Suns earned: ")
	for i := 0; i < suns; i++ {
		fmt.Print("☀ ")
	}
	fmt.Println()

	// Display breakdown
	fmt.Println("Score Breakdown:")

	// Landscapes
	fmt.Println("=== Landscapes ===")
	buildings := g.CountBuildings()
	fmt.Printf("Buildings: %d × 5 = %d points\n", buildings, buildings*5)

	trees := g.CountTrees()
	fmt.Printf("Trees: %d points\n", trees)

	mountains := g.CountMountains()
	fmt.Printf("Mountains: %d points\n", mountains)

	fields := g.CountFields()
	fmt.Printf("Fields: %d × 5 = %d points\n", fields, fields*5)

	rivers := g.CountRivers()
	fmt.Printf("Rivers: %d points\n", rivers)

	landscapePoints := buildings*5 + trees + mountains + fields*5 + rivers
	fmt.Printf("Total Landscape Points: %d\n", landscapePoints)

	// Rating
	fmt.Println("\nPerformance Rating:")
	if suns <= 1 {
		fmt.Println("You're just getting started. Keep practicing!")
	} else if suns <= 3 {
		fmt.Println("Good job! You're developing your landscape skills.")
	} else if suns <= 5 {
		fmt.Println("Excellent! You've created a harmonious habitat.")
	} else {
		fmt.Println("Masterful! Your landscape is truly a work of art.")
	}
}
