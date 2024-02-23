package main

import "fmt"

// The Solve function takes a string as input and counts the number of uppercase letters, lowercase
// letters, numbers, and special characters in the string, returning the counts as an array of
// integers.
func Solve(s string) []int {
	var upper, lower, number, special int
	for _, char := range s {
		switch {
		case char >= 'A' && char <= 'Z':
			upper++
		case char >= 'a' && char <= 'z':
			lower++
		case char >= '0' && char <= '9':
			number++
		default:
			special++
		}
	}

	return []int{upper, lower, number, special}
}

func main() {
	fmt.Println(Solve("Codewars@codewars123.com"))
	fmt.Println(Solve("bgA5<1d-tOwUZTS8yQ"))
	fmt.Println(Solve("P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H"))
}

// The best solution on codewars is:
// package kata
// func Solve(s string) []int {
//   r := make([]int, 4)
//   for _, c := range s {
//     switch {
//       case c >= 'A' && c <= 'Z': r[0]++;
//       case c >= 'a' && c <= 'z': r[1]++;
//       case c >= '0' && c <= '9': r[2]++;
//       default: r[3]++;
//     }
//   }
//   return r
// }