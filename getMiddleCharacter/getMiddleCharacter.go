package main

import (
	"fmt"
)

// The function "GetMiddle" returns the middle character(s) of a given string.
func GetMiddle(s string) string {
	if len(s) % 2 == 0 {
		return s[len(s) / 2 - 1 : len(s) / 2 + 1]
	}
	
	return s[len(s) / 2 : len(s) / 2 + 1]
}

func main() {
	fmt.Println(GetMiddle("test"))
	fmt.Println(GetMiddle("testing"))
	fmt.Println(GetMiddle("middle"))
}
