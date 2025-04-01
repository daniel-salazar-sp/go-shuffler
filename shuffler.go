package main

import (
	"fmt"
	"github.com/daniel-salazar-sp/go-shuffler/deck"
)

func main() {
	fmt.Println("Shuffler v0.0.1")

	deck := deck.NewDeck()

	deck.Print()
}
