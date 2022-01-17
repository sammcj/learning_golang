package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// length of deck is 16 cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// first card of deck is ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// last card is four of clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// Cleanup any previous tests
	os.Remove("_decktesting")

	deck := newDeck()

	// Test saving file
	deck.saveToFile("_decktesting")

	// Test loading file
	loadedDeck := newDeckFromFile("_decktesting")

	// Test loaded length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// Cleanup
	os.Remove("_decktesting")
}
