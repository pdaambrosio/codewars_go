package main

import (
	"fmt"
	"strings"
)

// AbbreviateName takes a full name as input and returns the abbreviated version with the
// first letter of each name capitalized and separated by a period.
func AbbreviateName(name string) string {
	names := strings.Split(name, " ")
	abbreviated := fmt.Sprintf("%s.%s", strings.ToUpper(string(names[0][0])), strings.ToUpper(string(names[1][0])))
	return abbreviated
}

func main() {
	fmt.Println(AbbreviateName("Sam Harris"))
	fmt.Println(AbbreviateName("Patrick Feeney"))
	fmt.Println(AbbreviateName("Evan Cole"))
}
