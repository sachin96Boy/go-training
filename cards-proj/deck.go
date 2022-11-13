package main

import "fmt"

//  crate a new kind of 'deck'
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}