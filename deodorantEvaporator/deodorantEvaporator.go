package main

import "fmt"

// The Evaporator function calculates the number of days it takes for a liquid to evaporate below a
// certain threshold based on the initial content and evaporation rate per day.
func Evaporator(content float64, evapPerDay int, threshold int) int {
  var days int
  var thresholdContent float64 = content * float64(threshold) / 100
  for content > thresholdContent {
	content -= content * float64(evapPerDay) / 100
	days++
  }

  return days
}

func main() {
  fmt.Println(Evaporator(10, 10, 10))
  fmt.Println(Evaporator(10, 10, 5))
  fmt.Println(Evaporator(100, 5, 5))
}