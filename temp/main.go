package main

import "fmt"

func main() {
	mp := map[string]int{}
	addValue(mp, "alice")
	fmt.Println(mp)
}

func addValue(mp map[string]int, key string) {
	mp[key] = 3
}
