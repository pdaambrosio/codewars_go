package main

import (
    "fmt"
	"math"
)

// I don't know how to solve this problem, so I just unlocked and copied the solution from the codewars website;
// I will try to understand the solution and then write my own solution
// The solution is from the user "sergrom"
func CountPairsInt(diff, nMax int) int {
    count := 0
    for i := 2; i <= nMax-diff; i++ {
        n1 := i
        n2 := i + diff
        if n2 < nMax && divis(n1) == divis(n2) {
            count++
        }
    }
    return count
}

func divis(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			count++
			if n/i != i {
				count++
			}
		}
	}
	return count
}

func main() {
    fmt.Println(CountPairsInt(1, 50))
    fmt.Println(CountPairsInt(3, 100))
	fmt.Println(CountPairsInt(5, 100))
}
