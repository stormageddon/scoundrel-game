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
	Id    string
}

type Deck []Card

const (
	Spade   Suits = "Spades"
	Heart         = "Hearts"
	Club          = "Clubs"
	Diamond       = "Diamonds"
)

func NewDeck() (Deck, error) {
	deck := []Card{}
	for _, suit := range []Suits{Spade, Heart, Club, Diamond} {
		for i := 1; i <= 13; i++ {
			value := 0
			name := ""
			id := ""

			switch i {
			case 1:
				name = "Ace"
				value = 14
			case 11:
				name = "Jack"
				value = 11
			case 12:
				name = "Queen"
				value = 12
			case 13:
				name = "King"
				value = 13
			default:
				name = fmt.Sprintf("%v", i)
				value = i
			}

			if value != 10 {
				id = fmt.Sprintf("%c%c", name[0], suit[0])
			} else {
				id = fmt.Sprintf("%s%c", name, suit[0])
			}

			card := Card{Name: fmt.Sprintf("%v of %s", name, suit),
				Value: value,
				Suit:  suit,
				Id:    id}
			deck = append(deck, card)
		}
	}

	//currentDeck = deck
	return deck, nil
}

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

func (d *Deck) Remove(cardName string) {
	for i, c := range *d {
		if c.Name == cardName {
			*d = append((*d)[:i], (*d)[i+1:]...)
			break
		}
	}
}

func (d *Deck) AddToBottom(card ...Card) {
	*d = append(*d, card...)
}
