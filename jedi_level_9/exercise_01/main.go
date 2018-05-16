package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go foo()
	wg.Wait()
}

func foo() {
	fmt.Println("Hello from foo")
	wg.Done()
}
