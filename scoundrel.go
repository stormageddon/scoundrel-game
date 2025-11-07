package main

import (
	"fmt"

	"caputo.io/scoundrel/components"
)

const (
	HandSize = 7
)

var hand []components.Card

func main() {
	deck, err := components.NewDeck()
	if err != nil {
		fmt.Println(err)
		return
	}

	for range HandSize {
		card, err := deck.Draw()

		if err != nil {
			fmt.Println(err)
			return
		}

		if card != nil {
			hand = append(hand, *card)
		}
	}
	fmt.Println(hand)
	fmt.Println(deck)
}
