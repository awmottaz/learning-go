package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementor int64

	fmt.Printf("BEFORE\n\tCPUs:\t%v\n\tGoroutines:\t%v\n", runtime.NumCPU(), runtime.NumGoroutine())

	numIterations := 100
	wg.Add(numIterations)

	for i := 0; i < numIterations; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Printf("Value: %v\n", atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("AFTER\n\tCPUs:\t%v\n\tGoroutines:\t%v\n", runtime.NumCPU(), runtime.NumGoroutine())
	fmt.Println(incrementor)
}
