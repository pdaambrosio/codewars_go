package main

import "fmt"

// The function takes an integer as input and returns an array containing the last digit of the input
// and the number of times the input was divided by 10 and subtracted twice the remainder.
func Seven(n int64) []int {
	count := 0
	for n > 99 {
		n = n/10 - 2*(n % 10)
		count++
	}

	return []int{int(n), count}
}

func main() {
	fmt.Println(Seven(1021))
	fmt.Println(Seven(1603))
	fmt.Println(Seven(371))
}
