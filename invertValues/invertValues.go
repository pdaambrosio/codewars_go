package main

import (
	"fmt"
)

// The function "invertValues" takes an array of integers as input and returns a new array with all the
// values negated.
func invertValues(arrNum []int) []int {
	result := make([]int, len(arrNum))
	for i, v := range arrNum {
		result[i] = -v
	}
	return result
}

// The main function calls the invertValues function with different arrays of integers and prints the
// results.
func main() {
	fmt.Println(invertValues([]int{1, 2, 3, 4, 5}))
	fmt.Println(invertValues([]int{-1, -2, -3, -4, -5}))
	fmt.Println(invertValues([]int{1, -2, 3, -4, 5}))
}
