package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bar() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
