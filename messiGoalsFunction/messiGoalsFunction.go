package main

import "fmt"

// The function "Goals" calculates the total number of goals scored in three different competitions.
func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}

func main() {
	fmt.Println(Goals(1, 2, 3))
	fmt.Println(Goals(5, 10, 2))
	fmt.Println(Goals(0, 0, 0))
}

/*
The best solution in codewars:
	package kata

	func Goals(goals ...int) (res int) {
	for _, i := range goals {
			res += i
		}
	return
	}
*/