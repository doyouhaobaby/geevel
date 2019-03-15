package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {
	fmt.Println("hello world")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line,n := range counts {
		//if n > 1 {
			fmt.Printf("%d - %s \n", n, line)
		//}
	}
}
