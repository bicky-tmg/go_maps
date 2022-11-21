package main

import (
	"fmt"
)

// func main() {
// var colors map[string]string

// colors := make(map[int]string)

// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "#00FF00",
// }

// 	colors[10] = "#ffffff"

// 	delete(colors, 10)

// 	fmt.Println(colors)
// }

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
