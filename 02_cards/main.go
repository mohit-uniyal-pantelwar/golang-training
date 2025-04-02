package main

import "fmt"

// short assignment cannot be used in global scope. It must be used inside a function.
// globalScopeShortAssignment := 10

func main() {
	//short variable declarations
	x := "stringVariable"
	fmt.Println(x)

	var y string = "normal declaration with initialization"
	fmt.Println(y)
}
