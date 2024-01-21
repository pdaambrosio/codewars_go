package main

import (
	"fmt"
	"sort"
)

// The function takes two strings, concatenates them, sorts the resulting string, removes any duplicate
// characters, and returns the final string.
func TwoToOne(s1 string, s2 string) string {
	s := s1 + s2
	slice := []byte(s)
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	s = string(slice)

	var result string
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] != s[i-1] {
			result += string(s[i])
		}
	}

	return result
}

func main() {
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
	fmt.Println(TwoToOne("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"))
}

/*
The best solutions in CodeWars:
	package kata

	import (
	"sort"
	"strings"
	)

	func TwoToOne(s1 string, s2 string) string {
	chars := strings.Split(s1 + s2, "")
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
		result = result + chr
		}
	}
	return result
	}
and
	package kata

	import "strings"

	func TwoToOne(s1 string, s2 string) string {
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s1+s2, char) {
		result += char
		}
	}
	return result
	}
*/