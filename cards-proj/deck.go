package main

import (
	"fmt"
	"os"
	"strings"
)

//  crate a new kind of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// declare the deal func
func deal(d deck, handSize int) (deck, deck) {
	// return two values from the deal function
	// 1. a deck of cards
	// 2. the remaining cards in the deck

	// here, we are taking from start to handsize and assigning it to hand
	// and the remaining cards to remainingCards
	return d[:handSize], d[handSize:]
}

// turn the deck to string
func (d deck) toString() string {
	// convert the deck to string
	// return a string
	return strings.Join([]string(d), ",")
}

// save the deck to a file
func (d deck) saveToFile(filename string) error {
	// return an error
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// give deck from a file
func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	// if there is an error, log the error and exit the program
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(data), ",")
	return deck(s)
}
