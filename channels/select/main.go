package main

import "fmt"

func main() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	// send
	go send(e, o, q)

	// receive
	receive(e, o, q)

	fmt.Println("ABOUT TO QUIT...")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("EVEN CHANNEL:", v)
		case v := <-o:
			fmt.Println("ODD CHANNEL:", v)
		case v := <-q:
			fmt.Println("QUIT CHANNEL:", v)
			return
		}
	}
}
