package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ToAlternatingCase takes a string as input and returns a new string with each
// character converted to its opposite case.
func ToAlternatingCase(str string) string {
	var result strings.Builder

	for _, char := range str {
		if unicode.IsLower(char) {
			result.WriteRune(unicode.ToUpper(char))
		} else {
			result.WriteRune(unicode.ToLower(char))
		}
	}

	return result.String()
}

func main() {
	fmt.Println(ToAlternatingCase("hello world"))
	fmt.Println(ToAlternatingCase("HELLO WORLD"))
	fmt.Println(ToAlternatingCase("hello WORLD"))
	fmt.Println(ToAlternatingCase("HeLLo WoRLD"))
	fmt.Println(ToAlternatingCase("12345"))
	fmt.Println(ToAlternatingCase("1a2b3c4d5e"))
}