package main

import (
	"fmt"
)

// The `DNAStrand` function takes a DNA string as input and returns the complementary DNA strand by
// replacing each nucleotide with its complementary base (A with T, T with A, C with G, and G with C).
func DNAStrand(dna string) string {
	var result string

	for _, char := range dna {
		switch char {
		case 'A':
			result += "T"
		case 'T':
			result += "A"
		case 'C':
			result += "G"
		case 'G':
			result += "C"
		}
	}

	return result
}

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}

// The best solution from codewars:
// package kata

// import "strings"

// var dnaReplacer *strings.Replacer = strings.NewReplacer(
//   "A", "T",
//   "T", "A",
//   "C", "G",
//   "G", "C",
// )

// func DNAStrand(dna string) string {
//   return dnaReplacer.Replace(dna)
// }