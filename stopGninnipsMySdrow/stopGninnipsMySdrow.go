package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverse(word)
		}
	}

	return strings.Join(words, " ")
}

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
