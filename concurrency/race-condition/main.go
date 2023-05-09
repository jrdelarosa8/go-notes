package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	counter := 0

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println("COUNTER:", counter)
}
