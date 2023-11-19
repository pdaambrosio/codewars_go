package main

import (
	"fmt"
)

func Past(h, m, s int) int {
	return (h*3600 + m*60 + s) * 1000
}

func main() {
	fmt.Println(Past(0, 1, 1))
	fmt.Println(Past(1, 1, 1))
	fmt.Println(Past(0, 0, 0))
	fmt.Println(Past(1, 0, 1))
	fmt.Println(Past(1, 0, 0))
}
