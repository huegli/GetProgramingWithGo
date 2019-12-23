package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	balance := 0
	dollars := 0
	cents := 0
	for balance < 2000 {
		switch rand.Intn(3) {
		case 0:
			balance += 5
		case 1:
			balance += 10
		case 2:
			balance += 25
		}
		dollars = balance / 100
		cents = balance % 100
		fmt.Printf("Balance %v.%v\n", dollars, cents)
	}
}
