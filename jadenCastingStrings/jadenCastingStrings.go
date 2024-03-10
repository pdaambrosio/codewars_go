package main

import (
	"fmt"
	"strings"
)

// The ToJadenCase function takes a string input and capitalizes the first letter of each word in the
// string.
func ToJadenCase(phrase string) string {
	arrayPhrase := strings.Split(phrase, " ")

	for word := range arrayPhrase {
		arrayPhrase[word] = strings.Title(arrayPhrase[word])
	}

	return strings.Join(arrayPhrase, " ")
}

func main() {
	fmt.Println(ToJadenCase("most trees are blue"))
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own"))
	fmt.Println(ToJadenCase("When I die. then you will realize"))
}
