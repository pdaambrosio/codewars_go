package main

import (
	"fmt"
	"strings"
)

// The SpinWords function takes a string, splits it into words, and reverses any word with a length of
// 5 or more before joining the words back together.
func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverse(word)
		}
	}

	return strings.Join(words, " ")
}

// The `reverse` function in Go takes a string as input and returns the string with its characters
// reversed.
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
	fmt.Println(SpinWords("This is a test"))
	fmt.Println(SpinWords("This is another test"))
}
