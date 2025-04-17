package main

import (
	"fmt"
	"harmonies/internal/game"
	"math/rand"
	"time"
)

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new gameInstance
	gameInstance := game.NewGame()

	fmt.Println("==============================")
	fmt.Println("     HARMONIES - SOLO MODE    ")
	fmt.Println("==============================")
	fmt.Println("Welcome to the solo mode of Harmonies!")
	fmt.Println("Build your landscape wisely to earn as many points as possible.")
	fmt.Println("The gameInstance ends when you have 2 or fewer empty spaces left")
	fmt.Println("or when the pouch is empty.")
	fmt.Println("\nPress Enter to start the gameInstance...")
	fmt.Scanln()

	// Game loop
	for !gameInstance.GameOver {
		gameInstance.PlayTurn()
	}

	// Calculate and display final results
	gameInstance.DisplayFinalResults()

	fmt.Println("\nThank you for playing Harmonies Solo Mode!")
}
