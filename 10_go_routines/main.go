package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.stackoverflow.com", "https://www.golang.org", "https://www.amazon.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}

	for value := range c {
		go checkLink(value, c)
	}
}

func checkLink(url string, c chan string) {
	time.Sleep(time.Second * 10)
	_, err := http.Get(url)
	if err == nil {
		fmt.Printf("%v is up\n", url)
	} else {
		fmt.Printf("%v is down\n", url)
	}
	c <- url
}
