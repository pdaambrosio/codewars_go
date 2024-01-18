package main

import "fmt"

// The function "RoundToNext5" takes an integer as input and returns the next multiple of 5.
func RoundToNext5(n int) int {
	for n % 5 != 0 {
		n++
	}

	return n
}

func main() {
	fmt.Println(RoundToNext5(0))
	fmt.Println(RoundToNext5(21))
	fmt.Println(RoundToNext5(-5))
}
