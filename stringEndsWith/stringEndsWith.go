package  main

import (
	"fmt"
)

// The solution function checks if a given string ends with a specific ending string.
func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}

	return str[len(str)-len(ending):] == ending
}

func main() {
	fmt.Println(solution("abcde", "cde"))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("a", "z"))
	fmt.Println(solution("", "t"))
}

/*
The best solutions in codewars:
	package kata
	import "strings"
	func solution(str, ending string) bool {
		return strings.HasSuffix(str, ending)
	}
and
	package kata
	func solution(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
	}
*/