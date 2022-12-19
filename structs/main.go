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
	// fmt.Println(alex)
	// print the value of the person

	// //////////////////////////////////////////////////
	// update first name of the person
	// alexPointer := &alex
	// alexPointer.updateName("Sachin")
	// alex.print()
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// update first name of the person without explicitly using pointer
	// update first name of the person
	alex.updateName("Sachin")
	alex.print()
	// //////////////////////////////////////////////////

	// //////////////////////////////////////////////////
	// alex.print()
	// //////////////////////////////////////////////////
}

// this method didn't give the result so updated method down below after this method
// func updateName(newFirstName string, p person) {
// 	p.firstName = newFirstName
// }

// pointers in GO
// here we define new struct of type person to alex
// that struct will be stored in the memory
// and the memory address of that struct will be stored in alex
// so alex is a pointer to that struct

// this method will give the result
// /////////////////////////////////////////////////////
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// /////////////////////////////////////////////////////

// /////////////////////////////////////////////////////
// first look at these new variables we used
// & variable - gives the memory address of the value this variable is pointing at
// * pointer - gives the value this memory address is pointing at

// /////////////////////////////////////////////////////
// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }
// /////////////////////////////////////////////////////
// in the reseiver we use * where the type is declared and not the value
//  here *person is a type description. it means we are working with a pointer to a person
//  here *pointerToPerson is a operator. it means we want to manipulate the value the pointer is referencing
// /////////////////////////////////////////////////////

// //////////////////////////////////////////////////

func (p person) print() {
	fmt.Printf("%+v", p)
}
