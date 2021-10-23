package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print every element of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal hands of a handSize of a deck and return 2 parts.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Return a deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saving a deck to a File
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Create a deck from a File
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// return newDeck()

		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(500)
	}

	return deck(strings.Split(string(bs), ","))

}

// Shuffle the deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
