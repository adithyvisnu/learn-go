package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// This can't be happening too
	// q[6] = 12
	// fmt.Println(q)

	q = append(q, 12)
	fmt.Println(q[6])

	s = append(s, struct {
		i int
		b bool
	}{17, false})

	// s = append(s, struct {
	// 	i int
	// 	b float32
	// }{17, 3.14})

	fmt.Println(s)
}
