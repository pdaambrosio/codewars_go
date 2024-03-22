package main

import (
	"fmt"
)

// The GetSize function calculates and returns the surface area and volume of a rectangular box based
// on its width, height, and depth.
func GetSize(w, h, d int) [2]int {
	return [2]int{2 * (w*h + h*d + w*d), w * h * d}
}

func main() {
	fmt.Println(GetSize(4, 2, 6)) // [88, 48]
	fmt.Println(GetSize(1, 1, 1)) // [6, 1]
	fmt.Println(GetSize(1, 2, 1)) // [10, 2]
}