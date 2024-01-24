package main

import (
	"fmt"
)

// The CountSheeps function counts the number of true values in an array of boolean values.
func CountSheeps(arrayOfSheep []bool) int {
	count := 0
	for _, v := range arrayOfSheep {
		if v {
			count++
		}
	}

	return count
}

func main() {
	arrayOfSheep := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	fmt.Println(CountSheeps(arrayOfSheep))
	arrayOfSheep1 := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, false, true,
		false, false, true, true}
	fmt.Println(CountSheeps(arrayOfSheep1))
	arrayOfSheep2 := []bool{true, true, true, false,
		true, true, true, true,
		true, false, false, false,
		true, false, false, true,
		true, false, true, true,
		false, false, true, true}
	fmt.Println(CountSheeps(arrayOfSheep2))
}

/*
The best solution on codewars:
	package kata

	func CountSheeps(numbers []bool) (count int) {
		for _, v := range numbers {
			if v { count++ }
		}
		return
	}
*/
