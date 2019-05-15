package main

import "fmt"

func main() {
	var fruits []string
	fruits = []string{"Apple", "Banana", "Mango", "Jack Fruit"}

	// just simple print a fruits slices
	fmt.Println(fruits)

	// what does 0:2 meant?
	// the colon was a separator for index number
	// that we will get the subsets from the slices
	// 0 (first number) means get the subset from 0 until
	// 2 (second number after colon) means get anything
	//    between 0 and 2 without index 2
	//    so it will produce [Apple Banana]
	fmt.Println(fruits[0:2])

	// what other this declaration?!
	// it tells go that we had a very early index (which is 0)
	// to be a default first number for getting subsets
	// produce [Apple Banana Mango]
	fmt.Println(fruits[:3])

	// oh i know this things
	// it tells go that we had a very last index (which is 3 in this case)
	// to be a default second number for getting subsets
	// produce [Banana Mango Jack Fruit]
	fmt.Println(fruits[1:])
}
