package main

import (
	"fmt"
)

// The `Combat` function in Go calculates the remaining health after taking damage.
func Combat(health, damage float64) float64 {
	if health - damage < 0 {
		return 0
	}

	return health - damage
}

func main() {
	fmt.Println(Combat(100, 5))
	fmt.Println(Combat(100.0, 50.0))
	fmt.Println(Combat(1.0, 100))
}