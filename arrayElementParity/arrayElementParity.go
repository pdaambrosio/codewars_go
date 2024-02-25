package main

import (
	"fmt"
)

// The Solve function takes an array of integers and returns the unique integer that appears only once
// in the array.
func Solve(arr []int) int {
    uniqueIntegers := make(map[int]bool)
    for _, num := range arr {
        if _, exists := uniqueIntegers[-num]; exists {
            delete(uniqueIntegers, -num)
        } else {
            uniqueIntegers[num] = true
        }
    }

    for num := range uniqueIntegers {
        return num
    }

    return 0
}

func main() {
	fmt.Println(Solve([]int{1, -1, 2, -2, 3}))
	fmt.Println(Solve([]int{-3, 1, 2, 3, -1, -4, -2}))
	fmt.Println(Solve([]int{1, -1, 2, -2, 3, 3}))
}

// The best solution in my opinion on codewars is:
// package kata 

// func Solve(arr []int) int {
// 	hash := make(map[int]bool)
// 	ans := 0
// 	for _, entry := range arr {
// 		if _, value := hash[entry]; !value {
// 			hash[entry] = true
// 			ans += entry
// 		}
// 	}
// 	return ans
// }