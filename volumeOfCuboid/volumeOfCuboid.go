package main

import (
	"fmt"
)

// The GetVolumeOfCuboid function calculates the volume of a cuboid based on its length, width, and
// height.
func GetVolumeOfCuboid(length float64, width float64, height float64) float64 {
	return length * width * height
}

func main() {
	fmt.Println(GetVolumeOfCuboid(1.0, 2.0, 2.0))
	fmt.Println(GetVolumeOfCuboid(6.3, 2.5, 5.0))
	fmt.Println(GetVolumeOfCuboid(1.0, 2.0, 2.0))
}
