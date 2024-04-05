package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is number\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()		
	}
}

// Thread 1 (Principal)
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	// Thread 2
	go task("A", &waitGroup)
	// Thread 3  
	go task("B", &waitGroup)
	
	// Thread 4 - Função Anônima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is number\n", i, "Anônima")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	// Sair - Fim Thread 1 (Principal)
	waitGroup.Wait()
}