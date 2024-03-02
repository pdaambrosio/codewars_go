package main

import "fmt"

// The function calculates the Fibonacci number at a given position using recursion.
func Fib(n int) int {
	if n == 0 {
	  return 0
	} 
	
	if n == 1 {
	  return 1
	} 
	
	return Fib(n-1) + Fib(n-2)
  }

func main() {
	fmt.Println(Fib(3))
	fmt.Println(Fib(4))
	fmt.Println(Fib(5))
}
