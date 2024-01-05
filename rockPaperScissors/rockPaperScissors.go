package main

import (
	"fmt"
)

// The function "rockPaperScissors" determines the winner of a rock-paper-scissors game based on the
// choices of two players.
func rockPaperScissors(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	switch p1 {
	case "rock":
		if p2 == "scissors" {
			return "Player 1 won!"
		}
	case "scissors":
		if p2 == "paper" {
			return "Player 1 won!"
		}
	case "paper":
		if p2 == "rock" {
			return "Player 1 won!"
		}
	}

	return "Player 2 won!"
}

func main() {
	fmt.Println(rockPaperScissors("rock", "scissors"))
	fmt.Println(rockPaperScissors("rock", "paper"))
	fmt.Println(rockPaperScissors("paper", "paper"))
}

// Best solution from codewars
/*
	package kata

	var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

	func Rps(a, b string) string {
		if a == b {
			return "Draw!"
		}
		if m[a] == b {
			return "Player 2 won!"
		}
		return "Player 1 won!"
	}
*/