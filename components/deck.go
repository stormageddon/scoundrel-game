package components

import (
	"errors"
	"fmt"
	"math/rand"
)

type Suits string

type Card struct {
	Name  string
	Value int
	Suit  Suits
}

type Deck []Card

const (
	Spade   Suits = "Spades"
	Heart         = "Hearts"
	Club          = "Clubs"
	Diamond       = "Diamonds"
)

//var currentDeck []Card = []Card{}

func NewDeck() (Deck, error) {
	deck := []Card{}
	for _, suit := range []Suits{Spade, Heart, Club, Diamond} {
		for i := 1; i <= 13; i++ {
			value := 0
			name := ""

			switch i {
			case 1:
				name = "Ace"
				value = 11
			case 11:
				name = "Jack"
				value = 10
			case 12:
				name = "Queen"
				value = 10
			case 13:
				name = "King"
				value = 10
			default:
				name = fmt.Sprintf("%v", i)
				value = i
			}

			card := Card{Name: fmt.Sprintf("%v of %s", name, suit), Value: value, Suit: suit}
			deck = append(deck, card)
		}
	}

	//currentDeck = deck
	return deck, nil
}

// func GetDeck() *[]Card {
// 	return &currentDeck
// }

func (d *Deck) Empty() {
	*d = []Card{}
}

func (d *Deck) Draw() (*Card, error) {
	if d == nil || len(*d) == 0 {
		return nil, errors.New("deck is empty")
	}

	card := &(*d)[0]
	*d = (*d)[1:]
	return card, nil
}

func (d *Deck) Shuffle() {
	n := len(*d)
	for i := range n {
		randIndex := rand.Intn(n)

		(*d)[i], (*d)[randIndex] = (*d)[randIndex], (*d)[i]
	}

}
