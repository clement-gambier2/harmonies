package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new game
	game := NewGame()

	fmt.Println("==============================")
	fmt.Println("     HARMONIES - SOLO MODE    ")
	fmt.Println("==============================")
	fmt.Println("Welcome to the solo mode of Harmonies!")
	fmt.Println("Build your landscape wisely to earn as many points as possible.")
	fmt.Println("The game ends when you have 2 or fewer empty spaces left")
	fmt.Println("or when the pouch is empty.")
	fmt.Println("\nPress Enter to start the game...")
	fmt.Scanln()

	// Game loop
	for !game.GameOver {
		game.PlayTurn()
	}

	// Calculate and display final results
	game.DisplayFinalResults()

	fmt.Println("\nThank you for playing Harmonies Solo Mode!")
}
