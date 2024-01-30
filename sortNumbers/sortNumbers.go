package main

import (
	"fmt"
	"sort"
)

// The function SortNumbers takes in a slice of integers and returns the sorted slice.
func SortNumbers (numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func main() {
	fmt.Println(SortNumbers([]int{1, 2, 3, 10, 5}))
	fmt.Println(SortNumbers([]int{20, 2, 10}))
	fmt.Println(SortNumbers([]int{2, 20, 10}))
}