package main

import "fmt"

func main() {
	// declare maps in go
	// here in this example all the keys are strings and all the values are strings
	// it should be importent that all keys should be only one type and
	// all values should be only one type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// print the map
	fmt.Println(colors)
}

// maps in go
// maps are key value pairs
// maps are reference types
// maps are like dictionaries in python
// maps are like objects in javascript
// maps are like hash tables in other languages
