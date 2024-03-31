package main

import (
	"fmt"
	"strings"
)

// The function AbbrevName takes a full name as input and returns the abbreviated version with the
// first letter of each name capitalized and separated by a period.
func AbbrevName(name string) string {
	splitName := strings.Split(name, " ")
	formatName := fmt.Sprintf("%s.%s", strings.ToUpper(string(splitName[0][0])), strings.ToUpper(string(splitName[1][0])))
	return formatName
}

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
	fmt.Println(AbbrevName("Patrick Feeney"))
	fmt.Println(AbbrevName("Evan Cole"))
}
