package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	var wg sync.WaitGroup
	var maxNumbers = 100

	wg.Add(maxNumbers)
	for i := 0; i < maxNumbers; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", atomic.LoadInt64(&counter))

}
