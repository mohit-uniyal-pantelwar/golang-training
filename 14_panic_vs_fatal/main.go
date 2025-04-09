package main

import "fmt"

func function2() {
	defer fmt.Println("Function 2 ends")
	fmt.Println("Function 2 begins")
	panic("Panic in function 2") // executed at very last when all the unwinding is done.
}

func function1() {
	defer fmt.Println("Function 1 Ends")
	fmt.Println("Function 1 begins")
	function2()
	fmt.Println("Function 1 body")
}

func main() {

	function1()

}
