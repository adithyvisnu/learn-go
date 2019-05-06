package main

import "fmt"

// create type called 'deck'
// which is slice of strings
// why we add type?
//     if we want to create a function based on
//     the type and what its belong
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
