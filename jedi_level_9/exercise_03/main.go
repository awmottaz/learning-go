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
	for i := 0; i < numIterations; i++ {
		go func() {
			tempValue := incrementor
			runtime.Gosched()
			tempValue++
			incrementor = tempValue
			fmt.Printf("Value: %v\n", incrementor)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("AFTER\n\tCPUs:\t%v\n\tGoroutines:\t%v\n", runtime.NumCPU(), runtime.NumGoroutine())
	fmt.Println(incrementor)
}
