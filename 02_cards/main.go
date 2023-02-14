package main

import "fmt"

func main() {
	// Different ways of declaring variables
	// var card string = "Ace of Spades"

	cards := newDeck()
	cards.toString()
	fmt.Println(cards.toString())
}
