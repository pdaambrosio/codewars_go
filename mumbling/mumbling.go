package main

import (
	"fmt"
	"strings"
)

// The `Accum` function takes a string as input and returns a modified version of the string where each
// character is repeated a number of times equal to its index, with the first occurrence capitalized
// and the rest in lowercase, and each character separated by a hyphen.
func Accum(s string) string {
	var result []string
	for i, c := range s {
		result = append(result, strings.ToUpper(string(c))+strings.Repeat(strings.ToLower(string(c)), i))
	}

	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("abcd"))
	fmt.Println(Accum("RqaEzty"))
	fmt.Println(Accum("cwAt"))
}
