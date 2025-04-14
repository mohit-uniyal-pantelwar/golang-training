package main

import (
	"fmt"
	"sync"
	"time"
)

func channel_num(number chan int) {

}

func channel_text(text chan string) {

}

func consumeChannelData(number chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println(<-number)
}

func reduceCounter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	wg.Done()
}

func main() {

	//1. read or write to/from empty channel

	// num := make(chan int)
	// text := make(chan string)

	// go channel_num(num)
	// go channel_text(text)

	// select {
	// case channel_1 := <-num:
	// 	fmt.Println("Data: ", channel_1)
	// case channel_2 := <-text:
	// 	fmt.Println("Data: ", channel_2)
	// }

	//2. sending data to full channel

	// num := make(chan int, 2)
	// num <- 1
	// num <- 2
	// num <- 3

	//3. check
	// c := make(chan int)

	// go consumeChannelData(c)
	// c <- 10
	// fmt.Println("not waiting")
	// close(c)

	// time.Sleep(5 * time.Second)

	//4. Iterating over channel
	// c := make(chan int, 5)
	// for i := 0; i < 5; i++ {
	// 	c <- i
	// }
	// close(c)

	// for value := range c {
	// 	fmt.Println(value)
	// }

	//5. wait groups

	var wg sync.WaitGroup

	wg.Add(2)
	go reduceCounter(&wg)
	go reduceCounter(&wg)

	wg.Wait()
	fmt.Println("No blocking...")

}
