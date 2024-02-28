package main

import (
	"fmt"
)

// The function Litres calculates the amount of water consumed in liters based on the time spent
// drinking.
func Litres(time float64) int {
	return int(time * 0.5)
}

func main() {
	fmt.Println(Litres(2))
	fmt.Println(Litres(1.4))
	fmt.Println(Litres(12.3))
}