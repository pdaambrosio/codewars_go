package main

import (
	"fmt"
	"strings"
)

// The function FizzBuzzCuckooClock takes a time string in the format "HH:MM" and returns a specific
// output based on the hour and minute values.
func FizzBuzzCuckooClock(time string) string {
	hour, minute := 0, 0
	fmt.Sscanf(time, "%d:%d", &hour, &minute)
	if minute == 0 {
		if hour == 0 {
			hour = 12
		} else if hour > 12 {
			hour -= 12
		}
		return "Cuckoo" + strings.Repeat(" Cuckoo", hour - 1)
	} else if minute == 30 {
		return "Cuckoo"
	}

	if minute % 3 == 0 && minute % 5 == 0 {
		return "Fizz Buzz"
	}

	if minute % 3 == 0 {
		return "Fizz"
	}

	if minute % 5 == 0 {
		return "Buzz"
	}

	return "tick"
}

func main() {
	fmt.Println(FizzBuzzCuckooClock("13:34"))
	fmt.Println(FizzBuzzCuckooClock("21:00"))
	fmt.Println(FizzBuzzCuckooClock("11:15"))
	fmt.Println(FizzBuzzCuckooClock("03:03"))
	fmt.Println(FizzBuzzCuckooClock("14:30"))
	fmt.Println(FizzBuzzCuckooClock("08:55"))
	fmt.Println(FizzBuzzCuckooClock("00:00"))
	fmt.Println(FizzBuzzCuckooClock("12:00"))
}

// The best solution in codewars is:
// package kata
// import (
//   "strings"
//   "strconv"
// )
// func FizzBuzzCuckooClock(time string) string {
// 	h, _ := strconv.Atoi(time[0:2])
//   m, _ := strconv.Atoi(time[3:5])
//   switch {
//     case m==0:
//       return strings.Repeat("Cuckoo ", (h+11)%12) + "Cuckoo"
//     case m==30:
//       return "Cuckoo"
//     case m%15 == 0:
//       return "Fizz Buzz"
//     case m%5 == 0:
//       return "Buzz"
//     case m%3 == 0:
//       return "Fizz"
//     default:
//       return "tick"
//   }
// }
