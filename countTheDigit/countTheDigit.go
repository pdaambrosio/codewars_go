package main

import (
	"fmt"
)

// The function `NbDig` counts the number of times a digit `d` appears in the squares of numbers from 0
// to `n`.
func NbDig(n int, d int) int {
	var count int

	for i := 0; i <= n; i++ {
		for _, v := range fmt.Sprintf("%d", i*i) {
			if int(v-'0') == d {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
	fmt.Println(NbDig(12224, 8))
}

/*
the best solution in codewars:

	package kata

	import (
	"strings"
	"strconv"
	)

	func NbDig(n int, d int) (out int) {
	for i := 0; i <= n; i++ {
		out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return
	}
*/