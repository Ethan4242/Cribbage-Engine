package main

import (
	"fmt"
)

const NumPlayers = 4

type Game struct {
	players [NumPlayers]Player
}

// Sets the initial player array in the game
func (g Game) SetPlayers(index int, p Player) {
	g.players[index] = p
	p.GetMove()
}

func PointTotals() {
	fmt.Println("Getting point totals.")
}

// Returns True if winner
func HasWinner() {
	fmt.Println("Is there a winner?")
}
