package main

import (
	"fmt"
)

// The function Divisors calculates the number of divisors of a given integer.
func Divisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(Divisors(1))
	fmt.Println(Divisors(10))
	fmt.Println(Divisors(11))
}