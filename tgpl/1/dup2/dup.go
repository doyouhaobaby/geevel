package main 

import (
	"fmt"
	"bufio"
	"os"
)

//go run dup.go ./hello.txt
func main () {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n",err)
				continue
			}
			countLines(f, counts)
		}
	}

	//fmt.Println("yy %v", counts)

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d : %s \n\n",n,line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		//fmt.Printf(input.Text())
		counts[input.Text()]++
	}
	//fmt.Println("xxx %v", counts)
}