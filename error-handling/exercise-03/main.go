package main

import "fmt"

type customErr struct {
	info string
}

func main() {
	err := customErr{
		info: "Not enough coffee...",
	}

	foo(err)
}

func (c customErr) Error() string {
	return fmt.Sprintf("Error: %s", c.info)
}

func foo(e error) {
	fmt.Println(e.Error())
}
