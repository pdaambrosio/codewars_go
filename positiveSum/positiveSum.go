package main

import (
	"fmt"
)

// The PositiveSum function takes in an array of integers and returns the sum of all positive numbers
// in the array.
func PositiveSum(numbers []int) int {
	var sum int = 0

	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}

	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}))
	fmt.Println(PositiveSum([]int{}))
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5}))
	fmt.Println(PositiveSum([]int{-1, 2, 3, 4, -5}))
}