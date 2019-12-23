package main

import (
	"fmt"
	"math/big"
)

func main() {
	distanceToCanis := new(big.Int)
	distanceToCanis.SetString("236000000000000000", 10)

	lightInAYear := big.NewInt(299792 * 86400 * 365)

	lightYears := new(big.Int)
	lightYears.Div(distanceToCanis, lightInAYear)

	fmt.Println("It takes", lightYears, "light years to travel to Canis.")
}
