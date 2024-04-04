package main

import "fmt"

// The Opposite function in Go returns the negative value of the input integer.
func Opposite(value int) int {
	return -value
}

func main() {
	fmt.Println(Opposite(1))
	fmt.Println(Opposite(-1))
	fmt.Println(Opposite(0))
}
