package main

import (
	"fmt"
)

func reverseWords(str string) string {
	var rev string
	var word string
	
	for _, i := range str {
	  if i == ' ' {
		rev = rev + word + " " // Adds word and space to result
		word = "" // Empties word variable
	  } else {
		word = string(i) + word // Adds letter to temporary word variable
	  } 
	}
	
	return rev + word// reverse those words
}



func main() {
	fmt.Println(reverseWords("The quick brown fox jumps over the lazy dog."))
	fmt.Println(reverseWords("apple"))
	fmt.Println(reverseWords("a b c d"))
	fmt.Println(reverseWords("double  spaced  words"))
	fmt.Print(reverseWords("stressed desserts"))
}
