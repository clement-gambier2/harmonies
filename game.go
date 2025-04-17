package main

import (
	"fmt"
)

// Game represents the overall game state
type Game struct {
	Pouch        *Pouch
	CentralBoard *CentralBoard
	Landscape    *Landscape
	Score        int
	TurnCount    int
	GameOver     bool
}

func NewGame() *Game {
	game := &Game{
		Pouch:        NewPouch(),
		CentralBoard: NewCentralBoard(),
		Landscape:    NewLandscape(),
		Score:        0,
		TurnCount:    0,
		GameOver:     false,
	}

	// Fill the central board initially
	game.FillCentralBoard()

	return game
}

func (g *Game) FillCentralBoard() {
	for i := range g.CentralBoard.Spaces {
		newTokens := g.Pouch.DrawTokens(3)
		g.CentralBoard.Spaces[i] = newTokens
	}
}

// Display the game state
func (g *Game) Display() {
	// Clear screen
	fmt.Print("\033[H\033[2J")

	// Display central board
	fmt.Println("=== CENTRAL BOARD ===")
	for i, space := range g.CentralBoard.Spaces {
		fmt.Printf("Space %d: ", i+1)
		for _, token := range space {
			fmt.Printf("%s %s \033[0m", ColorCode(token), ColorName(token))
		}
		fmt.Println()
	}
	fmt.Println()
	// Display personal board
	fmt.Println("=== YOUR PERSONAL BOARD ===")
	fmt.Println("    1   2   3   4   5   6   7  ")
	fmt.Println("  +---+---+---+---+---+---+---+")
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%d | ", i+1)
		for j := 0; j < BoardSize; j++ {
			token := g.Landscape.Tokens[i][j]
			if token.Color == Empty {
				fmt.Print("   | ")
			} else {
				symbol := "   "
				if token.Cube {
					symbol = " * "
				} else if token.Height > OneHigh {
					symbol = fmt.Sprintf(" %d ", int(token.Height))
				}
				fmt.Printf("%s%s\033[0m| ", ColorCode(token.Color), symbol)
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

// Play a turn
func (g *Game) PlayTurn() {
	g.TurnCount++

	// Display the game state
	g.Display()

	// Ask the player which space to take tokens from
	fmt.Print("Choose a space to take tokens from (1-3): ")
	var spaceChoice int
	fmt.Scanln(&spaceChoice)
	spaceChoice--

	if spaceChoice < 0 || spaceChoice >= len(g.CentralBoard.Spaces) || len(g.CentralBoard.Spaces[spaceChoice]) == 0 {
		fmt.Println("Invalid choice. Try again.")
		g.TurnCount--
		return
	}

	// Take tokens from the chosen space
	tokens := g.CentralBoard.Spaces[spaceChoice]
	g.CentralBoard.Spaces[spaceChoice] = []TokenColor{}

	// Place each token
	for i, token := range tokens {
		g.Display()
		fmt.Printf("Placing token %d/%d: %s\n", i+1, len(tokens), ColorName(token))

		var row, col int
		for {
			fmt.Print("Enter position (row column): ")
			_, err := fmt.Scanf("%d %d", &row, &col)
			if err != nil {
				fmt.Println("Invalid input. Enter row and column as numbers separated by space.")
				continue
			}

			// Adjust for 1-indexed input
			row--
			col--

			if g.CanPlaceToken(row, col, token) {
				break
			}
			fmt.Println("Cannot place token there. Try again.")
		}

		g.PlaceToken(row, col, token)
		g.Display()
	}

	// End of turn: discard 6 tokens, refill the board
	for i := range g.CentralBoard.Spaces {
		g.CentralBoard.Spaces[i] = []TokenColor{}
	}

	g.FillCentralBoard()
	g.CheckGameOver()
}

func (g *Game) CheckGameOver() {
	// Count empty spaces
	emptySpaces := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == Empty {
				emptySpaces++
			}
		}
	}

	// Game ends if there are 2 or fewer empty spaces
	if emptySpaces <= 2 {
		g.GameOver = true
		return
	}

	// Game ends if the pouch is empty
	if len(g.Pouch.Tokens) == 0 {
		g.GameOver = true
		return
	}
}

func (g *Game) CanPlaceToken(x, y int, color TokenColor) bool {
	// Check if the position is within bounds
	if x < 0 || x >= BoardSize || y < 0 || y >= BoardSize {
		return false
	}

	// Check if there's an animal cube on the space
	if g.Landscape.Tokens[x][y].Cube {
		return false
	}

	// Tokens can always be placed on empty spaces
	if g.Landscape.Tokens[x][y].Color == Empty {
		return true
	}

	// Check stacking rules
	// Gray can be stacked on Gray to creates Mountains
	if color == Gray && g.Landscape.Tokens[x][y].Color == Gray && g.Landscape.Tokens[x][y].Height < ThreeHigh {
		return true
	}

	if color == Brown && g.Landscape.Tokens[x][y].Color == Brown && g.Landscape.Tokens[x][y].Height < TwoHigh {
		return true
	}

	// Green can be stacked on Brown to complete trees
	if color == Green && g.Landscape.Tokens[x][y].Color == Brown && g.Landscape.Tokens[x][y].Height < ThreeHigh {
		return true
	}

	// Red can be stacked on Red, Gray, or Brown to create buildings
	if color == Red && (g.Landscape.Tokens[x][y].Color == Red || g.Landscape.Tokens[x][y].Color == Gray || g.Landscape.Tokens[x][y].Color == Brown) && g.Landscape.Tokens[x][y].Height < TwoHigh {
		return true
	}

	return false
}

func (g *Game) PlaceToken(x, y int, color TokenColor) bool {
	if !g.CanPlaceToken(x, y, color) {
		return false
	}

	token := &g.Landscape.Tokens[x][y]
	if token.Color == Empty {
		token.Color = color
		token.Height = OneHigh
	} else {
		token.Height++
		//Buildings
		if color == Red {
			token.Color = Red
		}
		//Trees
		if color == Green && token.Color == Brown {
			token.Color = Green
		}
	}

	return true
}
