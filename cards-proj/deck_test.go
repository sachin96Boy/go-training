package main

import (
	"os"
	"testing"
)

// in go we have to make sure that the test file is in the same package as the file we are testing
// we have to manually write and run tests based on the approach
// that we wrot functions, so based on different use cases, write different test functions

// //////////////////////////////////////////////////
//  run the test
//  go test
// //////////////////////////////////////////////////

// //////////////////////////////////////////////////
// lets look in to different functions we have and write tests for hem
//  1. newDeck()
//  ////////////////
//  for this we defined func TestNewDeck() below
//  **** Code to make sure that a deck is created with X number of cards
//  **** Code to make sure that the first card is Ace of Spades
// **** Code to make sure that the last card is Four of Clubs
// ////////////////
//  2. saveToFile()
//  ////////////////
//  **** Code to make sure that a file is created when we call the saveToFile() function
//  **** Code to make sure that the contents in the file are what we expect them to be
// ////////////////
//  3. newDeckFromFile()
//  ////////////////
//  **** Code to make sure that a deck is created with X number of cards
//  **** Code to make sure that the first card is Ace of Spades
// **** Code to make sure that the last card is Four of Clubs
// ////////////////
// //////////////////////////////////////////////////

// //////////////////////////////////////////////////
//  we pass t*testing.T as a parameter to the test function
//  t is a type of testing.T
//  t has a method called Errorf()
//  Errorf() takes a format string and a list of values
//  %v is a placeholder for the value we pass in
// //////////////////////////////////////////////////

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
