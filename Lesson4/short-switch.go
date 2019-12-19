package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    	rand.Seed(time.Now().UnixNano())
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galatic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}
