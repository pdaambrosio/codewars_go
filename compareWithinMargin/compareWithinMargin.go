package main

import "fmt"

func CloseCompare(a, b, margin float64) int {
	switch {
	case a < b - margin:
		return -1
	case a > b + margin:
		return 1
	default:
		return 0
	}
}

func main() {
	fmt.Println(CloseCompare(4.0, 5.0, 0.0,))
	fmt.Println(CloseCompare(5.0, 5.0, 0.0))
	fmt.Println(CloseCompare(6.0, 5.0, 0.0))
}