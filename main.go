package main

import "fmt"

func main() {
	//var card string
	//var card string = "Test"
	//card := newCard()

	//fmt.Println(card)

	cards := []string{"new card 1", newCard()}

	cards = append(cards, "new card 3")

	for i, card := range cards { //for-foreach like
		fmt.Println(i, card)
	}

	//fmt.Println(cards)

}

func newCard() string {
	return "new card 2"
}
