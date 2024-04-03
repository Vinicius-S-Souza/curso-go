package main

import "fmt"

type Conta struct {
	saldo int
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	fmt.Println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo:100}
	conta.simular(200)
	fmt.Println(conta.saldo)
}
