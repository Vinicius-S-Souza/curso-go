package main

func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "Word"

	println(<-ch)
	println(<-ch)
}