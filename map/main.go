package main

import "fmt"

func main() {
	// var colours map[string]string

	// colours := make(map[string]string)

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Add to map
	// colours["cool"] = "#000000"

	// Delete from map
	// delete(colours, "white")

	// fmt.Println(colours)

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
