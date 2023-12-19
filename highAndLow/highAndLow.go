package main

import (
	"fmt"
	"strconv"
	"strings"
)

// The function "HighAndLow" takes a string of numbers separated by spaces, finds the highest and
// lowest numbers, and returns them as a string.
func HighAndLow(in string) string {
	nums := strings.Split(in, " ")
	max, _ := strconv.Atoi(nums[0])
	min, _ := strconv.Atoi(nums[0])

	for _, v := range nums {
		num, _ := strconv.Atoi(v)

		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	return strconv.Itoa(max) + " " + strconv.Itoa(min)
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}
