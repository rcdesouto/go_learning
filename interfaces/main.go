package main

import (
	"fmt"
)

// Only structs with the function getGreeting()
// Can be apart of this club that is known as Bot
// If there are other methods defined in the bot
// and the stucts do not have these methods then they can not be a member
// of the bot club
type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
