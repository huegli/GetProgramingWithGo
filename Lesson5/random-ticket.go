package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Spaceline         Days Trip  type  Price")
	fmt.Println("========================================")

	for ticket := 0; ticket < 10; ticket++ {

		var spaceline = "Space Adventures"
		switch rand.Intn(3) {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "SpaceX"
		}

		factor := rand.Intn(15)
		speed := 16 + factor
		days := 62100000 / (speed * 60 * 60 * 24)
		price := 36 + factor

		tripType := "One-Way"
		if rand.Intn(2) == 1 {
			tripType = "Round-trip"
			price = price * 2
		}

		fmt.Printf("%-17s %4d %-10s  $%4d\n", spaceline, days, tripType, price)
	}

}
