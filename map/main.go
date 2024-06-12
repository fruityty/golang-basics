package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0001",
	}

	// var colors2 map[string]string

	// colors := make(map[string]string)

	colors["white"] = "ffffff"

	delete(colors, "green")

	fmt.Println(colors)
	printMap(colors)

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
