package main

import (
	"fmt"
)

// The function `ArrayDiff` takes two integer arrays `a` and `b`, and returns an array containing
// elements from `a` that are not present in `b`.
func ArrayDiff(a, b []int) (result []int) {
	for _, value := range a {
		if !contains(b, value) {
			result = append(result, value)
		}
	}

	return
}

// The function "contains" checks if a given value is present in a slice of integers and returns true
// if it is found.
func contains(arr []int, value int) bool {
	for _, value_compare := range arr {
		if value_compare == value {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2}, []int{1})) // [2]
	fmt.Println(ArrayDiff([]int{1, 2, 2, 2, 3}, []int{2})) // [1, 3]
	fmt.Println(ArrayDiff([]int{1, 2, 2, 2, 3}, []int{2, 3})) // [1]
}
