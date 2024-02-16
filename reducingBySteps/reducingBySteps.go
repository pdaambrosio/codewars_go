package main

import "fmt"

// The function calculates the greatest common divisor (GCD) of two integers using the Euclidean
// algorithm.
func Gcdi(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	if x < 0 {
		return -x
	}
	return x
}

// The function "Som" in Go takes two integers as input and returns their sum.
func Som(x, y int) int {
	return x + y
}

// The Maxi function returns the maximum value between two integers.
func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// The Mini function returns the smaller of two given integers.
func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// The Lcmu function calculates the least common multiple of two integers.
func Lcmu(x, y int) int {
	result := x * y / Gcdi(x, y)
	if result < 0 {
		return -result
	}
	return result
}

// The `OperArray` function takes a function `f`, an array `arr`, and an initial value `init`, and
// applies `f` to each element of `arr` in a cumulative manner, updating `init` with the result of each
// operation.
type FParam func(int, int) int
func OperArray (f FParam, arr []int, init int) []int {
	for i := range arr {
		arr[i] = f(init, arr[i])
		init = arr[i]
	}
	return arr
}

func main() {
	dta := []int{6, -72, -62, -22, -23, 80}
	fmt.Println(OperArray(Som, []int{18, 69, -90, -78, 65, 40}, 0))
	fmt.Println(OperArray(Maxi, []int{18, 69, -90, -78, 65, 40}, 1))
	fmt.Println(OperArray(Mini, []int{64, -67, -43, 12, -15, 108, 12, 104, -36}, 0))
	fmt.Println(OperArray(Lcmu, dta, dta[0]))
}