package main

import (
	"fmt"
)

func Between(a, b int) []int {
	var result []int

	for i := a; i <= b; i++ {
		result = append(result, i)
	}

	return result
}

func main() {
	fmt.Println(Between(1, 4))
	fmt.Println(Between(-2, 2))
	fmt.Println(Between(-10, 10))
}
