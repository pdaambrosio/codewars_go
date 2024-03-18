package main

import (
	"fmt"
)

// The FindNextSquare function in Go finds the next perfect square after a given number.
func FindNextSquare(sq int64) int64 {
	if sq < 0 {
		return -1
	}

	x := int64(0)
	for x * x < sq {
		x++
	}

	if x * x == sq {
		return (x + 1) * (x + 1)
	}

	return -1
}

func main() {
	fmt.Println(FindNextSquare(121))
	fmt.Println(FindNextSquare(625))
	fmt.Println(FindNextSquare(155))
}

// The best solution in code wars is:
// package kata

// import (
//   "math"
// )

// func FindNextSquare(sq int64) int64 {
//   res := math.Pow(math.Sqrt(float64(sq)) + 1, 2)
  
//   if res == math.Trunc(res) {
//     return int64(res)
//   }
//   return -1
// }