package main

import (
	"fmt"
)

// The SquareSum function calculates the sum of the squares of all the numbers in a given slice.
func SquareSum(numbers []int) (square int) {
	for _, number := range numbers {
		square += number * number
	}

	return square
}

func main () {
	fmt.Println(SquareSum([]int{1, 2, 2}))
	fmt.Println(SquareSum([]int{1, 2}))
	fmt.Println(SquareSum([]int{5, 3, 4}))
}
