package main

import (
	"fmt"
)

// The CreatePhoneNumber function formats an array of 10 numbers into a phone number string.
func CreatePhoneNumber(numbers [10]uint) string {
	format := "(%d%d%d) %d%d%d-%d%d%d%d"
	return fmt.Sprintf(format, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}