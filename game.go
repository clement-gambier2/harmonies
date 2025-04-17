package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Game represents the overall game state
type Game struct {
	Pouch          *Pouch
	CentralBoard   *CentralBoard
	Landscape      *Landscape
	AvailableCards []*Animal
	PlayerCards    []*Animal
	Score          int
	TurnCount      int
	GameOver       bool
}

// Initialize a new game
func NewGame() *Game {
	game := &Game{
		Pouch:          NewPouch(),
		CentralBoard:   NewCentralBoard(),
		Landscape:      NewLandscape(),
		AvailableCards: make([]*Animal, 0),
		PlayerCards:    make([]*Animal, 0),
		Score:          0,
		TurnCount:      0,
		GameOver:       false,
	}

	// Fill the central board initially
	game.FillCentralBoard()

	// Create initial animal cards
	game.CreateInitialAnimalCards()

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

	// Display available animal cards
	fmt.Println("=== AVAILABLE ANIMAL CARDS ===")
	for i, animal := range g.AvailableCards {
		fmt.Printf("%d. %s (Points: %v, Cubes left: %d)\n", i+1, animal.Name, animal.Points, animal.CubesLeft)
		for _, line := range animal.Card {
			fmt.Println(string(line))
		}
		fmt.Println()
	}

	// Display player cards
	fmt.Println("=== YOUR ANIMAL CARDS ===")
	for i, animal := range g.PlayerCards {
		status := "In Progress"
		if animal.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s (Points: %v, Status: %s)\n", i+1, animal.Name, animal.Points, status)
		for _, line := range animal.Card {
			fmt.Println(string(line))
		}
		fmt.Println()
	}

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
	fmt.Printf("Current Score: %d\n\n", g.Score)
}

// Create initial animal cards
func (g *Game) CreateInitialAnimalCards() {
	// Initialize with a few predefined animals
	animals := []*Animal{
		{
			Name:         "Bear",
			Points:       []int{0, 4, 8, 12},
			Color:        Brown,
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {0, 1}, {1, 0},
			},
			Card: [][]rune{
				[]rune("Bear (Brown)"),
				[]rune("Points: 0,4,8,12"),
				[]rune("Pattern:"),
				[]rune("B B"),
				[]rune("B  "),
			},
		},
		{
			Name:         "Fox",
			Points:       []int{0, 3, 7, 10},
			Color:        Red,
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {1, 0}, {1, 1},
			},
			Card: [][]rune{
				[]rune("Fox (Red)"),
				[]rune("Points: 0,3,7,10"),
				[]rune("Pattern:"),
				[]rune("R  "),
				[]rune("R R"),
			},
		},
		{
			Name:         "Eagle",
			Points:       []int{0, 5, 9, 14},
			Color:        Gray,
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {0, 1}, {0, 2},
			},
			Card: [][]rune{
				[]rune("Eagle (Gray)"),
				[]rune("Points: 0,5,9,14"),
				[]rune("Pattern:"),
				[]rune("G G G"),
			},
		},
		{
			Name:         "Deer",
			Points:       []int{0, 4, 8, 12},
			Color:        Green,
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {1, 0}, {2, 0},
			},
			Card: [][]rune{
				[]rune("Deer (Green)"),
				[]rune("Points: 0,4,8,12"),
				[]rune("Pattern:"),
				[]rune("G"),
				[]rune("G"),
				[]rune("G"),
			},
		},
		{
			Name:         "Fish",
			Points:       []int{0, 3, 6, 10},
			Color:        Blue,
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {0, 1}, {1, 1},
			},
			Card: [][]rune{
				[]rune("Fish (Blue)"),
				[]rune("Points: 0,3,6,10"),
				[]rune("Pattern:"),
				[]rune("B B"),
				[]rune("  B"),
			},
		},
		{
			Name:         "Rabbit",
			Points:       []int{0, 2, 5, 8},
			Color:        Yellow,
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {1, 1}, {2, 0},
			},
			Card: [][]rune{
				[]rune("Rabbit (Yellow)"),
				[]rune("Points: 0,2,5,8"),
				[]rune("Pattern:"),
				[]rune("Y  "),
				[]rune("  Y"),
				[]rune("Y  "),
			},
		},
	}

	// Generate all orientations for each animal
	for i := range animals {
		animals[i].Orientations = GenerateOrientations(animals[i].HabitatPattern)
	}

	// Shuffle the animals
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	// Take the first 3 for the initial available cards
	g.AvailableCards = animals[:3]
}

