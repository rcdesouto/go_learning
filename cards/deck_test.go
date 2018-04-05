package main

import (
	"os"
	"testing"
)

//specify package. This includes all testing files
// capital T for tests
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected value: 52, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected: Ace of Spaces, but got: %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected: King of Clubs, but got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected value: 52, but got: %v", len(ld))
	}

	os.Remove("_decktesting")
}
