package main

import (
	"fmt"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work...")
		}
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	// myChannel := make(chan string)
	// anotherChannel := make(chan string)

	// go func() {
	// 	myChannel <- "data1"
	// }()

	// go func() {
	// 	anotherChannel <- "data2"
	// }()

	// select {
	// case messageFromMyChannel := <-myChannel:
	// 	fmt.Println(messageFromMyChannel)
	// case messageFromAnotherChannel := <-anotherChannel:
	// 	fmt.Println(messageFromAnotherChannel)
	// }

	// **********For-Select loop*************

	//Example 1
	// chars := []rune{'A', 'B', 'C'}

	// charChannel := make(chan rune, 3)

	// for _, char := range chars {
	// 	charChannel <- char
	// }

	// close(charChannel)

	// for result := range charChannel {
	// 	fmt.Println(string(result))
	// }

	//*************Done Channel*************

	// done := make(chan bool)

	// go doWork(done)

	// time.Sleep(time.Second * 3)

	// close(done)

	// time.Sleep(time.Hour * 200)

	//*******pipeline**********

	nums := []int{1, 2, 3, 4, 5, 6}

	dataChannel := sliceToChannel(nums)
	finalChannel := sq(dataChannel)

	for n := range finalChannel {
		fmt.Println(n)
	}
}
