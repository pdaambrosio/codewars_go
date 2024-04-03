package main

import (
	"fmt"
	"strings"
)

// The `High` function takes a string of words, calculates the score of each word based
// on the sum of character values, and returns the word with the highest score.
func High(s string) string {
	words := strings.Split(s, " ")
	bestScore := 0
	bestWord := ""

	for _, word := range words {
		score := 0
		for _, char := range word {
			score += int(char) - 96
		}

		if score > bestScore {
			bestScore = score
			bestWord = word
		}
	}

	return bestWord
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
	fmt.Println(High("what time are we climbing up the volcano"))
	fmt.Println(High("take me to semynak"))
}
