package main

import (
	"fmt"
)

func CountPositivesSumNegatives(numbers []int) []int {
	var count int = 0
	var sum int = 0

	for _, number := range numbers {
		if number > 0 {
			count++
		} else if number < 0 {
			sum += number
		}
	}

	return []int{count, sum}
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(numbers))
}
