package main

import (
	"fmt"
)

// type Mahasiswa struct {
// 	NIM  string
// 	Nama string
// }

func main() {
	// var a [2]Mahasiswa
	// a[0] = Mahasiswa{
	// 	NIM:  "1612520088",
	// 	Nama: "Adithya Visnu Prasetyo Putra",
	// }
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a[0])
	// fmt.Println(a[1])

	var numbers = [4]int{1}

	// primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// primes[6] = 24
	// fmt.Println(primes)
	// This can't be happen because
	// declaration of array is 2
	// a[3] = "Good"
	// fmt.Println(a[3])
	// data, err := dummy(1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	// if data, err := dummy(2, errors.New("hehe")); err != nil {
	// 	fmt.Println(data)
	// 	fmt.Println(err)
	// }

	// str := ""
	// if str != nil {
	// 	fmt.Println("empty")
	// }

}

func dummy(i int, e error) (int, error) {
	return i, e
}
