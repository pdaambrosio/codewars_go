package main

import (
	"fmt"
)

// The RepeatStr function takes an integer and a string as input, and returns a new string that repeats
// the input string a specified number of times.
func RepeatStr(repititions int, value string) (result string) {
	for i := 0; i < repititions; i++ {
		result += value
	}

	return
}

func main() {
	fmt.Println(RepeatStr(6, "I")) // IIIIII
	fmt.Println(RepeatStr(5, "Hello")) // HelloHelloHelloHelloHello
	fmt.Println(RepeatStr(3, "World")) // WorldWorldWorld
}


// The best solution from codewars:
// package kata

// import "strings"

// func RepeatStr(repititions int, value string) string {
//   return strings.Repeat(value, repititions)
// }