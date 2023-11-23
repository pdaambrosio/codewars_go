package main

import (
	"fmt"
)

// The function "GetPlanetName" takes an ID as input and returns the corresponding planet name.
func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return ""
	}
}

func main() {
	fmt.Println(GetPlanetName(2))
}
