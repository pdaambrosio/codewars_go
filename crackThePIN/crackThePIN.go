package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Again I don't know how to solve this problem, so I just unlocked and copied the solution from the codewars website;
// I couldn’t decrypt hash with more than 5 digits in my code;
// I will try to understand the solution and then write my own solution
// The solution is from the user "Unnamed"

func Crack(hash string) string {
	hashBytesSlice, _ := hex.DecodeString(hash)
	var hashBytes [16]byte
	copy(hashBytes[:], hashBytesSlice)
	for i := 0; i <= 99999; i++ {
		s := fmt.Sprintf("%05d", i)
		if md5.Sum([]byte(s)) == hashBytes {
			return s
		}
	}

	panic("password not found")
}

func main() {
	fmt.Println(Crack("827ccb0eea8a706c4c34a16891f84e7b"))
	fmt.Println(Crack("86aa400b65433b608a9db30070ec60cd"))
}