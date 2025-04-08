package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkUrlStatus(url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	client := http.Client{}

	_, err := client.Do(req)
	return err
}

func getStatus(urls []string) {

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 5)

	for _, url := range urls {
		wg.Add(1)
		semaphore <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()
			err := checkUrlStatus(url)
			if err == nil {
				fmt.Printf("%v: Fetch successfull\n", url)
			} else {
				fmt.Printf("%v: Fetch unsuccessfull\n", url)
			}
		}()
	}

	wg.Wait()

}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.reddit.com",
		"https://www.stackoverflow.com",
		"https://www.nonexistentwebsite.xyz",
		"https://www.invalid-url.com",
	}

	getStatus(urls)

}
