package game

import (
	"harmonies/internal/model"
	"harmonies/pkg"
)

func (g *Game) CanPlaceToken(x, y int, color model.TokenColor) bool {
	// Check if the position is within bounds
	if x < 0 || x >= pkg.BoardSize || y < 0 || y >= pkg.BoardSize {
		return false
	}

	// Check if there's an animal cube on the space
	if g.Landscape.Tokens[x][y].Cube {
		return false
	}

	// Tokens can always be placed on empty spaces
	if g.Landscape.Tokens[x][y].Color == model.Empty {
		return true
	}

	// Check stacking rules
	// Gray can be stacked on Gray to creates Mountains
	if color == model.Gray && g.Landscape.Tokens[x][y].Color == model.Gray && g.Landscape.Tokens[x][y].Height < model.ThreeHigh {
		return true
	}

	if color == model.Brown && g.Landscape.Tokens[x][y].Color == model.Brown && g.Landscape.Tokens[x][y].Height < model.TwoHigh {
		return true
	}

	// Green can be stacked on Brown to complete trees
	if color == model.Green && g.Landscape.Tokens[x][y].Color == model.Brown && g.Landscape.Tokens[x][y].Height < model.ThreeHigh {
		return true
	}

	// Red can be stacked on Red, Gray, or Brown to create buildings
	if color == model.Red && (g.Landscape.Tokens[x][y].Color == model.Red || g.Landscape.Tokens[x][y].Color == model.Gray || g.Landscape.Tokens[x][y].Color == model.Brown) && g.Landscape.Tokens[x][y].Height < model.TwoHigh {
		return true
	}

	return false
}

func (g *Game) PlaceToken(x, y int, color model.TokenColor) bool {
	if !g.CanPlaceToken(x, y, color) {
		return false
	}

	token := &g.Landscape.Tokens[x][y]
	if token.Color == model.Empty {
		token.Color = color
		token.Height = model.OneHigh
	} else {
		token.Height++
		//Buildings
		if color == model.Red {
			token.Color = model.Red
		}
		//Trees
		if color == model.Green && token.Color == model.Brown {
			token.Color = model.Green
		}
	}

	return true
}
