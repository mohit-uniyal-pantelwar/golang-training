package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	fmt.Println(colors)

	colors["red"] = "#ff0000"
	colors["green"] = "#4bf745"
	colors["white"] = "#ffffff"

	for key, value := range colors {
		fmt.Printf("%v: %v\n", key, value)
	}

}
