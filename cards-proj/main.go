package main

// import "fmt"

func main() {
	// veriable declaration
	// var card string = "Ace of Spades"

	// lets break down the declaration for more details
	// //////////////////////////////////////////////////
	//  var => variable declaration keyword
	// card => name of the variable
	// string => type of the variable
	// "Ace of Spades" => value of the variable
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// go basic types
	// string => "Hello World"
	// bool => true, false
	// int => 1, 2, 3
	// float64 => 1.2, 3.4
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// declare a variable without a type
	//  card := "Ace of Spades"	 // first time we declare a variable, we need to use the var keyword
	// card = 5
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	//  use the function we declared to givethe value of the card
	// card := newCard()
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// there are two types of arrays in go
	// 1. fixed length array => [5]string => fixed length of 5
	// 2. slice => []string => slice is a dynamic array that can grow or shrink
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	//  use slices to store multiple values
	// cards := []string{"Ace of Diamonds", newCard()}
	// append a new value to the slice
	// cards = append(cards, "Six of Spades")

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

	// //////////////////////////////////////////////////
	// iterate over the slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards)
}

// //////////////////////////////////////////////////
// define a new function

func newCard() string {
	return "Five of Diamonds"
}