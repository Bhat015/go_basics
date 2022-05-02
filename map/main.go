package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	colors["white"] = "#ffffff"

	delete(colors, "green")

	fmt.Print(colors)

	// iterate over map
	printMap(colors)

}

func printMap(_map map[string]string) {
	for key, value := range _map {
		fmt.Println(key, value)
	}

}
