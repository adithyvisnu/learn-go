package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["Answer"] = "42"
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = "48"
	fmt.Println("The value:", m["Answer"])

	// delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// fmt.Println(m["Answer"])
	v, ok := m["Question"]
	fmt.Println("The value:", v, "Present?", ok)
}
