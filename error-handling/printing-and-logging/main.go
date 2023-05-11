package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f1, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	log.SetOutput(f1)

	f2, err := os.Open("fake.txt")
	if err != nil {
		log.Println("Error happened...", err)
	}

	defer f2.Close()

	fmt.Println("Check log.txt in the directory...")
}
