package main

import "fmt"

func main() {
	var name string
	name = "Bill"
	updateValue(name)
	fmt.Println(name)

	drinks := []string{
		"Aqua", "Amidis",
	}
	updateSliceValue(drinks)
	fmt.Println(drinks)
}

func updateValue(n string) {
	n = "Alex"
}

func updateSliceValue(slice []string) {
	slice[0] = "Nice Try"
}
