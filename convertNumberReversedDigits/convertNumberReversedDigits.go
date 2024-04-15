package main

import (
	"fmt"
)

// The function Digitize takes an integer input and returns a slice of its digits in reverse order.
func Digitize(n int) []int {
	var result []int

	for {
		result = append(result, n % 10)
		n = n / 10

		if n == 0 {
			return result
		}

	}
}

func main() {
	fmt.Println(Digitize(12345))
	fmt.Println(Digitize(987654321))
	fmt.Println(Digitize(0))
}
