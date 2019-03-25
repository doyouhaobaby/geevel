package main

import (
	"fmt"
	"sync"
	"runtime"
)


var (
	counter int
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCount(1)
	go incCount(2)

	wg.Wait()

	fmt.Println("完结统计结果", counter)
}

func incCount(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()

		value++

		counter = value
	}
}