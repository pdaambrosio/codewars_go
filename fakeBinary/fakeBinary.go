package main

import (
	"fmt"
)

// The function takes a string of digits and replaces each digit with either "0" or "1" based on
// whether the digit is less than 5 or not.
func FakeBin(s string) string {
	var result string

	for _, value := range s {
		if value < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}

	return result
}

func main() {
	fmt.Println(FakeBin("45385593107843568"))
	fmt.Println(FakeBin("509321967506747"))
	fmt.Println(FakeBin("366058562030849490134388085"))
}
