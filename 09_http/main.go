package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	bs := make([]byte, 999999)
	bytesRead, err := res.Body.Read(bs)
	fmt.Println(bytesRead)
	if err != nil && err != io.EOF {
		log.Fatal("Error: ", err)
	}
	fmt.Println(string(bs))
}
