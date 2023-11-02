package main

import "fmt"

// Function hero takes the number of bullets and the number of dragons and returns true if hero can kill all the dragons, false otherwise.
func hero(bullets, dragons int) bool {
	return bullets/2 >= dragons
}

func main() {
	fmt.Println(hero(10, 5))
	fmt.Println(hero(7, 4))
	fmt.Println(hero(4, 5))
	fmt.Println(hero(100, 40))
	fmt.Println(hero(1500, 751))
	fmt.Println(hero(0, 1))
}