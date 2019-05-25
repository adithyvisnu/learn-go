package main

import "fmt"

func main() {
	colors := map[string]string{
		"white": "#FFFFFF",
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	fmt.Println(colors)
	updateMapValue(colors)
	printMap(colors)
}

func updateMapValue(m map[string]string) {
	m["red"] = "Just RED!"
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("The", k, "color has a hexa value", v)
	}
}
