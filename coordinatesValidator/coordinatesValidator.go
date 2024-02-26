package main

import (
	"fmt"
	"regexp"
)

// The isValidCoordinates function in Go uses a regular expression to validate latitude and longitude
// coordinates.
func isValidCoordinates(coordinates string) bool {
	regex := regexp.MustCompile(`^-?(90|[0-8]?\d(\.\d+)?), -?(180|[01]?[0-7]?\d(\.\d+)?)$`)
	return regex.MatchString(coordinates)
}

func main() {
	fmt.Println(isValidCoordinates("-23, 25"))
	fmt.Println(isValidCoordinates("4, -3"))
	fmt.Println(isValidCoordinates("23.234, - 23.4234"))
}