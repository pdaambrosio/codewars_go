package main

import (
	"fmt"
)

// The NbMonths function calculates the number of months needed to save enough money to buy a new item
// while accounting for depreciation and monthly savings.
func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	months := 0
	savings := 0.0
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)
	loss := percentLossByMonth / 100.0

	for priceOld + float64(savings) < priceNew {
		months++
		if months % 2 == 0 {
			loss += 0.005
		}

		priceOld -= priceOld * loss
		priceNew -= priceNew * loss
		savings += float64(savingperMonth)
	}

	return [2]int{months, int(priceOld + float64(savings) - priceNew + 0.5)}
}

func main() {
	fmt.Println(NbMonths(2000, 8000, 1000, 1.5))
	fmt.Println(NbMonths(12000, 8000, 1000, 1.5))
	fmt.Println(NbMonths(8000, 12000, 500, 1.0))
}

// The best solution in codewars:
// package kata

// func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
//   months := 0
//   priceOld := float64(startPriceOld)
//   priceNew := float64(startPriceNew)

//   for ; priceOld + float64(months * savingperMonth) < priceNew; months++ {
//       if months % 2 == 1 {   
//         percentLossByMonth += 0.5
//       }
//       priceOld -= priceOld * percentLossByMonth / 100.0
//       priceNew -= priceNew * percentLossByMonth / 100.0
//   }

//   return [2]int{months, int(priceOld + float64(months * savingperMonth) - priceNew + 0.5)}
// }