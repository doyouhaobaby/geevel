package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main () {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("开始调度")

	go func() {
		defer wg.Done()

		for  count := 0; count <  3; count ++ {
			for char := 'a'; char < 'a' + 26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count :=0; count<3; count++ {
			for char := 'A'; char <'A' +26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("等待完成")
	wg.Wait()

	fmt.Println("完成")
}