package main

import "fmt"

func main() {
	// var : declare new variable
	// card : name of a new variable
	// string : telling go compiler
	//          only of type string
	//          that will ever be assigned
	//          to the variables
	// var card string = "Ace of Spades"
	// var declaration above is explicitly declared
	// we could use :=
	card := "Ace of Spades"
	// we no longer have to call initiation if we want to
	// reassign the variables
	card = "Five of Clovers"
	fmt.Println(card)

	var tf bool
	tf = false
	fmt.Println(tf)

	var i int
	i = 1002
	fmt.Println(i)
}
