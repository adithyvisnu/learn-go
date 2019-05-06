// To run this program use this command
// 		go run main.go deck.go
package main

func main() {
	cards := deck{"Ace of Spades", "Five of Clovers"}
	cards = append(cards, "Ten of Hearts")
	cards.print()
}
