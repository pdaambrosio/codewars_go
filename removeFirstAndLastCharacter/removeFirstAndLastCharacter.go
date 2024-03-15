package main

import (
	"fmt"
)

// The RemoveChar function in Go removes the first and last characters from a given string.
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func main() {
	fmt.Println(RemoveChar("eloquent"))
	fmt.Println(RemoveChar("country"))
	fmt.Println(RemoveChar("person"))
}
