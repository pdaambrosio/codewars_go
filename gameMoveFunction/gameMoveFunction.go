package main

import (
	"fmt"
)

// The Move function takes a position and a roll, and returns the new position after moving twice the
// value of the roll.
func Move(position int, roll int) int {
	return position + roll*2
}

func main() {
	fmt.Println(Move(0, 4))
	fmt.Println(Move(3, 6))
	fmt.Println(Move(2, 5))
}
