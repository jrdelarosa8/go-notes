package main

import "fmt"

func main() {
	c := make(chan int)

	send(c)
	receive(c)

	fmt.Println("ABOUT TO EXIT...")
}

// put 100 integers onto a channel
func send(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}

// pull 100 integers off of a channel
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
