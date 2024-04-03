package main

import (
	"fmt"
	"curso-go/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Retultado: ", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}

func Soma(i1, i2 int) {
	panic("unimplemented")
}
