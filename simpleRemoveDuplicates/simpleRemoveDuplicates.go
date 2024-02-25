package main

import (
	"fmt"
)

// The `removeDuplicates` function in Go removes duplicate elements from an integer slice while
// preserving the order of the original elements.
func removeDuplicates(arr []int) []int {
    lastIndex := make(map[int]int)
    result := make([]int, 0)

    for i := len(arr) - 1; i >= 0; i-- {
        if _, ok := lastIndex[arr[i]]; !ok {
            result = append(result, arr[i])
        }
        lastIndex[arr[i]] = i
    }

    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }

    return result
}

func main() {
	fmt.Println(removeDuplicates([]int{3,4,4,3,6,3}))
	fmt.Println(removeDuplicates([]int{1,2,1,2,1,2,3}))
	fmt.Println(removeDuplicates([]int{1,2,3,4}))
}

// The best solution in my opinion on codewars is:
// package kata 

// func Solve(arr []int) (res []int) {
//     visited := map[int]bool{}
//     for i := len(arr) - 1; i >= 0; i-- {
//         n := arr[i]
//         if visited[n] { continue }
      
//         visited[n] = true
//         res = append([]int{n}, res...)
//     }
  
//     return
// }