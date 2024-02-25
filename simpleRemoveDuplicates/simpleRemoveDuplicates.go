package main

import (
	"fmt"
)


// The `removeDuplicates` function takes an array of integers and returns a new array with duplicate
// elements removed while maintaining the original order.
func removeDuplicates(arr []int) []int {
    validDup := make(map[int]bool)
    result := make([]int, 0)
    for i := len(arr) - 1; i >= 0; i-- {
        fmt.Println(arr[i], validDup[arr[i]])
        if !validDup[arr[i]] {
            validDup[arr[i]] = true
            result = append([]int{arr[i]}, result...)
        }
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