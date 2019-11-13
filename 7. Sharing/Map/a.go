package main

import "fmt"

type Vertex struct {
	Name      string
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Long: -74.39967,
		Name: "Hello",
		Lat:  40.68433,
	}
	fmt.Println(m["Bell Labs"])

	m["Jakarta"] = Vertex{
		"Hello juga", 3.14, 22.33,
	}
	fmt.Println(m["Jakarta"])

	fmt.Println("=======================")
	for index, value := range m {
		fmt.Println(index, value)
	}
}
