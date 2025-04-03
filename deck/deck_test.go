package deck

import (
	"testing"
	"slices"
)

/*
 * Helpers
 */
func isValidDeck(deck Deck) bool {
	if len(deck.cards) != 52 {
		return false
	}

	for _, card := range deck.cards {
		if !(slices.Contains(Suits, card.suit) && slices.Contains(Ranks, card.rank)) {
			return false
		}
	}

	return true
}

/*
 * Tests
 */
func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if !isValidDeck(deck) {
		t.Errorf("It's not a valid deck")
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()

	if !isValidDeck(deck) {
		t.Errorf("It's no longer a valid deck")
	}

	first := deck.cards[0]
	if first.rank == "A" && first.suit == "A" {
		t.Errorf("Deck hasn't been shuffled")
	}
}

func TestDeal(t *testing.T) {
	deck := NewDeck()
	hand, deck := Deal(deck, 5)

	if len(hand.cards) != 5 {
		t.Errorf("Hand does not have the proper size; expected %d", len(hand.cards))
	}

	if len(deck.cards) != 47 {
		t.Errorf("Deck did not have the hand's cards removed")
	}
}
