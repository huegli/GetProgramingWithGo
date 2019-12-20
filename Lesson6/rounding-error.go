package main

import (
	"fmt"
)

func main() {
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "Deg F\n")
	fmt.Print((9.0/5.0*celsius)+32, "Deg F\n")
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "Deg F\n")
}
