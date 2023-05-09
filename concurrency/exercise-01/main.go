package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	wg.Add(2)

	go func() {
		fmt.Println("Thing 1")
		wg.Done()
	}()

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	go func() {
		fmt.Println("Thing 2")
		wg.Done()
	}()

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
}
