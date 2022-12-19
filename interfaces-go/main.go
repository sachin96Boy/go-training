package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	// //////////////////////////////////////////////////
	// create a new instance of englishBot
	eb := englishBot{}
	// create a new instance of spanishBot
	sb := spanishBot{}
	// //////////////////////////////////////////////////
	// call printGreeting function
	printGreeting(eb)
	printGreeting(sb)
	// //////////////////////////////////////////////////

}

// //////////////////////////////////////////////////
// ceate getGreeting method for englishBot
func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi There!"
}

// ceate getGreeting method for spanishBot
func (spanishBot) getGreeting() string {
	// very custom logic for generating an spanish greeting
	return "Hola!"
}

// //////////////////////////////////////////////////

// //////////////////////////////////////////////////
// create a printGreeting function
// this function will take a type of bot as a parameter
// and it will return a string
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// create a bot interface
// this interface will have one method called getGreeting
// and it will return a string
type bot interface {
	getGreeting() string
}

// advanced type declaration on interfaces
// type bot interface {
// 	getGreeting(string, int) (string, error)
// }
//  here we can define the method signature
//  getGreeting(string, int) (string, error)
//  this method will take two parameters
//  one is a string and another is an int
//  and it will return a string and an error
// //////////////////////////////////////////////////

// 1) interfaces are not generic types
// 2) interfaces are implicit
// 3) interfaces are satisfied implicitly
// 4) interfaces are a contract to help us manage types
// 5) interfaces are tough to understand, first is to understand the concept of polymorphism
//
