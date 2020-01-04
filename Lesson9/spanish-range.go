package main

import (
	"fmt"
)

func main() {
	question := "¿Cómo estás?"

	for _, c := range question {
		fmt.Printf("%c ", c)
	}
}
