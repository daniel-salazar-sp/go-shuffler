package deck

import (
	"fmt"
	"math/rand"
)

/*
 * Structs
 */
type Card struct {
	suit string
	rank string
}

type Deck struct {
	cards []Card
}

var Suits = []string{"C", "D", "H", "S"}
var Ranks = []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K"}

/*
 * Functions
 */
func NewDeck() Deck {
	var cards []Card

	for suit := 0; suit < 4; suit++ {
		for rank := 0; rank < 13; rank++ {
			cards = append(cards, Card{suit: Suits[suit], rank: Ranks[rank]})
		}
	}

	return Deck{cards: cards}
}

func (card Card) ToString() string {
	return fmt.Sprintf("%s of %s", card.rank, card.suit)
}

func (deck *Deck) Print() {
	for _, card := range deck.cards {
		fmt.Println(card.ToString())
	}
}

func (deck *Deck) Shuffle() {
	for i := (len(deck.cards) - 1); i > 0; i-- {
		j := rand.Intn(i)

		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}

func Deal(deck Deck, handSize int) (Deck, Deck) {
	var hand Deck

	for i := 0; i < handSize; i++ {
		index := rand.Intn(len(deck.cards))

		hand.cards = append(hand.cards, deck.cards[index])
		deck.cards = append(deck.cards[:index], deck.cards[index + 1:]...)
	}

	return hand, deck
}

