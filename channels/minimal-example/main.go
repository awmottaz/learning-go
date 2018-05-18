// Demonstrate a minimal way to work with channels and goroutines.
package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("Exiting main...")
}

// foo sends data to passed send channel
func foo(c chan<- int) {
	c <- 42
}

// bar reads data off the passed receive channel
func bar(c <-chan int) {
	fmt.Println("Bar received", <-c)
}
