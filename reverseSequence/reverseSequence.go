package main

import (
	"fmt"
)

// The reverseSequence function takes an integer as input and returns a sequence of numbers in reverse
// order.
func reverseSequence(number int) []int {
	result := make([]int, number)
	for i := 0; i < number; i++ {
		result[i] = number - i
	}

	return result
}

// The main function calls the reverseSequence function and prints the result.
func main() {
	fmt.Println(reverseSequence(5))
}
