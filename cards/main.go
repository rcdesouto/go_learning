package main

// File is names main.go as it is required for go to create
// a executable file during compile. I connect this with
// node.js best practice of naming main file running as index.js

import (
	"fmt"
)

// this needs to be named main to run.
func main() {
	// created a deck "type" in the deck.go
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// Running a deal function returns two deck types
	// Definition of this can be found in deck.go
	hand, remaingCards := deal(cards, 5)

	// Show the hand that was created from above
	fmt.Println("The Hand")
	hand.print()

	// Save Cards to a file in the current directory
	remaingCards.saveToFile("my_cards")

	// Load a deck from a file
	// oldDeck := newDeckFromFile("my_cards")
	// oldDeck.print()
}
