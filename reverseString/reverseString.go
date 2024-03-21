package main

import (
	"fmt"
)

// The Solution function in Go reverses a given string by iterating through its characters and
// concatenating them in reverse order.
func Solution(word string) string {
	var rev string

	for _, i := range word {
		rev = string(i) + rev
	}

	return rev
}

func main() {
	fmt.Println(Solution("world"))
	fmt.Println(Solution("hello"))
	fmt.Println(Solution("apple"))
}