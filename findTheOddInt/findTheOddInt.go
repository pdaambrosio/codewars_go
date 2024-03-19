package main

import (
	"fmt"
)

// The FindOdd function in Go calculates the odd number that appears an odd number of times in a given
// sequence of integers.
func FindOdd(seq []int) int {
	odd := 0
	for _, v := range seq {
		odd ^= v
	}

	return odd
}

func main() {
	fmt.Println(FindOdd([]int{1, 1, 2, 2, 3, 3, 4, 4, 5}))
	fmt.Println(FindOdd([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5}))
	fmt.Println(FindOdd([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 5}))
}
