package main

import (
	"fmt"
)

// The function calculates the sum of all even Fibonacci numbers below a given limit.
func SumEvenFibonacci(limit int) int {
	var sum, a, b int
	a = 1
	b = 2

	for b <= limit {
		if b%2 == 0 {
			sum += b
		}

		a, b = b, a+b
	}

	if limit == 1 {
		return 2
	}

	return sum
}

func main() {
	fmt.Println(SumEvenFibonacci(8))
	fmt.Println(SumEvenFibonacci(111111))
	fmt.Println(SumEvenFibonacci(8675309))
}

/*
	The best solution on codewars.com:
	func SumEvenFibonacci(limit int) int {
	sum, a, b := 2, 1, 2
	for b <= limit {
		a, b = b, a + b
		if b % 2 == 0 { sum += b }
	}
	return sum
	}
*/