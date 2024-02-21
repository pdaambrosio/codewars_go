package main

import (
	"fmt"
	"strings"
)

// This Go function finds the length of the shortest word in a given string.
func FindShort(s string) int {
	words := strings.Split(s, " ")
	shortestWord := words[0]
	for i := 0; i < len(words); i++ {
		if len(words[i]) < len(shortestWord) {
			shortestWord = words[i]
		}
	}

	return len(shortestWord)
}


func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))
	fmt.Println(FindShort("turns out random test cases are easier than writing out basic ones"))
	fmt.Println(FindShort("Let's travel abroad shall we"))
}

// The best solution in codewars is:
// package kata

// import "strings"

// func FindShort(s string) int {
//   shortest := len(s)
//   for _, word := range strings.Split(s, " ") {
//     if len(word) < shortest {
//       shortest = len(word)
//     }
//   }
//   return shortest
// }