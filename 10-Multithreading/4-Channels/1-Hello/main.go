package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Canal Vazio

	// Thread 2
	go func() {
		canal <- "Olá Mundo!" // Canal Cheio
	}()

	msg := <-canal // Esvazia Canal
	fmt.Println(msg)
}
