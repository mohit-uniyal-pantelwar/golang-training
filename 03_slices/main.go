package main

import "fmt"

func main() {
	//Array - fixed size list of homogeneous elements
	array := [5]int{1, 2, 3}
	fmt.Println(array)

	//Slice - Dynamic array which grows at runtime.
	fruits := []string{"apple", "banana", "grapes"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
