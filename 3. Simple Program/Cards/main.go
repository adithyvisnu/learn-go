// To run this program use this command
// 		go run main.go deck.go
package main

func main() {
	// this one the first approach for learning
	// the function receiver method
	// it commented due to creating a new code
	// for initializing a deck itself
	// ==============================================
	// cards := deck{"Ace of Spades", "Five of Clubs"}
	// cards = append(cards, "Ten of Hearts")
	// cards.print()

	// this is a new code for declaring a deck
	// declaring deck with the new function
	cards := newDeck()
	// cards.print()

	// this is a multiple return values function
	hands, _ := deal(cards, 11)
	// fmt.Println(hands)
	// fmt.Println(remainingCards)

	// fmt.Println(hands.toString())

	hands.saveToFile("mycards")

	// why this variable below cannot use the function
	// saveToFile() or toString() which has a []string as a param?
	// because it is not a deck type
	// var anotherSameType []string
	// anotherSameType = []string{"a", "b", "cd"}
	// anotherSameType

	deck := newDeckFromFile("mycards")
	deck.print()
}
