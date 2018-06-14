package main

import (
	"fmt"
)

func main() {
	fmt.Println("Beginning Cribbage-Engine..")
	startGame()
}

// Initializes all of the relevant game values and runs main loop
func startGame() {
	fmt.Println("Beginning Cribbage Game..")

	g := Game{} // Create game instance

	// Do initializations
	initializePlayers(g)
	
}

func initializePlayers(g Game) {
	fmt.Println("Initializing players.. ")

	// Set up player instances in game
	for i := 0; i < 4; i++ {
		g.SetPlayers(i, Ethan{i})
	}

	// TODO: Set up the teams
	// TODO: Set up game board
	

}