package main

import "fmt"

// The MultiplicationTable function generates a multiplication table of a given size.
func MultiplicationTable(size int) [][]int {
  var result [][]int
  for i := 1; i <= size; i++ {
	var row []int
	for j := 1; j <= size; j++ {
	  row = append(row, i * j)
	}
	result = append(result, row)
  }

  return result
}

func main() {
	fmt.Println(MultiplicationTable(1))
	fmt.Println(MultiplicationTable(2))
	fmt.Println(MultiplicationTable(3))
}
