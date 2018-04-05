package main

// Part of package main
// For Executable programs there must be a package main
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three",
		"Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swapping positions
		// Take item from newPosition and put it in position i
		// Take item from i posision and put it in newPosition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(404)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
