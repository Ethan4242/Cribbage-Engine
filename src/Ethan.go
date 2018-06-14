package main

import (
	"fmt"
)

var x = 0

type Ethan struct {
	ID int
}

func (e Ethan) GetMove() {
	fmt.Println("Getting move from Ethan: ", e.ID)
}