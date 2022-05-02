package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Println(cards.toString())

	//save to file
	cards.saveToFile("my_cards")

	//read from file
	cards2 := newDeckFromFile("my_cards")

	cards2.print()
}
