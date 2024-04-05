package main

import (
	"fmt"
	"strings"
	"unicode"
)

// The High function in Go takes a string of words, calculates the score for each word, and returns the
// word with the highest score.
func High(s string) string {
	words := strings.Fields(s)
	bestScore := 0
	bestWord := ""
	for _, word := range words {
		score := calculateScore(word)
		if score > bestScore {
			bestScore = score
			bestWord = word
		}
	}

	return bestWord
}

// The calculateScore function calculates the score of a word by summing the values of its letters
// based on their position in the alphabet.
func calculateScore(word string) int {
	score := 0
	for _, char := range word {
		if unicode.IsLetter(char) {
			score += int(unicode.ToLower(char) - 'a' + 1)
		}
	}

	return score
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
	fmt.Println(High("what time are we climbing up the volcano"))
	fmt.Println(High("take me to semynak"))
}
