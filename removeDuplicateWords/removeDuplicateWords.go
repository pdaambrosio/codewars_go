package main

import (
	"fmt"
	"strings"
)

// The function RemoveDuplicateWords takes a string as input and returns a new string with duplicate
// words removed.
func RemoveDuplicateWords(str string) string {
	var words []string
	for _, word := range strings.Split(str, " ") {
		if !strings.Contains(strings.Join(words, " "), word) {
			words = append(words, word)
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta"))
	fmt.Println(RemoveDuplicateWords("my cat is my cat fat"))
	fmt.Print(RemoveDuplicateWords("hello hello world"))
}
