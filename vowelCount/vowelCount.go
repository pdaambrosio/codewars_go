package main

import (
	"fmt"
)

// The function "GetCount" counts the number of vowels in a given string.
func GetCount(str string) (count int) {
	for _, vowels := range str {
		switch vowels {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	return
}

func main() {
	fmt.Println(GetCount("abracadabra"))
	fmt.Println(GetCount("o a kak ushakov lil vo kashu kakao"))
	fmt.Println(GetCount("Akatsuki no Yona"))
}
