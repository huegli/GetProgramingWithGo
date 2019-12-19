package main

import "fmt"

func main() {
	const distance, traveltime = 58000000, 28

	speed := distance / (traveltime * 24)

	fmt.Printf("You need to go %v km/h to reach Malacandra in %v days\n", speed, traveltime)
}
