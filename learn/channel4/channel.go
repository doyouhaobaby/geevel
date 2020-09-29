package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "hello world!"
}

func callerB(c chan string) {
	c <- "dddd!"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg := <-a:
			fmt.Printf("%s A \n", msg)
		case msg := <-b:
			fmt.Printf("%s B \n", msg)
		default:
			fmt.Println("default")
		}
	}
}
