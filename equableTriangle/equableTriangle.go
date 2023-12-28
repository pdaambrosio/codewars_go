package main

import (
	"fmt"
	"math"
)

// The function EquableTriangle checks if a triangle with side lengths a, b, and c is an equable
// triangle.
func EquableTriangle(a, b, c int) bool {
	p := (a + b + c) / 2
	return math.Sqrt(float64(p*(p-a)*(p-b)*(p-c))) == float64(a+b+c)
}

func main() {
	fmt.Println(EquableTriangle(5, 12, 13))
	fmt.Println(EquableTriangle(2, 3, 4))
	fmt.Println(EquableTriangle(6, 8, 10))
}

// The best practice solution from the codewars website:
/*
func EquableTriangle(a, b, c int) bool {
    p := a+b+c
    return 16*p == (p-2*a)*(p-2*b)*(p-2*c)
}
*/