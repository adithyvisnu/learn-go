package main

import "fmt"

func main() {
	// primes := [6]int{}
	// fmt.Println(primes)
	// primes[1] = 12
	// fmt.Println(primes)

	// primes[2] = "string"
	// fmt.Println(primes)
	// var s = primes[1:4]
	// s ngambil dari variabel primes
	// 1 = index mulai
	// 4 = index akhir tapi gak diambil value nya
	// var s = primes[:3]
	// fmt.Println(s)

	// kartu := []string{"padung", "lope", "keriting", "ketupat"}
	// fmt.Println(kartu)

	// This can't be happening too
	// primes[6] = 0
	// fmt.Println(primes)

	// kartu[5] = "lain"
	// fmt.Println(kartu)

	numbers := make([]int, 3, 5)
	// for index := 0; index < len(numbers); index++ {

	// }

	for index, value := range numbers {
		numbers[index] = value + index
	}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	fmt.Println(numbers[2:4])
	// numbers[4] = 12
	// fmt.Println(numbers)
	numbers = append(numbers, 3)
	fmt.Println(numbers)

	// numbers = append(numbers[:1], numbers[1:0])
	// fmt.Println(numbers)

	memory := numbers[2:8]
	fmt.Println(cap(memory))
	memory[2] = 10
	fmt.Println(memory)

	// fmt.Println(cap(memory))

	// g := cap(memory)
	// fmt.Println(g)
	// for index := 0; index < g; index++ {
	// 	memory[index] = index
	// 	fmt.Printf("%d", memory[index])
	// }
	// fmt.Println(memory)

	// for index, value := range cap(memory) {
	// 	memory[index] = value + index
	// }
	// fmt.Println(memory)
	// numbers[5] = 10
	// fmt.Println(numbers[2:5])
	// numbers[3] = 12
	// fmt.Println(numbers)
}
