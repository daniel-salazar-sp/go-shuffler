package deck

import (
	"fmt"
)

/*
 * Structs
 */
type Card struct {
	suit string
	rank string
}

type Deck struct {
	cards [52]Card
}

var	Suits = []string{"A", "C", "X", "H"}
var Ranks = []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K"}

/*
 * Functions
 */
func NewDeck() Deck {
	var cards [52]Card

	for rank := 0; rank < 13; rank++ {
		for suit := 0; suit < 4; suit++ {
			cards[rank * 4 + suit] = Card{suit: Suits[suit], rank: Ranks[rank]}
		}
	}

	return Deck{cards: cards}
}

func (deck Deck) Print() {
	for index, card := range deck.cards {
		fmt.Printf("%d: %s of %s\n", index, card.rank, card.suit)
	}
}
