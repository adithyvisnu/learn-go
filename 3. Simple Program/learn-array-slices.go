// To run this program use this command
// 		go run learn-array-slices.go learn-function.go
package main

import "fmt"

func main() {
	// array : fixed length
	// slices : an array that can growth and shrink
	//          slices should implements the same type
	//          of data structure of value
	fmt.Println("Array and Slices")

	// declare a new slices of type string
	cards := []string{"Ten of Spades", newCard()}
	fmt.Println(cards)

	cards = append(cards, "Six of Clovers")
	// fmt.Println(cards)

	// iterate over slices
	// i : meant to be index
	// card : meant to be value of the index
	//        card just a variable that contains a value
	// cards : slices above and loop over it
	// inside brackets : prints the index and value based on type
	for i, card := range cards {
		fmt.Printf("%d -> %s, ", i, card)
	}
}
