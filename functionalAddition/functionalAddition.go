package main

import (
	"fmt"
)

// The Add function returns a closure that adds a given number to its argument.
func Add(n int) func(int) int {
	return func(m int) int {
		return n + m
	}
}

func main() {
	fmt.Println(Add(1)(3))
	fmt.Println(Add(2)(3))
	fmt.Println(Add(3)(3))
}