package main

import (
	"fmt"
)

// The DigitalRoot function calculates the digital root of a given number by repeatedly summing its
// digits until a single digit is obtained.
func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	var sum int = 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return DigitalRoot(sum)
}

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(195))
	fmt.Println(DigitalRoot(992))
}

/*
The best solutions in codewars:
	package kata
	func DigitalRoot(n int) int {
		return (n - 1) % 9 + 1
	}
and
	package kata
	func DigitalRoot(n int) int {
		for n > 9 {
			n = (n / 10) + n%10
		}
		return n
	}
*/