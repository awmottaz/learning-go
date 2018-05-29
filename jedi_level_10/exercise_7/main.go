package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	c := make(chan int)
	var count int64

	numGoRoutines := 10

	for gr := 0; gr < numGoRoutines; gr++ {
		go func(gr int) {
			for i := 0; i < 10; i++ {
				c <- i + 10*gr
			}
			atomic.AddInt64(&count, 1)
			if atomic.LoadInt64(&count) == int64(numGoRoutines) {
				close(c)
			}
		}(gr)
	}

	for v := range c {
		fmt.Println(v)
	}
}
