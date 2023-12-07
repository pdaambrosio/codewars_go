package main

import (
	"fmt"
)

func NumberToString(num int) string {
	return fmt.Sprintf("%d", num)
}

func main() {
	fmt.Println(NumberToString(67))
	fmt.Println(NumberToString(123))
	fmt.Println(NumberToString(999))
}