package main

import (
	"fmt"
	"strconv"
)

// The function takes a binary string as input and converts it to a decimal integer.
func BinToDec(bin string) int {
	if bin, err := strconv.ParseInt(bin, 2, 64); err == nil {
		return int(bin)
	}

	return 0
}

func main() {
	fmt.Println(BinToDec("10"))
	fmt.Println(BinToDec("11"))
	fmt.Println(BinToDec("101010"))
	fmt.Println(BinToDec("1001001"))
}
