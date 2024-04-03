package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 5, 45, 50))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
