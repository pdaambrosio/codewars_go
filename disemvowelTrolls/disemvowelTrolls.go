package main

import (
	"fmt"
	"strings"
)

// The Disemvowel function removes all vowels from a given string.
func Disemvowel(comment string) string {
	var vowels string = "aeiouAEIOU"

	for _, vowel := range vowels {
		comment = strings.ReplaceAll(comment, string(vowel), "")
	}

	return comment
}

func main() {
	fmt.Println(Disemvowel("This website is for losers LOL!"))
	fmt.Println(Disemvowel("No offense but,\nYour writing is among the worst I've ever read"))
	fmt.Println(Disemvowel("What are you doing?"))
}

// The best solution in the Codewars:
// package kata

// import "regexp"

// func Disemvowel(comment string) string {
//   return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
// }