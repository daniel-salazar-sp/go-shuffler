package deck

import (
	"testing"
	"slices"
)

/*
 * Helpers
 */
func isValidDeck(deck Deck) bool {
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
