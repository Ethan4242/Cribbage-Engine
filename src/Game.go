package main

import (
	"fmt"
)

type Game struct {
	NumPlayers int
}


func (g Game) setPlayers() int {
	g.NumPlayers = 4
	return g.NumPlayers
}

func PointTotals() {

}

// Returns True if winner
func HasWinner() {
	fmt.Println("Testing worked!")
}
