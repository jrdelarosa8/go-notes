package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0

	const routines = 100
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementer)
}
