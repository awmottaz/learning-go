package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementor := 0

	fmt.Printf("BEFORE\n\tCPUs:\t%v\n\tGoroutines:\t%v\n", runtime.NumCPU(), runtime.NumGoroutine())

	numIterations := 100
	wg.Add(numIterations)

	mutex := &sync.Mutex{}
	for i := 0; i < numIterations; i++ {
		go func() {
			mutex.Lock()
			val := incrementor
			val++
			incrementor = val
			fmt.Printf("Value: %v\n", incrementor)

			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("AFTER\n\tCPUs:\t%v\n\tGoroutines:\t%v\n", runtime.NumCPU(), runtime.NumGoroutine())
	fmt.Println(incrementor)
}
