package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Length of 16 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %s", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	const fileName = "_decktesting.txt"
	os.Remove(fileName)
	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromeFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16, got %d", len(loadedDeck))
	}
	os.Remove(fileName)
}
