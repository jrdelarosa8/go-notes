package main

import "fmt"

func main() {
	e := make(chan int)
	o := make(chan int)
	q := make(chan bool)

	go send(e, o, q)

	receive(e, o, q)

	fmt.Println("ABOUT TO EXIT...")
}

// send
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

// receive
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("THE VALUE RECEIVED FROM EVEN CHANNEL:", v)
		case v := <-o:
			fmt.Println("THE VALUE RECEIVED FROM ODD CHANNEL:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("FROM COMMA OK:", i, ok)
				return
			} else {
				fmt.Println("FROM COMMA OK:", i)
			}
		}
	}
}
