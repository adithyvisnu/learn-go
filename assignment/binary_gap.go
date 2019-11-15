package main

import "fmt"

func main() {
	BinaryGap(1025)
}

func BinaryGap(i int) {
	var s = fmt.Sprintf("%b", i)
	fmt.Println(s)
	var tmp = []int{0, 0}
	var max = 0
	for _, charString := range s {
		if tmp[0] == 0 && charString == 49 {
			tmp[0] = 1
		} else if tmp[0] == 1 && charString == 49 {
			if tmp[1] <= max {
				tmp[1] = max
				max = 0
			}
			tmp[0] = 0
		} else if tmp[0] == 1 && charString == 48 {
			max = max + 1
		}
	}

	// fmt.Println(tmp[1])
	return
}
