package main

import (
	"fmt"
)

// The function calculates the weight of potatoes after a change in moisture content.
func Potatos(p0, w0, p1 int) int {
	return w0 * (100 - p0) / (100 - p1)
}

func main() {
	fmt.Println(Potatos(99, 100, 98))
	fmt.Println(Potatos(82, 127, 80))
	fmt.Println(Potatos(93, 129, 91))
}
