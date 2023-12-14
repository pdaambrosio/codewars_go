package main

import (
	"fmt"
	"strconv"
)

// The function takes a string as input and returns the sum of all the integers present in the string.
func SumOfIntergersInString(strInput string) int {
	var sum int
	var strNum string

	for _, value := range strInput {
		if value >= '0' && value <= '9' {
			strNum += string(value)
		} else {
			if strNum != "" {
				num, _ := strconv.Atoi(strNum)
				sum += num
				strNum = ""
			}
		}
	}

	if strNum != "" {
		num, _ := strconv.Atoi(strNum)
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(SumOfIntergersInString("12.4"))
	fmt.Println(SumOfIntergersInString("h3ll0w0rld"))
	fmt.Println(SumOfIntergersInString("2 + 3 = "))
	fmt.Println(SumOfIntergersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
}

// best solution on codewars using regex
/*
package kata

import ("regexp"; "strconv")

func SumOfIntegersInString(strng string) int {
  r := 0
  for _, s := range regexp.MustCompile(`\d+`).FindAllString(strng, -1) {
    x, _ := strconv.Atoi(s)
    r += x
  }
  return r
}
*/