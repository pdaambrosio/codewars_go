package main

import "fmt"

// The `DoubleInteger` function in Go doubles the input integer value.
func DoubleInteger(n int) int {
	return n * 2
}

func main() {
	fmt.Println(DoubleInteger(2))
	fmt.Println(DoubleInteger(4))
	fmt.Println(DoubleInteger(8))
}