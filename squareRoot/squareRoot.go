package main

import (
	"fmt"
	"math"
)

// Function SquareOrSquareRoot returns a new array (list) of the squares of the numbers in the input array (list) 
// if the square root is an integer; otherwise, returns the square root of the number.
func SquareOrSquareRoot(arr []int) []int {
	var result []int
	for _, num := range arr {
		if math.Sqrt(float64(num)) == math.Floor(math.Sqrt(float64(num))) {
			result = append(result, int(math.Sqrt(float64(num))))
		} else {
			result = append(result, num*num)
		}
	}
	return result
}

func main() {
	fmt.Println(SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
	fmt.Println(SquareOrSquareRoot([]int{10, 20, 30, 40, 50}))
	fmt.Println(SquareOrSquareRoot([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
