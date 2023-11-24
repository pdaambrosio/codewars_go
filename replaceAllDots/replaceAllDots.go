package main

import (
	"fmt"
	"regexp"
)

// The function replaces all occurrences of dots in a string with hyphens.
func ReplaceDots(str string) string {
	return regexp.MustCompile(`\.`).ReplaceAllString(str, "-")
}

func main() {
	fmt.Println(ReplaceDots("one.two.three"))
	fmt.Println(ReplaceDots("one.two.three.four.five"))
	fmt.Println(ReplaceDots("one.two.three.four.five.six.seven.eight.nine.ten"))
}
