package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 5, 45, 50) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
