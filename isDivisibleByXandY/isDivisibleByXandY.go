package main

import (
	"fmt"
)

// The IsDivisible function checks if a number n is divisible by both x and y.
func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func main() {
	fmt.Println(IsDivisible(3, 3, 4))
	fmt.Println(IsDivisible(12, 3, 4))
	fmt.Println(IsDivisible(8, 3, 4))
	fmt.Println(IsDivisible(48, 3, 4))
	fmt.Println(IsDivisible(100, 5, 10))
	fmt.Println(IsDivisible(100, 5, 3))
}
