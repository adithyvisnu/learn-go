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
	cards.print()
}
