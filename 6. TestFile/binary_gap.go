package main

import "fmt"

func main() {
	var input = 123
	var s = fmt.Sprintf("%b", input)

	// startCount:
	for index, charString := range s {
		fmt.Println(index, charString)
	}
	return
}
