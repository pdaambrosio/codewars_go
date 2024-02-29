package main

import (
	"fmt"
)

// Gps function returns the highest average speed of the segments.
// s is an integer representing the time in seconds between each segment.
// x is a slice of floats representing the distance in meters of each segment.
func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}

	max := 0.0
	for i := 0; i < len(x) - 1; i++ {
		speed := (x[i + 1] - x[i]) * 3600 / float64(s)
		if speed > max {
			max = speed
		}
	}

	return int(max)
}

func main() {
	fmt.Println(Gps(20, []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}))
	fmt.Println(Gps(12, []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}))
	fmt.Println(Gps(20, []float64{}))
}

// The best solution on codewars:
// package kata

// import "math"

// func Gps(s int, segments []float64) int {
// 	var maxSpeed float64
// 	for index := 1; index < len(segments); index++ {
// 		startSegment, endSegment := segments[index-1], segments[index]
// 		kmPerHour := 3600.0 * (endSegment - startSegment) / float64(s)
// 		maxSpeed = math.Max(maxSpeed, kmPerHour)
// 	}
// 	return int(maxSpeed)
// }