package main

import "fmt"

// any type that has a getGreeting function with
// string as the return value will included on
// the type bot interfaces automatically
type bot interface {
	getGreeting() string
}

type englishBot struct {
}
type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// the parameter inside this function is an interfaces
// which must be a type and has a getGreeting function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// we can't create the same function name even if
// there are different on parameters
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	// very custom logic for get greeting on english bot
	return "Hello there"
}

func (sb spanishBot) getGreeting() string {
	// very custom logic for get greeting on spanish bot
	return "Hola"
}
