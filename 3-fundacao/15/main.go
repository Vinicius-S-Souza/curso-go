package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showtype(x)
	showtype(y)
}

func showtype(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}