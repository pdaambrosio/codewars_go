package main

import (
	"fmt"
)

// The SortByLength function sorts an array of strings based on their lengths in ascending order.
func SortByLength(array []string) []string {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if len(array[i]) > len(array[j]) {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func main() {
	fmt.Println(SortByLength([]string{"beg", "life", "i", "to"}))
	fmt.Println(SortByLength([]string{"", "moderately", "brains", "pizza"}))
	fmt.Println(SortByLength([]string{"longer", "longest", "short"}))
}
