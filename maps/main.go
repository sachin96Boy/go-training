package main

import "fmt"

func main() {
	// declare maps in go
	// here in this example all the keys are strings and all the values are strings
	// it should be importent that all keys should be only one type and
	// all values should be only one type

	// //////////////////////////////////////////////////
	// // declare maps in go in a different way
	// // here in this example all the keys are strings and all the values are strings
	// // it should be importent that all keys should be only one type and
	// // all values should be only one type
	emails := make(map[string]string)
	// here when we declare a map we have to use make keyword
	// and we have to pass the type of
	// the key and the type of the value
	// initialy the map will be empty
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// // assign key values
	// here since we declare keys and values as strings
	// we have to pass strings
	// we use this brackets syntax to assign key values
	// the key values should be unique and its type should be the same as the type of the key
	emails["Bob"] = "bobgmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@mikkyway.com"
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// // delete from map
	// pass the key to delete with the map name
	delete(emails, "Bob")
	// //////////////////////////////////////////////////

	// /////////////////////////////////////////////////////////
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// /////////////////////////////////////////////////////////
	// print the map
	fmt.Println(colors)
	fmt.Println(emails)
}

// maps in go
// maps are key value pairs
// maps are reference types
// maps are like dictionaries in python
// maps are like objects in javascript
// maps are like hash tables in other languages
