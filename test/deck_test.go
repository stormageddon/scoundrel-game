package test

import (
	"reflect"
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

	l := len(testDeck)

	if testDeck == nil || len(testDeck) != 52 {
		t.Errorf("Expected size %v but deck was size %v", 52, l)
	}
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
	originalDeck := make(components.Deck, len(testDeck))
	copy(originalDeck, testDeck)
	originalLength := len(originalDeck)
	testDeck.Shuffle()

	if len(testDeck) != 52 {
		t.Errorf("Length of deck (%d) after shuffle was not %d", len(testDeck), originalLength)
	}

	if reflect.DeepEqual(originalDeck, testDeck) {
		t.Error("Deck has same order after shuffling")
	}
}

func TestRemoveCard(t *testing.T) {
	setup(t)

	testDeck.Remove("Ace of Spades")

	if len(testDeck) != 51 {
		t.Errorf("Expected deck size %v but got %v", 51, len(testDeck))
	}

	for _, card := range testDeck {
		if card.Name == "Ace of Spades" {
			t.Error("Ace of Spades was not correctly removed")
		}
	}
}
