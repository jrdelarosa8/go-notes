package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
