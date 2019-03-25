package main

import (
	"time"
	"sync"
	"sync/atomic"
	"fmt"
)

var (
	shutdown int64

	wg sync.WaitGroup
)


func main() {
	wg.Add(2)

	go doWork("A")		
	go doWork("B")		

	time.Sleep(1 * time.Second)
	fmt.Println("停止工作")

	atomic.StoreInt64(&shutdown, 1)
}
 
func doWork(x string) {
	defer wg.Done()

	for {
		fmt.Printf("do %s work \n", x)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("工作 %s  停止\n", x)
			break;
		}
	}

}
