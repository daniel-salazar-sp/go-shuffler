package main

import (
	"github.com/daniel-salazar-sp/go-shuffler/deck"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	cards := deck.NewDeck()

	cards.Shuffle()
	hand, cards := deck.Deal(cards, 5)

	err = hand.SaveToFile("hand.txt")
	check(err)
	err = cards.SaveToFile("deck.txt")
	check(err)

	hand, err = deck.ReadFromFile("hand.txt")
	check(err)

	hand.Print()
}
