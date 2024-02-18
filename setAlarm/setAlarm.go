package main

import "fmt"

// The setAlarm function returns true if the person is employed and not on vacation.
func setAlarm(employed bool, vacation bool) bool {
	return employed && !vacation
}

func main() {
	fmt.Println(setAlarm(true, true))
	fmt.Println(setAlarm(false, true))
	fmt.Println(setAlarm(false, false))
	fmt.Println(setAlarm(true, false))
}