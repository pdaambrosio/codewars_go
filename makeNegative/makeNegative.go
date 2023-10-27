package main

// The function takes an integer as input and returns its negative value if it is positive, otherwise
// it returns the input itself.

func makeNegative(num int) int {
	if num > 0 {
		return num * -1
	}
	return num
}

func main() {
	println(makeNegative(1))
	println(makeNegative(-5))
	println(makeNegative(0))
}
