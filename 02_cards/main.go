package main

func main() {
	// Different ways of declaring variables
	// var card string = "Ace of Spades"

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
