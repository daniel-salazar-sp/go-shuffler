package main

import (
	"github.com/daniel-salazar-sp/go-shuffler/deck"
)

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()

	hand, cards := deck.Deal(cards, 5)
	hand.Print()
}
