package main

import (
	"fmt"
)

// The function calculates the sum of all numbers below a given number that are divisible by 3 or 5.
func Multiple3And5(number int) int {
	var sum int
	for i := 1; i < number; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	
	return sum
}

func main() {
	fmt.Println(Multiple3And5(10))
	fmt.Println(Multiple3And5(20))
	fmt.Println(Multiple3And5(200))
}
