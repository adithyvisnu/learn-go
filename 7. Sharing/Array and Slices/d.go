package main

import "fmt"

func main() {
	q := []int{3, 5}
	fmt.Println(q)

	r := []int{5, 4}

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{}
	// fmt.Println(s)

	// This can't be happening too
	// q[6] = 12
	// fmt.Println(q)

	// s[0] = struct {
	// 	i int
	// 	b bool
	// }{
	// 	1, true,
	// }
	fmt.Println("==========")
	q = append(q, r...)
	fmt.Println(q)
	fmt.Println(r)

	// s = append(s, struct {
	// 	i int
	// 	b bool
	// }{17, false})

	// s = append(s, struct {
	// 	i int
	// 	b float32
	// }{17, 3.14})

	// fmt.Println(s)
}
