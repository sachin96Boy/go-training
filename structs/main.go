package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// creeate a struct that represent a person
	// person will have a first name and a last name
	// create a variable of type person and assign a value to it
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "sachinsupunthaka96@gmail.com",
			zipCode: 12345,
		}}
	fmt.Println(alex)
}
