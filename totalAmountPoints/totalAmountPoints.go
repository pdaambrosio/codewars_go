package main

import (
	"fmt"
	"strings"
)

// The function calculates the total points earned in a series of games based on the scores.
func Points(games []string) int {
	var total int

	for _, game := range games {
		splitGame := strings.Split(game, ":")
		x, y := splitGame[0], splitGame[1]

		switch {
		case x > y:
			total += 3
		case x == y:
			total += 1
		}

		// or
		// if game[0] > game[2] {
		// 	total += 3
		// } else if game[0] == game[2] {
		// 	total += 1
		// }
	}

	return total
}

func main() {
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))
	fmt.Println(Points([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}))
	fmt.Println(Points([]string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}))
}