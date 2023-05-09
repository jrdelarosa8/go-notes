package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	incrementer := 0

	const routines = 100
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementer)
}
