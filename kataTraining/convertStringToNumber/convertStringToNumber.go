package main

import (
	"fmt"
	"strconv"
)

// The function "convertStringToNumber" converts a string to an integer in Go.
func convertStringToNumber(value string) int {
	convResult, _ := strconv.Atoi(value)
	return convResult
}

func main() {
	fmt.Println(convertStringToNumber("1234"))
	fmt.Println(convertStringToNumber("605"))
	fmt.Println(convertStringToNumber("1405"))
	fmt.Println(convertStringToNumber("-7"))
}
