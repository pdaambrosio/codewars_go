package main

import (
	"fmt"
)

// The function calculates the persistence of a number by multiplying its digits until it becomes a
// single digit number, and returns the number of times the multiplication was performed.
func Persistence(n int) int {
	var count int
	for n >= 10 {
		var sum int = 1
		for n > 0 {
			sum *= n % 10
			n /= 10
		}

		n = sum
		count++
	}

	return count
}

func main() {
	fmt.Println(Persistence(39))
	fmt.Println(Persistence(999))
	fmt.Println(Persistence(4))
}
