package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	if !ok {
		fmt.Println("NOTHING ON THE CHANNEL")
	} else {
		fmt.Println(v, ok)
	}

	close(c)

	v, ok = <-c
	if !ok {
		fmt.Println("NOTHING ON THE CHANNEL")
	} else {
		fmt.Println(v, ok)
	}
}
