package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("throw", i)
	}
}

func cacher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("catch", num)
	}
}

func main() {
	c := make(chan int, 3)
	go thrower(c)
	go cacher(c)
	time.Sleep(100 * time.Microsecond)
}
