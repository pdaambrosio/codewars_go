package main

import (
	"fmt"
)

func High(s string) string {
	var highestWord string
	var highestScore int

	words := split(s)
	for _, word := range words {
		score := getScore(word)
		if score > highestScore {
			highestScore = score
			highestWord = word
		}
	}

	return highestWord
}

func main() {
	fmt.Println(High("abad"))
	fmt.Println(High("taxi"))
	fmt.Println(High("volcano"))
}
