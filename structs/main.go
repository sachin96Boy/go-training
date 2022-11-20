package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// creeate a struct that represent a person
	// person will have a first name and a last name
	// create a variable of type person and assign a value to it
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
