package main


// The function checks if a given number is a Wilson prime.
func AmIWilson(n int) bool {
	return n == 5 || n == 13 || n == 563
}

func main() {
	println(AmIWilson(5))
	println(AmIWilson(13))
	println(AmIWilson(563))
	println(AmIWilson(12))
}