// Display the final score and results
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

	// Animal cards
	fmt.Println("=== Animals ===")
	animalPoints := 0
	for _, animal := range g.PlayerCards {
		if animal.CubesLeft < len(animal.Points) {
			points := animal.Points[len(animal.Points)-animal.CubesLeft-1]
			animalPoints += points
			fmt.Printf("%s: %d points\n", animal.Name, points)
		}
	}
	fmt.Printf("Total Animal Points: %d\n\n", animalPoints)

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

	fmt.Printf("\nTOTAL SCORE: %d\n", animalPoints+landscapePoints)

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

		// After placing a token, offer to place an animal cube
		g.Display()
		fmt.Println(g.Landscape)
		//fmt.Print("Do you want to place an animal cube? (y/n): ")
		//var wantCube string
		//fmt.Scanln(&wantCube)

		//if strings.ToLower(wantCube) == "y" {
		//	for {
		//		if len(g.PlayerCards) == 0 {
		//			fmt.Println("No animal cards available.")
		//			break
		//		}
		//
		//		fmt.Print("Choose an animal card (1-", len(g.PlayerCards), ") or 0 to skip: ")
		//		var cardChoice int
		//		fmt.Scanln(&cardChoice)
		//
		//		if cardChoice == 0 {
		//			break
		//		}
		//
		//		cardChoice--
		//		if cardChoice < 0 || cardChoice >= len(g.PlayerCards) {
		//			fmt.Println("Invalid choice. Try again.")
		//			continue
		//		}
		//
		//		if g.PlaceAnimalCube(cardChoice) {
		//			fmt.Println("Animal cube placed successfully!")
		//			break
		//		} else {
		//			fmt.Println("Cannot place cube for this animal. Try another card.")
		//		}
		//	}
		//}
	}

	// Option to take an animal card
	g.Display()
	fmt.Print("Do you want to take an animal card? (y/n): ")
	var wantCard string
	fmt.Scanln(&wantCard)

	if strings.ToLower(wantCard) == "y" {
		if len(g.PlayerCards) >= MaxCardsAtOnce {
			fmt.Println("You already have the maximum number of cards.")
		} else if len(g.AvailableCards) > 0 {
			fmt.Print("Choose an animal card (1-", len(g.AvailableCards), "): ")
			var cardChoice int
			fmt.Scanln(&cardChoice)
			cardChoice--

			if cardChoice >= 0 && cardChoice < len(g.AvailableCards) {
				g.PlayerCards = append(g.PlayerCards, g.AvailableCards[cardChoice])
				g.AvailableCards = append(g.AvailableCards[:cardChoice], g.AvailableCards[cardChoice+1:]...)
			} else {
				fmt.Println("Invalid choice.")
			}
		} else {
			fmt.Println("No animal cards available.")
		}
	} else {
		// If the player didn't take a card, offer to discard one
		fmt.Print("Do you want to discard an available animal card? (y/n): ")
		var wantDiscard string
		fmt.Scanln(&wantDiscard)

		if strings.ToLower(wantDiscard) == "y" && len(g.AvailableCards) > 0 {
			fmt.Print("Choose a card to discard (1-", len(g.AvailableCards), "): ")
			var cardChoice int
			fmt.Scanln(&cardChoice)
			cardChoice--

			if cardChoice >= 0 && cardChoice < len(g.AvailableCards) {
				g.AvailableCards = append(g.AvailableCards[:cardChoice], g.AvailableCards[cardChoice+1:]...)
			} else {
				fmt.Println("Invalid choice.")
			}
		}
	}

	// End of turn: discard 6 tokens, refill the board
	// In solo mode, discard remaining tokens
	for i := range g.CentralBoard.Spaces {
		g.CentralBoard.Spaces[i] = []TokenColor{}
	}

	// Refill the central board
	g.FillCentralBoard()

	// Refill animal cards
	for len(g.AvailableCards) < 3 && len(g.Pouch.Tokens) > 0 {
		// Create a new animal card (simplified)
		animal := &Animal{
			Name:         fmt.Sprintf("Animal %d", rand.Intn(100)),
			Points:       []int{0, rand.Intn(4) + 2, rand.Intn(4) + 6, rand.Intn(4) + 10},
			Color:        TokenColor(rand.Intn(6) + 1),
			CubePosition: Coordinate{0, 0},
			CubesLeft:    3,
			HabitatPattern: Pattern{
				{0, 0}, {rand.Intn(2), rand.Intn(2)}, {rand.Intn(2), rand.Intn(2)},
			},
		}

		// Generate card display
		animal.Card = [][]rune{
			[]rune(fmt.Sprintf("%s (%s)", animal.Name, ColorName(animal.Color))),
			[]rune(fmt.Sprintf("Points: %v", animal.Points)),
			[]rune("Pattern:"),
		}

		// Generate orientations
		animal.Orientations = GenerateOrientations(animal.HabitatPattern)

		g.AvailableCards = append(g.AvailableCards, animal)
	}

	// Check if the game is over
	g.CheckGameOver()
}

