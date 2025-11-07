package test

import (
	"fmt"
	"testing"

	"caputo.io/scoundrel/components"
)

var testDeck components.Deck

func setup(t *testing.T) {
	d, err := components.NewDeck()
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	testDeck = d
	t.Cleanup(func() { testDeck.Empty() })
}

func TestCreateDeck(t *testing.T) {
	setup(t)

	//d := deck.Create()
	l := len(testDeck)

	if testDeck == nil || len(testDeck) != 52 {
		t.Errorf("Expected size %v but deck was size %v", 52, l)
	}

	//t.Cleanup(func() { deck.Empty() })
}

func TestDrawFromEmptyDeck(t *testing.T) {
	testDeck.Empty()
	_, err := testDeck.Draw()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestDrawCard(t *testing.T) {
	setup(t)

	card, err := testDeck.Draw()
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	if card == nil {
		t.Errorf("Expected card but got nil")
	}
	if len(testDeck) != 51 {
		t.Errorf("Expected deck size %v but got %v", 51, len(testDeck))
	}
}

func TestShuffleDeck(t *testing.T) {
	setup(t)
	originalDeck := make(components.Deck, 52)
	copy(originalDeck, testDeck)
	originalLength := len(originalDeck)
	testDeck.Shuffle()

	fmt.Printf("TestDeck: %v, OriginalDeck: %v", testDeck, originalDeck)

	if len(testDeck) != 52 {
		t.Errorf("Length of deck (%d) after shuffle was not %d", originalLength, len(testDeck))
	}
	numSamePlace := 0

	for i := range len(testDeck) {
		if testDeck[i] == (originalDeck)[i] {
			numSamePlace += 1
		}
	}

	if numSamePlace == originalLength {
		t.Error("Deck has same order after shuffling")
	}
}
