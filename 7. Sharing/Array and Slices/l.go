package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	str := "hello"
	// str = ["h", "e", "l", "l", "o"]
	for _, val := range str {
		fmt.Printf("%s ", string(val))
	}

	fmt.Println()
}
