package main

import "fmt"

// The function "FindMultiples" takes an integer and a limit as input and returns a slice of all the
// multiples of the integer up to the limit.
func FindMultiples(integer, limit int) []int {
	var result []int
	for i := 1; i <= limit; i++ {
		if i % integer == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	fmt.Println(FindMultiples(5, 25))
	fmt.Println(FindMultiples(1, 2))
	fmt.Println(FindMultiples(5, 7))
}

/*
The best solution in codewars:
	package kata

	func FindMultiples(integer, limit int) (res[]int) {
	for i := integer; i <= limit; i += integer { res = append(res, i)}
	return
	}
*/