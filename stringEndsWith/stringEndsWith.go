package  main

import (
	"fmt"
)

// The solution function checks if the ending of a given string matches a specified ending string.
func solution(str, ending string) bool {
	return str[len(str)-len(ending):] == ending
}

func main() {
	fmt.Println(solution("abcde", "cde"))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("a", "z"))
}