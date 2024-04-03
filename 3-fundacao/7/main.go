package main

import (
	"fmt"
)

func main() {

	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Wesley"] = 1000

	for nome, salarios := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salarios)
	}

	for _, salarios := range salarios {
		fmt.Printf("O salario é %d\n", salarios)
	}

}
