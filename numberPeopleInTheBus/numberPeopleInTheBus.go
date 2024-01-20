package main

import "fmt"

// The function calculates the total number of people on a bus after a series of stops.
func Number(busStops [][2]int) int {
	var total int = 0
	for _, stop := range busStops {
		total += stop[0] - stop[1]
	}

	return total
}

func main() {
	fmt.Println(Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}))
}
