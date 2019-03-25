package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("开始调度")

	go foo("a")

	go foo("b")

	fmt.Println("等待完成")
	
	wg.Wait()

	fmt.Println("执行结束")
}

func foo (prefix string) {
	defer wg.Done()
//	println(prefix)

	next:
	 for outer := 2; outer < 5000; outer++ {
		 for inner := 2; inner < outer; inner++ {
			 if outer % inner == 0 {
				 continue next
			 }
		 }
		 fmt.Printf("%s: %d \n", prefix, outer)
	 }
}