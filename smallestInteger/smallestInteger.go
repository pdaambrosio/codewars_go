package main

import (
	"fmt"
)

// The SmallestIntegerFinder function takes in an array of integers and returns the smallest integer in
// the array.
func SmallestIntegerFinder(numbers []int) int {
	var min int = numbers[0]
	for _, value := range numbers { // it's possible use sort.Ints(numbers) and return numbers[0] yet
		if value < min {
			min = value
		}
	}
	return min
}


func main() {
	fmt.Println(SmallestIntegerFinder([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}))
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100}))
}