package deck

import (
	"os"
	"bufio"
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

/*
 * I/O
 */
func (deck Deck) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if (err != nil) {
		return err
	}

	defer file.Close()

	writeBuf := bufio.NewWriter(file)

	for _, card := range deck.cards {
		_, err := writeBuf.WriteString(fmt.Sprintf("%s%s\n", card.suit, card.rank))
		if (err != nil) {
			return err
		}
	}

	writeBuf.Flush()
	return nil
}

func ReadFromFile(filename string) (Deck, error) {
	file, err := os.Open(filename)
	if (err != nil) {
		return Deck{}, err
	}

	defer file.Close()

	var cards []Card
	readBuf := bufio.NewScanner(file)

	for readBuf.Scan() {
		line := readBuf.Text()

		cards = append(cards, Card{suit: string(line[0]), rank: string(line[1:])})
	}

	return Deck{cards: cards}, nil
}
