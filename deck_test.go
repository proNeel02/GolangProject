package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Length of 16 but got %d", len(d))
	}
}
