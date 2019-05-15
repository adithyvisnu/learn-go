package main

import "fmt"

func main() {
	var first, second int
	first, second = 1, 6
	fmt.Println(first, second)
	first, second = exchange(first, second)
	fmt.Println("===> ", first, second)
}

func exchange(a, b int) (int, int) {
	return b, a
}
