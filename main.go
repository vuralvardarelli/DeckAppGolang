package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	fmt.Println("------")
	hand, remainingCards := deal(cards, 5)
	hand.print()

	fmt.Println("------")
	remainingCards.print()

	fmt.Println("------")
	fmt.Println(remainingCards.toString())

	fmt.Println("------")
	remainingCards.saveToFile("Testfile.txt")
	fmt.Println("File saved.")

	fmt.Println("------")
	cards2 := newDeckFromFile("Testfile.txt")
	fmt.Println("File read.")
	cards2.print()

	fmt.Println("------")
	cards2.shuffle()
	cards2.print()
}
