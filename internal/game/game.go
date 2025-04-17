package game

import (
	"fmt"
	"harmonies/internal/model"
	"harmonies/pkg"
)

// Game represents the overall game state
type Game struct {
	Pouch        *model.Pouch
	CentralBoard *model.CentralBoard
	Landscape    *model.Landscape
	Score        int
	TurnCount    int
	GameOver     bool
}

func NewGame() *Game {
	game := &Game{
		Pouch:        model.NewPouch(),
		CentralBoard: model.NewCentralBoard(),
		Landscape:    model.NewLandscape(),
		Score:        0,
		TurnCount:    0,
		GameOver:     false,
	}
	game.FillCentralBoard()
	return game
}

func (g *Game) FillCentralBoard() {
	for i := range g.CentralBoard.Spaces {
		newTokens := g.Pouch.DrawTokens(3)
		g.CentralBoard.Spaces[i] = newTokens
	}
}

func (g *Game) PlayTurn() {
	g.TurnCount++
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
	g.CentralBoard.Spaces[spaceChoice] = []model.TokenColor{}

	// Place each token
	for i, token := range tokens {
		g.Display()
		fmt.Printf("Placing token %d/%d: %s\n", i+1, len(tokens), model.ColorName(token))

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
		g.CentralBoard.Spaces[i] = []model.TokenColor{}
	}

	g.FillCentralBoard()
	g.CheckGameOver()
}

func (g *Game) CheckGameOver() {
	// Count empty spaces
	emptySpaces := 0
	for i := 0; i < pkg.BoardSize; i++ {
		for j := 0; j < pkg.BoardSize; j++ {
			if g.Landscape.Tokens[i][j].Color == model.Empty {
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