// Check if the game is over
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

// Check if a token can be placed at a given position
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

	// Check stacking rules:
	// 1. Gray can be stacked on Gray
	if color == Gray && g.Landscape.Tokens[x][y].Color == Gray && g.Landscape.Tokens[x][y].Height < ThreeHigh {
		return true
	}

	// 2. Green can be stacked on Brown
	if color == Green && g.Landscape.Tokens[x][y].Color == Brown && g.Landscape.Tokens[x][y].Height < ThreeHigh {
		return true
	}

	// 3. Red can be stacked on Red, Gray, or Brown
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
		// Handle stacking
		token.Height++
		// If stacking red on something, the color changes to red
		if color == Red {
			token.Color = Red
		}
		// If stacking green on brown, the color changes to green
		if color == Green && token.Color == Brown {
			token.Color = Green
		}
	}

	return true
}

// Check if a habitat pattern matches at a given position and orientation
func (g *Game) CheckHabitatMatch(animal *Animal, startX, startY int, orientation int) bool {
	pattern := animal.Orientations[orientation]

	// Check each position in the pattern
	for _, coord := range pattern {
		x, y := startX+coord.X, startY+coord.Y

		// Check if position is valid
		if x < 0 || x >= BoardSize || y < 0 || y >= BoardSize {
			return false
		}

		// For the position where the cube goes, check if it's empty of cubes
		if coord.X == animal.CubePosition.X && coord.Y == animal.CubePosition.Y {
			if g.Landscape.Tokens[x][y].Cube {
				return false
			}

			// Also check if the color matches
			if g.Landscape.Tokens[x][y].Color != animal.Color {
				return false
			}
		}

		// For mountain and tree heights, check if they match
		// This is simplified - you would need more detailed logic for exact matching
		if g.Landscape.Tokens[x][y].Color == Empty {
			return false
		}
	}

	return true
}

// Try to place an animal cube on the board
func (g *Game) PlaceAnimalCube(animalIndex int) bool {
	if animalIndex < 0 || animalIndex >= len(g.PlayerCards) {
		return false
	}

	animal := g.PlayerCards[animalIndex]
	if animal.CubesLeft <= 0 || animal.Completed {
		return false
	}

	// Try all possible positions and orientations
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			for orientation := 0; orientation < 4; orientation++ {
				if g.CheckHabitatMatch(animal, x, y, orientation) {
					// Found a match, place the cube
					pattern := animal.Orientations[orientation]
					for _, coord := range pattern {
						if coord.X == animal.CubePosition.X && coord.Y == animal.CubePosition.Y {
							g.Landscape.Tokens[x+coord.X][y+coord.Y].Cube = true
							animal.CubesLeft--
							if animal.CubesLeft == 0 {
								animal.Completed = true
							}
							return true
						}
					}
				}
			}
		}
	}

	return false
}
