package main

import "fmt"

func EvenOrOdd(number int) string {
  if (number % 2 == 0) {
	return "Even"
  }

  return "Odd"
}


func main() {
  fmt.Println(EvenOrOdd(2))
  fmt.Println(EvenOrOdd(0))
  fmt.Println(EvenOrOdd(7))
  fmt.Println(EvenOrOdd(1))
}
