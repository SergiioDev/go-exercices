package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Printf("CPUs: %v\nGoroutines: %v\n", runtime.NumCPU(), runtime.NumGoroutine())

	var counter int64

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	//	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//mu.Lock()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Printf("Counter: %v\n", atomic.LoadInt64(&counter))
			//mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("CPUs: %v\nGoroutines: %v\n", runtime.NumCPU(), runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
