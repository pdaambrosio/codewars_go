package main

import (
	"fmt"
	"strings"
)

// Function IsUpperCase checks if a string is uppercase
func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}

// Function main is the main function of this package
func main() {
	fmt.Println(IsUpperCase("c"))
	fmt.Println(IsUpperCase("C"))
	fmt.Println(IsUpperCase("hello I AM DONALD"))
	fmt.Println(IsUpperCase("HELLO I AM DONALD"))
	fmt.Println(IsUpperCase("ACSKLDFJSgSKLDFJSKLDFJ"))
	fmt.Println(IsUpperCase("ACSKLDFJSGSKLDFJSKLDFJ"))
}
