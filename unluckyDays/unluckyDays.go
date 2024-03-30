package main

import (
	"fmt"
	"time"
)

// The function `UnluckDays` calculates the number of Fridays that fall on the 13th day of each month
// in a given year.
func UnluckDays(year int) int {
	var unluckCount int = 0
	for month := 1; month <= 12; month++ {
		if time.Date(year, time.Month(month), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5 {
			unluckCount++
		}
	}

	return unluckCount
}

func main() {
	fmt.Println(UnluckDays(2015))
	fmt.Println(UnluckDays(1986))
	fmt.Println(UnluckDays(2019))
}
