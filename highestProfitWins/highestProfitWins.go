package main

import (
	"fmt"
)

// The MinMax function takes an array of integers and returns an array containing the minimum and
// maximum values in the input array.
func MinMax(arr []int) [2]int {
	var min, max int = arr[0], arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return [2]int{min, max}
}

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinMax([]int{1, 2, -3, 4, 5}))
	fmt.Println(MinMax([]int{1, 9, 3, 4, -5}))
}
