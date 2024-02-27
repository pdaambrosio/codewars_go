package main

import (
	"fmt"
	"strings"
)

// The ToCamelCase function converts a string with words separated by hyphens or underscores into camel
// case format.
func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})

	for i, word := range words {
		if i != 0 {
			words[i] = strings.Title(word)
		}
	}

	return strings.Join(words, "")
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
	fmt.Println(ToCamelCase("the-stealth-warrior"))
}

// The best solution in codewars is:
// package kata

// import (
// 	"regexp"
// 	"strings"
// )

// func ToCamelCase(s string) string {
// 	words := regexp.MustCompile("-|_").Split(s, -1)

// 	for i, w := range words[1:] {
// 		words[i+1] = strings.Title(w)
// 	}

// 	return strings.Join(words, "")
// }