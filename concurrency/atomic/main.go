package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	var counter int64

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("COUNTER:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println("COUNTER:", counter)
}
